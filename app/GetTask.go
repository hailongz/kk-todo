package app

/*B(Import)*/
	/*E(Import)*/

type /*B(Result)*/ GetTaskResult /*E(Result)*/ struct {
	/*B(Result.Base)*/
	/*E(Result.Base)*/

	/*B(Output)*/ /*E(Output)*/
	/*B(Output.todo)*/
	Todo *Todo `json:"todo,omitempty" title:"TODO"`
	/*E(Output.todo)*/
	/*B(Output.user)*/
	User *User `json:"user,omitempty" title:"用户"`
	/*E(Output.user)*/
}

type /*B(Task)*/ GetTask /*E(Task)*/ struct {
	/*B(Task.Base)*/
	/*E(Task.Base)*/

	/*B(Input)*/ /*E(Input)*/
	/*B(Input.rid)*/
	Rid int64 `json:"rid" title:"根级ID"`
	/*E(Input.rid)*/
	/*B(Input.uid)*/
	Uid int64 `json:"uid" title:"用户ID"`
	/*E(Input.uid)*/
	/*B(Input.id)*/
	Id int64 `json:"id" title:"ID"`
	/*E(Input.id)*/

	/*B(Task.Result)*/
	Result GetTaskResult `json:"-"`
	/*E(Task.Result)*/
}

/*B(name)*/
func (T *GetTask) GetName() string {
	return "todo/get"
}

/*E(name)*/

/*B(title)*/
func (T *GetTask) GetTitle() string {
	return "获取"
}

/*E(title)*/

/*B(Task.GetResult)*/
func (T *GetTask) GetResult() interface{} {
	return &T.Result
}

/*E(Task.GetResult)*/
