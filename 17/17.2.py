import datetime


REGISTER = 'Register'

class Day17(object):
    def __init__(self, a, inputs):
        self._inputs = inputs
        self._outputs = []

        self._a = a  # ignore register_a value from contents
        self._b = 0
        self._c = 0

        self._instruction_pointer = 0
        self._has_jumped = False

    @property
    def inputs(self):
        return self._inputs

    @property
    def outputs(self):
        return self._outputs

    def raise_if_outputs_invalid(self):
        for i, output in enumerate(self._outputs):
            if output != self._inputs[i]:
                raise Exception()

    def print_outputs(self):
        print(','.join([str(x) for x in self.outputs]))

    def is_copy(self):
        if len(self._inputs) != len(self._outputs):
            return False

        for i in range(len(self._inputs)):
            if self._inputs[i] != self._outputs[i]:
                return False
        return True

    @property
    def a(self):
        return self._a

    @property
    def b(self):
        return self._b

    @property
    def c(self):
        return self._c

    @property
    def literal_operand(self):
        return self._operand

    @property
    def combo_operand(self):
        mapping = {
            0: 0,
            1: 1,
            2: 2,
            3: 3,
            4: self.a,
            5: self.b,
            6: self.c,
        }
        return mapping[self._operand]

    def run(self):
        while self._instruction_pointer < len(self.inputs) - 1:
            self._has_jumped = False
            opcode = self.inputs[self._instruction_pointer]
            self._operand = self.inputs[self._instruction_pointer + 1]
            self.run_instruction(opcode)
            if not self._has_jumped:
                self._instruction_pointer += 2

    def run_instruction(self, opcode):
        mapping = {
            0: self.adv,
            1: self.bxl,
            2: self.bst,
            3: self.jnz,
            4: self.bxc,
            5: self.out,
            6: self.bdv,
            7: self.cdv,
        }
        return mapping[opcode]()

    # 0
    def adv(self):
        self._a = self.a // (2 ** self.combo_operand)

    # 1
    def bxl(self):
        self._b = self.bitwise_xor(self.b, self.literal_operand)

    # 2
    def bst(self):
        self._b = self.combo_operand % 8

    # 3
    def jnz(self):
        if self.a == 0:
            return

        self._instruction_pointer = self.literal_operand
        self._has_jumped = True

    # 4
    def bxc(self):
        self._b = self.bitwise_xor(self.b, self.c)

    # 5
    def out(self):
        val = self.combo_operand % 8
        self._outputs.append(val)
        self.raise_if_outputs_invalid()

    # 6
    def bdv(self):
        self._b = self.a // (2 ** self.combo_operand)

    # 7
    def cdv(self):
        self._c = self.a // (2 ** self.combo_operand)

    def bitwise_xor(self, a, b):
        return a ^ b

def parse_contents(contents):
    contents, input_contents = contents.split('\n\n')
    registers = {}
    for line in contents.split('\n'):
        if line:
            register_data = line.split(': ')
            code = register_data[0][len(REGISTER) + 1:]
            register_value = int(register_data[1])
            registers[code] = register_value

    inputs = []
    for line in input_contents.split('\n'):
        if line:
            inputs_data = line.split(': ')
            for input in inputs_data[1].split(','):
                if input:
                    inputs.append(int(input))

    return registers, inputs

def is_valid(initial_A):
    return (
        (2 == ((((initial_A % 8) ^ 5) ^ (initial_A // (2 ** ((initial_A % 8) ^ 5)))) ^ 6) % 8) and
        (4 == (((((initial_A // 8) % 8) ^ 5) ^ ((initial_A // 8) // (2 ** (((initial_A // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (1 == ((((((initial_A // 8) // 8) % 8) ^ 5) ^ (((initial_A // 8) // 8) // (2 ** ((((initial_A // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (5 == (((((((initial_A // 8) // 8) // 8) % 8) ^ 5) ^ ((((initial_A // 8) // 8) // 8) // (2 ** (((((initial_A // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (7 == ((((((((initial_A // 8) // 8) // 8) // 8) % 8) ^ 5) ^ (((((initial_A // 8) // 8) // 8) // 8) // (2 ** ((((((initial_A // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (5 == (((((((((initial_A // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ ((((((initial_A // 8) // 8) // 8) // 8) // 8) // (2 ** (((((((initial_A // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (4 == ((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ (((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** ((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (5 == (((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ ((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** (((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (0 == ((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ (((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** ((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (3 == (((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ ((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** (((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (1 == ((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ (((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** ((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (6 == (((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ ((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** (((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (5 == ((((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ (((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** ((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (5 == (((((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ ((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** (((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (3 == ((((((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ (((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** ((((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (0 == (((((((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5) ^ ((((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // (2 ** (((((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) % 8) ^ 5)))) ^ 6) % 8) and
        (0 == (((((((((((((((initial_A // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8) // 8)
    )

def check(expected, y):
    return expected == ((((y % 8) ^ 5) ^ (y // (2 ** ((y % 8) ^ 5)))) ^ 6) % 8

def get_range(y_prev, y_prev_exp):
    return range(min_range(y_prev, y_prev_exp), max_range(y_prev, y_prev_exp))

def min_range(y_prev, y_prev_exp):
    return ((y_prev) * 8 ** y_prev_exp) // (8 ** (y_prev_exp - 1))

def max_range(y_prev, y_prev_exp):
    return ((y_prev + 1) * 8 ** y_prev_exp) // (8 ** (y_prev_exp - 1))

def solution(inputs):
    y16 = 0
    for y15 in get_range(y16, 16):
        if not check(inputs[15], y15):
            continue
        for y14 in get_range(y15, 15):
            if not check(inputs[14], y14):
                continue
            for y13 in get_range(y14, 14):
                if not check(inputs[13], y13):
                    continue
                for y12 in get_range(y13, 13):
                    if not check(inputs[12], y12):
                        continue
                    for y11 in get_range(y12, 12):
                        if not check(inputs[11], y11):
                            continue
                        for y10 in get_range(y11, 11):
                            if not check(inputs[10], y10):
                                continue
                            for y9 in get_range(y10, 10):
                                if not check(inputs[9], y9):
                                    continue
                                for y8 in get_range(y9, 9):
                                    if not check(inputs[8], y8):
                                        continue
                                    for y7 in get_range(y8, 8):
                                        if not check(inputs[7], y7):
                                            continue
                                        for y6 in get_range(y7, 7):
                                            if not check(inputs[6], y6):
                                                continue
                                            for y5 in get_range(y6, 6):
                                                if not check(inputs[5], y5):
                                                    continue
                                                for y4 in get_range(y5, 5):
                                                    if not check(inputs[4], y4):
                                                        continue
                                                    for y3 in get_range(y4, 4):
                                                        if not check(inputs[3], y3):
                                                            continue
                                                        for y2 in get_range(y3, 3):
                                                            if not check(inputs[2], y2):
                                                                continue
                                                            for y1 in get_range(y2, 2):
                                                                if not check(inputs[1], y1):
                                                                    continue
                                                                for y0 in get_range(y1, 1):
                                                                    if not check(inputs[0], y0):
                                                                        continue
                                                                    return y0

if __name__ == '__main__':
    with open('17/day_17_input.txt', 'r') as f:
        contents = f.read()

    registers, inputs = parse_contents(contents)

    start = datetime.datetime.now()
    initial_A = solution(inputs)
    program = Day17(initial_A, inputs)
    try:
        program.run()
        if program.is_copy():
            print('SOLUTION IS: ', initial_A)
    except:
        print('NOPE!')

    print('TOOK: ', datetime.datetime.now() - start)
