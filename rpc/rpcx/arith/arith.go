package arith

import "context"

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith)Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

type Arith2 int

func (t *Arith2)Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B *10
	return nil
}