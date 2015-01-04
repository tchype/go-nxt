package commands

import "github.com/tonyheupel/go-nxt"

func NewStartProgram(filename string, replyChannel chan *nxt.ReplyTelegram) *nxt.Command {
	file := append([]byte(filename), 0) // null-terminated string

	return nxt.NewDirectCommand(0x00, file, replyChannel)
}

func NewStopProgram(replyChannel chan *nxt.ReplyTelegram) *nxt.Command {
	return nxt.NewDirectCommand(0x01, nil, replyChannel)
}

func NewGetCurrentProgramName(replyChannel chan *nxt.ReplyTelegram) *nxt.Command {
	return nxt.NewDirectCommand(0x11, nil, replyChannel)
}

type GetCurrentProgramNameReply struct {
	*nxt.ReplyTelegram
	Filename string
}

func ParseGetCurrentProgramNameReply(reply *nxt.ReplyTelegram) *GetCurrentProgramNameReply {
	return &GetCurrentProgramNameReply{
		ReplyTelegram: reply,
		Filename: string(reply.Message),
	}
}
