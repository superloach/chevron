#!/usr/bin/env python3
import sys, re, random

def decap(regex):
	expr = r"""\(([^?](?:[^:](?:[^)]+)?)?)?\)"""
	return re.sub(expr, r"(?:\1)", regex)

class TXT(str):
	def __new__(cls, value):
		if '__txt__' in dir(value):
			value = value.__txt__()
		return  super(TXT, cls).__new__(cls, value)

	def __repr__(self):
		return '<STR %s>' % str(self).__repr__()

class NUM(int):
	regex = r"(-?[0-9]*(?:\.[0-9]+)?)"

	def __new__(cls, value):
		if '__num__' in dir(value):
			value = value.__num__()
		return  super(NUM, cls).__new__(cls, value)

	def __repr__(self):
		return '<NUM %s>' % self

	def __str__(self):
		return self.__txt__()

	def __txt__(self):
		return str(int(self))

class VAR:
	regex = "\^(_.|.)"
	all = {}

	def __init__(self, name):
		self.name = name

	def set(self, value):
		VAR.all[self.name] = value

	def get(self):
		return VAR.all.get(self.name, None)

	def __str__(self):
		return self.name

	def __repr__(self):
		return '<VAR %s=%s>' % (self.name, self.get().__repr__())

class MIX:
	def __init__(self, text):
		self.text = TXT(text)

	def __txt__(self):
		out = self.text

		for name in VAR.all:
			var = VAR(name)
			regex = r"\^%s" % re.escape(var.name)
			out = re.sub(regex, TXT(var.get()), out)

		return out

	def __num__(self):
		return TXT(self)

class MAT:
	oper = {
		'+': lambda a, b: a + b,
		'-': lambda a, b: a - b,
		'/': lambda a, b: a // b,
		'*': lambda a, b: a * b,
		'%': lambda a, b: a % b,
		'<': lambda a, b: a < b,
		'>': lambda a, b: a > b,
		'=': lambda a, b: a == b,
	}
	test = {
		'p': lambda n: not sum([n % x == 0 for x in range(2, n)]),
		'o': lambda n: not not (n % 2),
		'e': lambda n: not (n % 2),
		'r': lambda n: random.randint(0, n),
	}
	opr = '|'.join([re.escape(o) for o in [*oper.keys()] + ['~']])
	ter = '|'.join([re.escape(t) for t in test.keys()])
	var = decap(VAR.regex)
	num = decap(NUM.regex)

	regex = r"((?:%s)|%s)(?:(%s)((?:%s)|%s|%s))?" % (var, num, opr, var, num, ter)

	def parse(text):
		match = re.match(r'^%s$' % MAT.regex, text)
		return MAT(*match.groups())

	def __init__(self, a, oper, b):
		self.a = MIX(a)
		self.oper = oper
		self.b = MIX(b)

	def __num__(self):
		self.a = NUM(self.a)

		if None in [self.oper, self.b]:
			return self.a

		if self.oper == '~':
			test = MAT.test[str(TXT(self.b))]
			value = test(self.a)
		else:
			self.b = NUM(self.b)
			oper = MAT.oper[self.oper]
			value = oper(self.a, self.b)

		return NUM(value)

	def __txt__(self):
		return NUM(self)

class COM:
	regex = "<>(.*)"

	def __init__(self, comment):
		self.comment = MIX(comment)
		self.var = VAR('_c')

	def __call__(self):
		self.var.set(TXT(self.comment))

class OUT:
	regex = r">([^<>][^>]*)"

	def __init__(self, text):
		self.text = MIX(text)

	def __call__(self):
		value = TXT(self.text)
		print(value)

class TIN:
	regex = r">([^>]*)>\^?(.)"

	def __init__(self, prompt, var):
		self.prompt = MIX(prompt)
		self.var = VAR(var)

	def __call__(self):
		prompt = TXT(self.prompt)
		text = input(prompt)
		self.var.set(TXT(text))

class NIN:
	regex = r">([^>]*)>>\^?(.)"

	def __init__(self, prompt, var):
		self.prompt = MIX(prompt)
		self.var = VAR(var)

	def __call__(self):
		prompt = TXT(self.prompt)
		number = input(prompt)
		self.var.set(NUM(number))

