package app

/*B(Import)*/
	import "github.com/hailongz/kk-micro/micro"
	/*E(Import)*/

type /*B(Result)*/ QueryTaskResult /*E(Result)*/ struct {
	/*B(Result.Base)*/
	micro.QueryTaskResult
	/*E(Result.Base)*/

	/*B(Output)*/ /*E(Output)*/
	/*B(Output.todos)*/
	Todos []Todo `json:"todos,omitempty" title:"TODO"`
	/*E(Output.todos)*/
	/*B(Output.users)*/
	Users []User `json:"users,omitempty" title:"用户"`
	/*E(Output.users)*/
}

type /*B(Task)*/ QueryTask /*E(Task)*/ struct {
	/*B(Task.Base)*/
	micro.QueryTask
	/*E(Task.Base)*/

	/*B(Input)*/ /*E(Input)*/
	/*B(Input.endTime)*/
	EndTime interface{} `json:"endTime" title:"开始时间(秒)"`
	/*E(Input.endTime)*/
	/*B(Input.startTime)*/
	StartTime interface{} `json:"startTime" title:"开始时间(秒)"`
	/*E(Input.startTime)*/
	/*B(Input.q)*/
	Q string `json:"q" title:"关键字"`
	/*E(Input.q)*/
	/*B(Input.orderBy)*/
	OrderBy string `json:"orderBy" title:"desc 降序默认 asc 升序"`
	/*E(Input.orderBy)*/
	/*B(Input.type)*/
	Type string `json:"type" title:"类型 多个逗号分割"`
	/*E(Input.type)*/
	/*B(Input.uid)*/
	Uid int64 `json:"uid" title:"用户ID"`
	/*E(Input.uid)*/
	/*B(Input.id)*/
	Id int64 `json:"id" title:"钱包ID"`
	/*E(Input.id)*/

	/*B(Task.Result)*/
	Result QueryTaskResult `json:"-"`
	/*E(Task.Result)*/
}

/*B(name)*/
func (T *QueryTask) GetName() string {
	return "todo/query"
}

/*E(name)*/

/*B(title)*/
func (T *QueryTask) GetTitle() string {
	return "查询"
}

/*E(title)*/

/*B(Task.GetResult)*/
func (T *QueryTask) GetResult() interface{} {
	return &T.Result
}

/*E(Task.GetResult)*/
