package main
  
import "fmt"

func myAtoi(s string) int {
    ret := 0
    if len(s) == 0 {
        return ret
    }
    sByteArr := []byte(s)
    isZhengshu := true
    for _, v :=range sByteArr{
        if v == ' ' {
            continue
        }
        if ret > 0 && (v == '-' || v == '+') {
            break
        }else if v == '-' {
            isZhengshu = false
            continue
        }
        isNum := (v >= '0' && v <= '9')
        if isNum == false {
            break
        }     
        val := int(v - '0')
        ret = ret * 10 + val
    }
    if isZhengshu == false {
        return -ret
    }else{
        return ret
    }
}

func main() {
    fmt.Println(myAtoi("  -420s2 "))
    fmt.Println("xxx")
}
