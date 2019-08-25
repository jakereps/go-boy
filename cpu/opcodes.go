package cpu

type Opcode byte

var InstructionSet = map[Opcode]Instruction{
	// LD r,r' (r <- r') 1
	Opcode(0x1<<6 | a<<3 | a): LD,
	Opcode(0x1<<6 | a<<3 | b): LD,
	Opcode(0x1<<6 | a<<3 | c): LD,
	Opcode(0x1<<6 | a<<3 | d): LD,
	Opcode(0x1<<6 | a<<3 | e): LD,
	Opcode(0x1<<6 | a<<3 | h): LD,
	Opcode(0x1<<6 | a<<3 | l): LD,
	Opcode(0x1<<6 | b<<3 | a): LD,
	Opcode(0x1<<6 | b<<3 | b): LD,
	Opcode(0x1<<6 | b<<3 | c): LD,
	Opcode(0x1<<6 | b<<3 | d): LD,
	Opcode(0x1<<6 | b<<3 | e): LD,
	Opcode(0x1<<6 | b<<3 | h): LD,
	Opcode(0x1<<6 | b<<3 | l): LD,
	Opcode(0x1<<6 | c<<3 | a): LD,
	Opcode(0x1<<6 | c<<3 | b): LD,
	Opcode(0x1<<6 | c<<3 | c): LD,
	Opcode(0x1<<6 | c<<3 | d): LD,
	Opcode(0x1<<6 | c<<3 | e): LD,
	Opcode(0x1<<6 | c<<3 | h): LD,
	Opcode(0x1<<6 | c<<3 | l): LD,
	Opcode(0x1<<6 | d<<3 | a): LD,
	Opcode(0x1<<6 | d<<3 | b): LD,
	Opcode(0x1<<6 | d<<3 | c): LD,
	Opcode(0x1<<6 | d<<3 | d): LD,
	Opcode(0x1<<6 | d<<3 | e): LD,
	Opcode(0x1<<6 | d<<3 | h): LD,
	Opcode(0x1<<6 | d<<3 | l): LD,
	Opcode(0x1<<6 | e<<3 | a): LD,
	Opcode(0x1<<6 | e<<3 | b): LD,
	Opcode(0x1<<6 | e<<3 | c): LD,
	Opcode(0x1<<6 | e<<3 | d): LD,
	Opcode(0x1<<6 | e<<3 | e): LD,
	Opcode(0x1<<6 | e<<3 | h): LD,
	Opcode(0x1<<6 | e<<3 | l): LD,
	Opcode(0x1<<6 | h<<3 | a): LD,
	Opcode(0x1<<6 | h<<3 | b): LD,
	Opcode(0x1<<6 | h<<3 | c): LD,
	Opcode(0x1<<6 | h<<3 | d): LD,
	Opcode(0x1<<6 | h<<3 | e): LD,
	Opcode(0x1<<6 | h<<3 | h): LD,
	Opcode(0x1<<6 | h<<3 | l): LD,
	Opcode(0x1<<6 | l<<3 | a): LD,
	Opcode(0x1<<6 | l<<3 | b): LD,
	Opcode(0x1<<6 | l<<3 | c): LD,
	Opcode(0x1<<6 | l<<3 | d): LD,
	Opcode(0x1<<6 | l<<3 | e): LD,
	Opcode(0x1<<6 | l<<3 | h): LD,
	Opcode(0x1<<6 | l<<3 | l): LD,
	// LD r,n (r <- n) 2
	Opcode(0x0 | a<<3 | 0x6): LD,
	Opcode(0x0 | b<<3 | 0x6): LD,
	Opcode(0x0 | c<<3 | 0x6): LD,
	Opcode(0x0 | d<<3 | 0x6): LD,
	Opcode(0x0 | e<<3 | 0x6): LD,
	Opcode(0x0 | h<<3 | 0x6): LD,
	Opcode(0x0 | l<<3 | 0x6): LD,
	// LD r,(HL) (r <- (HL)) 2
	Opcode(0x1<<6 | a<<3 | 0x6): LD,
	Opcode(0x1<<6 | b<<3 | 0x6): LD,
	Opcode(0x1<<6 | c<<3 | 0x6): LD,
	Opcode(0x1<<6 | d<<3 | 0x6): LD,
	Opcode(0x1<<6 | e<<3 | 0x6): LD,
	Opcode(0x1<<6 | h<<3 | 0x6): LD,
	Opcode(0x1<<6 | l<<3 | 0x6): LD,
	// LD (HL),r ((HL) <- r) 2
	Opcode(0x1<<6 | 0x6<<3 | a): LD,
	Opcode(0x1<<6 | 0x6<<3 | b): LD,
	Opcode(0x1<<6 | 0x6<<3 | c): LD,
	Opcode(0x1<<6 | 0x6<<3 | d): LD,
	Opcode(0x1<<6 | 0x6<<3 | e): LD,
	Opcode(0x1<<6 | 0x6<<3 | h): LD,
	Opcode(0x1<<6 | 0x6<<3 | l): LD,
	// LD (HL),n ((HL) <- n) 3
	Opcode(0x0 | 0x6<<3 | 0x6): LD,
	// LD A,(BC) (A <- BC) 2
	Opcode(0x0 | 0x1<<3 | 0x2): LD,
	// LD A,(DE) (A <- DE) 2
	Opcode(0x0 | 0x3<<3 | 0x2): LD,
	// LD A,(C) (A <- (FF00H+C)) 2
	Opcode(0x3<<6 | 0x6<<3 | 0x2): LD,
	// LD (C),A ((FF00H+C) <- A) 2
	Opcode(0x3<<6 | 0x4<<3 | 0x2): LD,
	// LD A,(n) (A <- (n)) 3
	Opcode(0x3<<6 | 0x6<<3 | 0x0): LD,
	// LD (n),A ((n) <- A) 3
	Opcode(0x3<<6 | 0x4<<3 | 0x0): LD,
	// LD A,(nn) (A <- (nn)) 4
	Opcode(0x3<<6 | 0x7<<3 | 0x2): LD,
	// LD (nn),A ((nn) <- A) 4
	Opcode(0x3<<6 | 0x5<<3 | 0x2): LD,
	// LD A,(HLI) (A <- HL, HL <- HL+1) 2
	Opcode(0x0 | 0x5<<3 | 0x2): LD,
	// LD A,(HLD) (A <- HL, HL <- HL-1) 2
	Opcode(0x0 | 0x7<<3 | 0x2): LD,
	// LD (BC),A ((BC) <- A) 2
	Opcode(0x0 | 0x0 | 0x2): LD,
	// LD (DE),A ((DE) <- A) 2
	Opcode(0x0 | 0x2<<3 | 0x2): LD,
	// LD (HLI),A ((HL) <- A, HL <- HL+1) 2
	Opcode(0x0 | 0x4<<3 | 0x2): LD,
	// LD (HLD),A ((HL) <- A, HL <- HL-1) 2
	Opcode(0x0 | 0x6<<3 | 0x2): LD,
}
