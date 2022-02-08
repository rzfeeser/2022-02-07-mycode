type HttpError struct {
    Code int
    Reason string
}

func (e *SyntaxError) Error() string {
    return fmt.Sprintf("%v:%v: syntax error", e.Reason, e.Code)
}

type InternalError struct {
    Path string
}

func (e *InternalError) Error() string {
    return fmt.Sprintf("parse %v: internal error", e.Path)
}

if err := Foo(); err != nil {
    switch e := err.(type) {
    case *HttpError:
        // Do something interesting with e.Reason and e.Code
    case *InternalError:
        // Abort and file an issue.
    default:
        log.Println(e)
    }
}