class TAS:
	regex = r"%s<([^<>][^>]*)" % VAR.regex

	def __init__(self, var, text):
		self.text = MIX(text)
		self.var = VAR(var)

	def __call__(self):
		self.var.set(TXT(self.text))

class NAS:
	regex = r"%s<<(%s)" % (VAR.regex, decap(MAT.regex))

	def __init__(self, var, expr):
		self.expr = MAT.parse(expr)
		self.var = VAR(var)

	def __call__(self):
		self.var.set(NUM(self.expr))

class IDX:
	regex = r"%s<(%s)>(.*)" % (VAR.regex, decap(MAT.regex))

	def __init__(self, to, idx, frm):
		self.to = VAR(to)
		self.idx = MAT.parse(idx)
		self.frm = MIX(frm)

	def __call__(self):
		index = NUM(math)
		string = TXT(self.frm)
		if index > len(string):
			char = ''
		else:
			char = string[index - 1]
		self.to.set(TXT(char))

class HOP:
	regex = r"->(.)(%s)" % decap(MAT.regex)

	def __init__(self, expos, line):
		self.rel = False
		if expos == '+':
			self.rel = True
		else:
			line = expos + line
		self.line = MAT.parse(line)
		self.var = VAR('_#')

	def __call__(self):
		line = NUM(self.line) - 1
		if line < 0: self.rel = True
		if self.rel: line = self.var.get() + line
		self.var.set(line)

class SKP:
	regex = r"->(.)(%s)\?(%s)" % (decap(MAT.regex), decap(MAT.regex))

	def __init__(self, expos, line, expr):
		self.rel = False
		if expos == '+':
			self.rel = True
		else:
			line = expos + line
		self.expr = MAT.parse(expr)
		self.line = MAT.parse(line)
		self.var = VAR('_#')

	def __call__(self):
		do = NUM(self.expr)
		if do:
			line = NUM(self.line) - 1
			if line < 0: self.rel = True
			if self.rel: line = self.var.get() + line
			self.var.set(line)

class JMP:
	regex = r"->(.)(%s)\?\?([^=]*)=(.*)" % decap(MAT.regex)

	def __init__(self, expos, line, mix1, mix2):
		self.rel = False
		if expos == '+':
			self.rel = True
		else:
			line = expos + line
		self.line = MAT.parse(line)
		self.mix1 = MIX(mix1)
		self.mix2 = MIX(mix2)
		self.var = VAR('_#')

	def __call__(self):
		txt1 = TXT(self.mix1)
		txt2 = TXT(self.mix2)
		if txt1 == txt2:
			line = NUM(self.line) - 1
			if line < 0: self.rel = True
			if self.rel: line = self.var.get() + line
			self.var.set(line)

class DIE:
	regex = r"><(.*)"

	def __init__(self, text):
		self.text = MIX(text)

	def __call__(self):
		value = TXT(self.text)
		if len(value): sys.exit(value)
		sys.exit()

def find(line):
	for cmd in [COM, OUT, TIN, NIN, TAS, NAS, IDX, HOP, SKP, JMP, DIE]:
		match = re.match(r"^%s$" % cmd.regex, line)
		if match:
			argc = cmd.__init__.__code__.co_argcount
			args = [*match.groups()][:argc - 1]
			command = cmd(*args)
			return command

def main(filename):
	VAR('_r').set('>')
	VAR('_l').set('<')
	VAR('_u').set('^')
	VAR('_q').set('?')
	VAR('_d').set('-')
	VAR('_e').set('=')
	VAR('_n').set('\n')
	VAR('__').set('')
	VAR('_#').set(1)
	VAR('_i').set(sys.stdin.isatty())

	file = open(filename)

	lines = []
	program = []
	for line in file:
		line = line.rstrip('\n')
		lines.append(line)
		if len(line):
			cmd = find(line)
			program.append(cmd)

	file.close()

	start = 0
	if None in program:
		while None in program[start:]:
			line = start + program[start:].index(None)
			print("couldn't parse line %d" % line)
			print(lines[line].__repr__())
			start = line + 1

		DIE('errors parsing lines, good luck ^_u~^_u')()

	linenum = VAR('_#')

	while 1:
		try:
			cmd = program[int(linenum.get()) - 1]
		except IndexError:
			DIE('')()

		cmd()
#		print(VAR.all)

		linenum.set(linenum.get() + 1)

if __name__ == '__main__':
	main(*sys.argv[1:])
