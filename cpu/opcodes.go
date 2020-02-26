package cpu

type Opcode byte

func buildInstruction(c int, op Operation, src, dest register) Instruction {
	return Instruction{
		Cycles:      c,
		Operation:   op,
		Source:      src,
		Destination: dest,
	}
}

var InstructionSet = map[Opcode]Instruction{
	// 8-Bit Transfer Instructions
	// NOTE: A 111, B 000, C 001, D 010, E 011, H 100, L 101
	// LD r,r' (r <- r') 1
	Opcode(0x1<<6 | a<<3 | a): buildInstruction(1, LD, a, a),
	Opcode(0x1<<6 | a<<3 | b): buildInstruction(1, LD, b, a),
	Opcode(0x1<<6 | a<<3 | c): buildInstruction(1, LD, c, a),
	Opcode(0x1<<6 | a<<3 | d): buildInstruction(1, LD, d, a),
	Opcode(0x1<<6 | a<<3 | e): buildInstruction(1, LD, e, a),
	Opcode(0x1<<6 | a<<3 | h): buildInstruction(1, LD, h, a),
	Opcode(0x1<<6 | a<<3 | l): buildInstruction(1, LD, l, a),
	Opcode(0x1<<6 | b<<3 | a): buildInstruction(1, LD, a, b),
	Opcode(0x1<<6 | b<<3 | b): buildInstruction(1, LD, b, b),
	Opcode(0x1<<6 | b<<3 | c): buildInstruction(1, LD, c, b),
	Opcode(0x1<<6 | b<<3 | d): buildInstruction(1, LD, d, b),
	Opcode(0x1<<6 | b<<3 | e): buildInstruction(1, LD, e, b),
	Opcode(0x1<<6 | b<<3 | h): buildInstruction(1, LD, h, b),
	Opcode(0x1<<6 | b<<3 | l): buildInstruction(1, LD, l, b),
	Opcode(0x1<<6 | c<<3 | a): buildInstruction(1, LD, a, c),
	Opcode(0x1<<6 | c<<3 | b): buildInstruction(1, LD, b, c),
	Opcode(0x1<<6 | c<<3 | c): buildInstruction(1, LD, c, c),
	Opcode(0x1<<6 | c<<3 | d): buildInstruction(1, LD, d, c),
	Opcode(0x1<<6 | c<<3 | e): buildInstruction(1, LD, e, c),
	Opcode(0x1<<6 | c<<3 | h): buildInstruction(1, LD, h, c),
	Opcode(0x1<<6 | c<<3 | l): buildInstruction(1, LD, l, c),
	Opcode(0x1<<6 | d<<3 | a): buildInstruction(1, LD, a, d),
	Opcode(0x1<<6 | d<<3 | b): buildInstruction(1, LD, b, d),
	Opcode(0x1<<6 | d<<3 | c): buildInstruction(1, LD, c, d),
	Opcode(0x1<<6 | d<<3 | d): buildInstruction(1, LD, d, d),
	Opcode(0x1<<6 | d<<3 | e): buildInstruction(1, LD, e, d),
	Opcode(0x1<<6 | d<<3 | h): buildInstruction(1, LD, h, d),
	Opcode(0x1<<6 | d<<3 | l): buildInstruction(1, LD, l, d),
	Opcode(0x1<<6 | e<<3 | a): buildInstruction(1, LD, a, e),
	Opcode(0x1<<6 | e<<3 | b): buildInstruction(1, LD, b, e),
	Opcode(0x1<<6 | e<<3 | c): buildInstruction(1, LD, c, e),
	Opcode(0x1<<6 | e<<3 | d): buildInstruction(1, LD, d, e),
	Opcode(0x1<<6 | e<<3 | e): buildInstruction(1, LD, e, e),
	Opcode(0x1<<6 | e<<3 | h): buildInstruction(1, LD, h, e),
	Opcode(0x1<<6 | e<<3 | l): buildInstruction(1, LD, l, e),
	Opcode(0x1<<6 | h<<3 | a): buildInstruction(1, LD, a, h),
	Opcode(0x1<<6 | h<<3 | b): buildInstruction(1, LD, b, h),
	Opcode(0x1<<6 | h<<3 | c): buildInstruction(1, LD, c, h),
	Opcode(0x1<<6 | h<<3 | d): buildInstruction(1, LD, d, h),
	Opcode(0x1<<6 | h<<3 | e): buildInstruction(1, LD, e, h),
	Opcode(0x1<<6 | h<<3 | h): buildInstruction(1, LD, h, h),
	Opcode(0x1<<6 | h<<3 | l): buildInstruction(1, LD, l, h),
	Opcode(0x1<<6 | l<<3 | a): buildInstruction(1, LD, a, l),
	Opcode(0x1<<6 | l<<3 | b): buildInstruction(1, LD, b, l),
	Opcode(0x1<<6 | l<<3 | c): buildInstruction(1, LD, c, l),
	Opcode(0x1<<6 | l<<3 | d): buildInstruction(1, LD, d, l),
	Opcode(0x1<<6 | l<<3 | e): buildInstruction(1, LD, e, l),
	Opcode(0x1<<6 | l<<3 | h): buildInstruction(1, LD, h, l),
	Opcode(0x1<<6 | l<<3 | l): buildInstruction(1, LD, l, l),
	// LD r,n (r <- n) 2
	Opcode(0x0 | a<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x0 | b<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x0 | c<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x0 | d<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x0 | e<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x0 | h<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x0 | l<<3 | 0x6): buildInstruction(2, LD, a, a),
	// LD r,(HL) (r <- (HL)) 2
	Opcode(0x1<<6 | a<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | b<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | c<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | d<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | e<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | h<<3 | 0x6): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | l<<3 | 0x6): buildInstruction(2, LD, a, a),
	// LD (HL),r ((HL) <- r) 2
	Opcode(0x1<<6 | 0x6<<3 | a): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | 0x6<<3 | b): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | 0x6<<3 | c): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | 0x6<<3 | d): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | 0x6<<3 | e): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | 0x6<<3 | h): buildInstruction(2, LD, a, a),
	Opcode(0x1<<6 | 0x6<<3 | l): buildInstruction(2, LD, a, a),
	// LD (HL),n ((HL) <- n) 3
	Opcode(0x0 | 0x6<<3 | 0x6): buildInstruction(3, LD, a, a),
	// LD A,(BC) (A <- BC) 2
	Opcode(0x0 | 0x1<<3 | 0x2): buildInstruction(2, LD, a, a),
	// LD A,(DE) (A <- DE) 2
	Opcode(0x0 | 0x3<<3 | 0x2): buildInstruction(2, LD, a, a),
	// LD A,(C) (A <- (FF00H+C)) 2
	Opcode(0x3<<6 | 0x6<<3 | 0x2): buildInstruction(2, LD, a, a),
	// LD (C),A ((FF00H+C) <- A) 2
	Opcode(0x3<<6 | 0x4<<3 | 0x2): buildInstruction(2, LD, a, a),
	// LD A,(n) (A <- (n)) 3
	Opcode(0x3<<6 | 0x6<<3 | 0x0): buildInstruction(3, LD, a, a),
	// LD (n),A ((n) <- A) 3
	Opcode(0x3<<6 | 0x4<<3 | 0x0): buildInstruction(3, LD, a, a),
	// LD A,(nn) (A <- (nn)) 4
	Opcode(0x3<<6 | 0x7<<3 | 0x2): buildInstruction(4, LD, a, a),
	// LD (nn),A ((nn) <- A) 4
	Opcode(0x3<<6 | 0x5<<3 | 0x2): buildInstruction(4, LD, a, a),

	// LD A,(HLI) (A <- HL, HL <- HL+1) 2
	Opcode(0x0 | 0x5<<3 | 0x2): buildInstruction(2, LD, a, a),
	// LD A,(HLD) (A <- HL, HL <- HL-1) 2
	Opcode(0x0 | 0x7<<3 | 0x2): buildInstruction(2, LD, a, a),
	// LD (BC),A ((BC) <- A) 2
	Opcode(0x0 | 0x0 | 0x2): buildInstruction(2, LD, a, a),
	// LD (DE),A ((DE) <- A) 2
	Opcode(0x0 | 0x2<<3 | 0x2): buildInstruction(2, LD, a, a),
	// LD (HLI),A ((HL) <- A, HL <- HL+1) 2
	Opcode(0x0 | 0x4<<3 | 0x2): buildInstruction(2, LD, a, a),
	// LD (HLD),A ((HL) <- A, HL <- HL-1) 2
	Opcode(0x0 | 0x6<<3 | 0x2): buildInstruction(2, LD, a, a),

	// 16-Bit Transfer Instructions
	// NOTE: dd = BC 00, DE 01, HL 10, SP 11
	// NOTE: qq = BC 00, DE 01, HL 10, AF 11
	// LD dd,nn (dd <- nn) 3
	// LD SP, HL (SP <- HL) 2
	// PUSH qq ((SP-1)<-qqH (SP-2)<-qqL SP<-(SP-2)) 4
	// POP qq (qqL <- (SP) qqH<-(SP+1) SP<-(SP+2)) 3
	// NOTE: Affects flags - added before cycle count.
	// LDHL SP,e (HL <- SP+e) **00 3
	// LD (nn),SP ((nn)<-SPL (nn+1)<-SPH) 5

	// 8-Bit Arithmetic and Logical Operation Instructions
	// NOTE: s is any of r,n,(HL)
	// NOTE: CYCLE - 1 when s is r, 2 when s is n or (HL)
	// TODO: Figure out how to handle that in buildInstruction...
	// NOTE: FLAGS added before cycles (CY,H,N,Z) - flag is affected according
	//       to the result of the operation.
	//       Z: Zero flag. z=1 if the result of the operation is 0
	//       C: Carry/liink flag. C=1 if the operation produced a carry
	//       from the MSB of the operand or result
	//       H: Half-carry flag
	//       N: Add/Subject flag.
	// TODO: Figure out how to handle that in buildInstruction...
	// ADD A,r (A<-A+r) **0* 1
	// ADD A,n (A<-A+n) **0* 2
	// ADD A,(HL) (A<-A+(HL)) **0* 2
	// ADC A,s (A<-A+s+CY) **0* 1,2
	// SUB s (A<-A-s) **1* 1,2
	// SBC A,s (A<-A-s-CY) **1* 1,2
	// AND s (A<-A&s) 010* 1,2
	// OR s (A|s) 000* 1,2
	// XOR s (A^s) 0008 1,2
	// CP s (A-s) **1* 1,2
	// INC r (r<-r+1) -*0* 1
	// INC (HL) ((HL) <- (HL)+1) -*0* 3
	// DEC r (r <- r-1) -*1* 3
	// DEC (HL) ((HL) <- (HL)-1) -*1* 3

	// 16-Bit Arithmetic Operation Instructions
	// NOTE: BC 00, DE 01, HL 10, SP 11
	// ADD HL,ss (HL <- HL+ss) **0- 2
	// ADD SP,e (SP <- SP+e) **00 4
	// INC ss (ss <- ss+1) 2
	// DEC ss (ss <- ss-1) 2
}
