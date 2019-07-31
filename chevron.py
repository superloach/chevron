#!/usr/bin/env python3
import sys, re, random

global DEBUG
DEBUG = False

def decap(regex):
	expr = r"\(([^?](?:[^:](?:[^)]+)?)?)?\)"
	return re.sub(expr, r"(?:\1)", regex)

class TXT(str):
	def __new__(cls, value):
		if '__txt__' in dir(value):
			value = value.__txt__()
		return super(TXT, cls).__new__(cls, value)

	def __repr__(self):
		return '<TXT %s>' % str(self).__repr__()

class NUM(int):
	regex = r"(-?[0-9]*(?:\.[0-9]+)?)"

	def __new__(cls, value):
		if '__num__' in dir(value):
			value = value.__num__()
		return super(NUM, cls).__new__(cls, value)

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
			value = TXT(var.get()).replace('\\', '\\\\')
			regex = r"\^" + re.escape(var.name)
			out = re.sub(regex, value, out)

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
		'~': None,
	}
	special = {
		'p': lambda n: not sum([n % x == 0 for x in range(2, int(n))]),
		'o': lambda n: not not (int(n) % 2),
		'e': lambda n: not (int(n) % 2),
		'r': lambda n: random.randint(0, int(n)),
		'n': lambda n: not int(n),
		'l': lambda t: t.lower(),
		'u': lambda t: t.upper(),
		'v': lambda t: ''.join(t[::-1]),
	}
	opr = '|'.join([re.escape(o) for o in oper.keys()])
	opr2 = ''.join([*oper.keys()])

	regex = r"([^?|%s]+)(?:(%s)(.+))?" % (opr2, opr)

	def parse(text):
		if text[0] == '-': text = '0' + text
		expr = TXT(MIX(text))
		match = re.match(r'^%s$' % MAT.regex, expr)
		return MAT(*match.groups())

	def __init__(self, a, oper, b):
		self.a = a
		self.oper = oper
		self.b = b

	def __num__(self):
		if None in [self.oper, self.b]:
			return self.a

		if self.oper == '~':
			special = MAT.special[self.b]
			value = special(self.a)
		else:
			self.a = NUM(self.a)
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
	regex = r">([^>]+)>%s" % VAR.regex

	def __init__(self, prompt, var):
		self.prompt = MIX(prompt)
		self.var = VAR(var)

	def __call__(self):
		prompt = TXT(self.prompt)
		text = input(prompt)
		self.var.set(TXT(text))

class NIN:
	regex = r">([^>]+)>>%s" % VAR.regex

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
		self.expr = expr
		self.var = VAR(var)

	def __call__(self):
		value = MAT.parse(self.expr)
		self.var.set(NUM(value))

class IDX:
	regex = r"%s<(%s)>(.+)" % (VAR.regex, decap(MAT.regex))

	def __init__(self, to, idx, frm):
		self.to = VAR(to)
		self.idx = idx
		self.frm = MIX(frm)

	def __call__(self):
		index = NUM(MAT.parse(self.idx))
		string = TXT(self.frm)
		if index > len(string) or index == 0:
			char = ''
		else:
			index = index - (index > 0)
			char = string[index]
		self.to.set(TXT(char))

class CUT:
	regex = r"%s<(%s)\|(%s)>(.+)" % (VAR.regex, decap(MAT.regex), decap(MAT.regex))

	def __init__(self, to, idx, num, frm):
		self.to = VAR(to)
		self.idx = idx
		self.num = num
		self.frm = MIX(frm)

	def __call__(self):
		index = NUM(MAT.parse(self.idx))
		number = NUM(MAT.parse(self.num))
		string = TXT(self.frm)
		if index > len(string) or index == 0:
			char = ''
		else:
			index = index - (index > 0)
			char = string[index:index + number]
		self.to.set(TXT(char))

class HOP:
	regex = r"->(.)(%s)" % decap(MAT.regex)

	def __init__(self, expos, line):
		self.rel = False
		if expos == '+':
			self.rel = True
		else:
			line = expos + line
		self.line = line
		self.var = VAR('_#')

	def __call__(self):
		line = NUM(MAT.parse(self.line)) - 1
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
		self.expr = expr
		self.line = line
		self.var = VAR('_#')

	def __call__(self):
		do = NUM(MAT.parse(self.expr))
		if do:
			line = NUM(MAT.parse(self.line)) - 1
			if line < 0: self.rel = True
			if self.rel: line = self.var.get() + line
			self.var.set(line)

class JMP:
	ops = {
		'=': lambda a, b: a == b,
		'>': lambda a, b: b not in a,
		'<': lambda a, b: b in a
	}
	regex = r"->(.)(%s)\?\?([^=<>]+)([=<>])(.+)" % decap(MAT.regex)

	def __init__(self, expos, line, mix1, op, mix2):
		self.rel = False
		if expos == '+':
			self.rel = True
		else:
			line = expos + line
		self.line = line
		self.mix1 = MIX(mix1)
		self.op = JMP.ops[op]
		self.mix2 = MIX(mix2)
		self.var = VAR('_#')

	def __call__(self):
		txt1 = TXT(self.mix1)
		txt2 = TXT(self.mix2)
		if self.op(txt1, txt2):
			line = NUM(MAT.parse(self.line)) - 1
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
	for cmd in [COM, OUT, TIN, NIN, TAS, NAS, IDX, CUT, HOP, SKP, JMP, DIE]:
		match = re.match(r"^%s$" % cmd.regex, line)
		if match:
			argc = cmd.__init__.__code__.co_argcount
			args = [*match.groups()][:argc - 1]
			command = cmd(*args)
			return command

def run(interpreter, *args):
	filename = None
	pargs = []
	for arg in args:
		if arg == '--debug':
			DEBUG = True
		elif not filename:
			filename = arg
		else:
			pargs.append(arg)

	if not filename: DIE('filename must be provided')()

	VAR('_r').set('>')
	VAR('_l').set('<')
	VAR('_u').set('^')
	VAR('_q').set('?')
	VAR('_d').set('-')
	VAR('_e').set('=')
	VAR('_n').set('\n')
	VAR('_a').set('abcdefghijklmnopqrstuvwxyz')
	VAR('_i').set('1234567890')
	VAR('_g').set('\x00'.join(pargs))
	VAR('_x').set('\x00')
	VAR('__').set('')
	VAR('_#').set(1)
	VAR('_t').set(sys.stdin.isatty())

	file = open(filename)

	lines = []
	program = []
	for rline in file:
		for line in rline.rstrip('\n').split(';'):
			if len(line):
				lines.append(line)
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
		line = int(linenum.get()) - 1
		try:
			cmd = program[line]
		except IndexError:
			DIE('')()

		cmd()

		linenum.set(linenum.get() + 1)

if __name__ == '__main__':
	run(*sys.argv)
