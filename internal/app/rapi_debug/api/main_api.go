package api

type MainApi struct {
}

func (r MainApi) GetIndex(args ArgsEmpty) string {
	return "Test"
}

func (r MainApi) GetFuck(args ArgsId) int {
	return 32
}

func (r MainApi) GetSuck(args ArgsAB) int {
	return args.A + args.B
}

func (r MainApi) PostSuper(args ArgsSuper) map[string]interface{} {
	return map[string]interface{}{
		"sas": 3,
		"gas": "xax",
	}
}
