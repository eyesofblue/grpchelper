package public

const (
	//打印到标准输出
	ConsoleLogConfig = `<seelog>
    <outputs formatid="main">   
        <filter levels="trace,info,debug,critical,error"> 
            <console />  
        </filter>
    </outputs>
    <formats>
        <format id="main" format="%Date %Time [%LEVEL] %File:%Line[%FuncShort] %Msg%n"/>  
    </formats>
</seelog>`

	//打印输出到文件
	FileLogConfig = `<seelog>
    <outputs formatid="main">   
        <filter levels="trace,info,debug,critical,error"> 
            <rollingfile type="date" filename="logs/log" fullname="true" datepattern="2006-01-02" maxrolls="365" />
        </filter>
    </outputs>
    <formats>
        <format id="main" format="%Date %Time [%LEVEL] %File:%Line[%FuncShort] %Msg%n"/>  
    </formats>
</seelog>`
)

