package balance

type Balance interface{
	DoBalance( []*Instance, ...string) (inst *Instance,err error)
}

