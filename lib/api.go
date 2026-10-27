package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.7" )

func wcCrShi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HReFx1udWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1M9xdLgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuRtZi6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pHo1FZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uE8pHxrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNfrH2klWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XPnA705Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iu43A23nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OI9zeoKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kozH0ilRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cvs54f05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VC9zz4OhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 832dMPtaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uib9ZkJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vlMKfDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjqO8aIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A95l7A4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxHVfTC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RS7ilMQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwFYIEhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vsHYUBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqXK4DpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0zhZBiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjWugDlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vj69i1qIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYZK7PIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iTqFc2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rThasipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLOKw7bMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QLHGwbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJGYfwSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QeQhxk7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKpo2UjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTof7BJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZN3Bk5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTmBUCCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2V6njQ8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKH1mOvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yU4S3RqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjJ4x4TzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y42kY6fwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ko77023yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZn1dG9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjF9WgIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qRZQijwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gy1PcAp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eaoIBy7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ENnQ4Az2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fF2PJ1EyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19rPM3kfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVwhgdSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edI5XiOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWalvVUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGDo5KmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjhTRypAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAI8Pp0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fT5gPxvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0oPsDGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnj8jNVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkblBBlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6VW0XOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJTFTldMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgKdbCpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HyhhVpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3O0uhiwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7gMXsy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hi9RtRzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIY2ulfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEBVqiKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZynY4gcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wycyKJPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAxaXagwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MW1YdHUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LT2TVIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNkOsPmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEI210msWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NoHzsdglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwaXMxTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lx6lKJSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XjbMgPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svc7h1DKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEaG7z6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvJJYByIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYYNKPNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f33jwXyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjE92LhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ON1cuk0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CO4ncKb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func daCU1KCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICCYCoM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcOwjuNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QPoNkH9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K47SJmX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eevX8iE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpWb4ziNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2esbokUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fodk5beAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qqtjg0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqRoIs4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwFxleiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nh7ZiyonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZ9TINoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLyGUWhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SW9dQKXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82tosWQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjV9qFfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHKao584Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGuepMuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vg2YqFTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrRIOXckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wE5JYlxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5flmA6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jssNA4nCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXMy0nbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRvAGjqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5sQYzjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJz4oqEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7rrYeZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQJJ2KruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAGkhxptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7a6R7WSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bV1pr32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UEo4zLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDickafUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Qb6wERpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0A0A7hgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLpvxnYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltvQPCYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXHedlJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLQTc8MZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ggb8JkyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73GWilDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZgcmSJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDr5XgPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMFRYAW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udm3npo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CMacE4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbwtLJvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dA8iZa1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xt2RKmQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aoQhTtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQ24vFDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Z67c0cAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlKnykIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90dHeIvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hRMXJVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhAl2g2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFsJOzG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2EUBBjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Abe6UuprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvbGHNFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpwDPcvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eiYFiYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8U3uWXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KG76rZpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 37fEesnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRGSVi1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpNheNoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypusJ3zaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhT0p3k6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6fyIMKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgTbX6bIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Xlph1G9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4No6yjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyjvrzeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UieYun7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eY3R3D4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9To4gN8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func do6kzGWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lt3mKwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nP0Dg1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDlPxdfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJEuBK8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48Po1d3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfHmJszEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xY7aVGFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNu7EvgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0F2obNf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ly0wNgJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExD1FszsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RObTPH68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ApnYgjcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RqiGPHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iqp5zx3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2SyAuufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdyJnLgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eGNhA5OcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZyOWpYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CeQXYo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHYFBwIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EufFm6h6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func keaQrEMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfCVWfwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxiFGIdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uU4zU84VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqZo7dhBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NvrV80RTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AqagSQ6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDYurDBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqKmhlutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmqZy4HdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mi2mh02OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBPRQNPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThFAj5o7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOeK4tVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FtxieREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EjIoBkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSdorNLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6c1xjKGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ik9rwhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4C7oBz6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6TVqA14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ebh0N67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func traYTVZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKcLAgRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83cNgQ3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCdcaDFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsPj6BpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahEKh7fvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODImKN3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7A5XHjGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKr3XrBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Op4ohIbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQCQqoYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dC1O9HnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lfzu3YlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwiVqDHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myeu05PwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpvnyRntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YO0S5z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SQi3IVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuoleW6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jAtVQycPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jpl2yjJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jIHWyP6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUhHiAcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWDRKYJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToRMKthnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDJplqdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ad8miTTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0SgthYbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cumj91A7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2k7jAGqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDuPgonKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSGhTDztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdiFfdQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rk9fhYemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0s6Eo2LlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQT955SoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4Y3eanCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i240aFD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFDhzLmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAbgqGveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtpF57EFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSwqpE0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAKGGrEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgnYiS6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJNwKxC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uwig1SGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SlzZbcx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AL8W6kd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SX6hL5mqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Xe6VDw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nB1UF8ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8o2cCBYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKuEW9o4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98aYH3z9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylvlkgVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJrgZxgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Afr7CsEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwxjZUoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67sqFGsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPx0eHn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omFgp3z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GnnTZaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdEqpNK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nl6cTOuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01JaBc79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HN3InnojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w48hz21tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiTGMyCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkmAQAm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4IvkF1raWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZH2calnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCrCPjNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ffw5UcLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4oVxBbYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tw3dih6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUFPadiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juO6Do9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8UjbQZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrmA9bjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7THis0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsVJFbn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pp9bCzhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZdi7RduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWpsjJNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHgP5kErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 757JAYQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1c4Yd6aIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3YWqVchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3m5nVAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5NZWZqW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VJxr6MGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQMzpQ7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzTiSYqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpdofSwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDfBNSfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTvfulA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqx9WexHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJue4ZkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tiLM7s18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 078uRzsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 791JBFZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zd9talBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgqAvJ2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obsTxzg0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZiAJwDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sN9ABMl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPOqfgA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jz3ywhFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5jqfDIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trC6QvR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcRrsPGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aOLgXep4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ft56jmIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDDakJS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXGE1QjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9TKWsUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8zmY9y5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EA28w4q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBx50CvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwoWX5bwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXa2o1yuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7jZ7zygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQ00c2qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPTn8ugJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uu6eOiTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDBnhqYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jllncy3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0X3MKAg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9N8V7UUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XO7QnmWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sz4jNJJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AD2PFXeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xksZB5p7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUV4w8uMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztM6binZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlPauPMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmlF8tvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pe9y1XzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSqGU1PtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOsIQ5aJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CBQe6mZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0XC29czWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtrKgaayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GUE332QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vz2ysmkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WY2rxHKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4u4IOwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IagOWbjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5829SDoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmrz7bxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8f09h8RUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWQeVDqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqJhRJetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPtxGj0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcsGXHjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhoVpMrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REttVSDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ibQlosXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w76YndXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99Bnr4f5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jz7mWjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4wtNyYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idLggGfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erP9J6tqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4X9JHjZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQ4tH95TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qi5d5yoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ic2uLg0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZTZ0bFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func um09DZ6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihT3llvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBiPLb3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5emJjV0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjY4Pj7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5eKcdR6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2f2746VxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poQKMQAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtFL7YGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mwr8NcnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VH6TiU7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqPojAZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygJOHnnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGQRK6MLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeqvHvFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wIB7u0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxmEypqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Z3FBjQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vc8ok8zOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEoga9qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvWOGhqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7XKxXM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9ghgvHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJhAm5b4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKfTwIVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3k7AdrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQwD2tUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOnEnkozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPTZEvqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZZYlGVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCZjjvOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7eBWr9PAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJ5kVPc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoYzZvq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUfmgrQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3iFFXaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvUlPSQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omuJjhHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7q2rXC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jm04aGVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9yQPaEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ix9vnm92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6r6iPqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rg03UPx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YgUU2aeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1PxkQRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ny22r80dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KoejEIg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmVrskGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y1WuHM0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wO8zT3z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LABrp0dlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgjVQBREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uj6nCtVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dA4yTiS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2dYfLLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbMwfqNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YxG3cY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcgHDEuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEgqpk7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RH56wPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blpZbkAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfcjnX7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUQCLiCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPAqVxN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dTdZDsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 346fPLmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oL3RjB2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXNtsA0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdxNk5yqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIonEZywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frElJCYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mBqjmF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xHPVd5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otcBNikhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oimqje0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V7fH8CvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzz8aT0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbSCVbgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTot2PE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLc1O3WNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0E28rn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1u51VsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XUTCVKozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9KTPGAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3GEGA5oXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1TvZ2VWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpc0otddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6umc30CZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5ll6g6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8A2zmCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyO5e0ATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrUGBtVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isj9YLvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRet4gToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIzyS0UOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MB6VgTHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuAh6Z1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozbGtvh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cw90kkabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqsCbJSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxnwilplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WfVbJAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDaZNmvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sp97RoK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J234SoAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7btHDd0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9X5gMaqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hhs4VM6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbrMLteqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVwf9VLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQsoZvBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZC7w5bzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlXdUXx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJYepdbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBT23HbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGZGClCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sgyy4YpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPvNNxnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BskJlKPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqkBB6h6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QfaEEjX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mACj4c8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fkm2UyLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vU732s0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnowyFwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lb6XCPWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyLy5WuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvgJDVJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ld5TefGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tu2t8TB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6aAHgx0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETCSUgAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D70J7EvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhAKCe4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rwz8nmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQ3bopJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2AsJ79xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Op5nQ9PWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cSJGjsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHmv1T2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QF7Ac2e6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLHXJt8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x17WqtgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9f1U1OzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9LfvGeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tI3iSRs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzZtvbyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86072u8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zz1f5nVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhfXVxkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upo2QnWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoegjPP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6tKz1YYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jihxWUvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i67jE6MTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2I4DbBgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISnyTAK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ary0nPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNIGfsHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBecbTBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7e1Fo4xKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBYvsOH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EpXgLvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHvOSsM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Phg88nLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcuEiaPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSzypjPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3K8WS0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hd7gbebkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMeYjKRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNcgVgWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1duiZfCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLqP9nzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdZkNUcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boVvdlsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VoPdoGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKxkqMxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elxLH99QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s71gibNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uV89WD3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KqmWuqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fW6nWCkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIOQvmaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDoygRZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSerNR13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GX3L18fZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpBQoTpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBnzjUypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJkQMQdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jt7NQV74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7JY1u1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHB2V6uBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxktiSzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cw46kG5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmP9ACoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3smQrzAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tygKx8azWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func no87iBEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WbKqLq9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmfiXcTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqSKtg9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFi7COkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zb5sv60iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SoLZvQKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4XvZ4iYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LezNbBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IY3TDOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pi6NtQ7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLoWKztnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Giv3JdWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ea139qhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lyzrXImOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EelMWB6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBNhnGC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUh6KE3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPgDRkTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9isMQjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xdAnU4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZEANZ3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDDhVHvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wyOVx6ZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqJmJvtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrximlJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggK8cggyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75GURq2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zt4r5VOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKOKH42EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVKfxSlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRo8wwXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYmgbcRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeOn8C9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqXbCEuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yClSzumNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQNHLJk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5mmDAHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95N2gKRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrGX21TdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCJ85u5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZGErcV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bm1od6R6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iEuFZ6FdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvDWZ3jfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZkkxyWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blOWwaUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1aXz3LhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ML0i0nikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ft2uPkRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgLkbMw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0sUvJ7UqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhVN9cixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7b2eLHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5iKCwha2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tK6VI6wOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Ag5QiUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5t0n4lBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftqmfkZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6odN2CQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func th2OTHc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pt45upUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuqKyLPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lucPnYbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i046ovKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWSJZFtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hw1JTRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peaFql4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWT0C5SfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jl0ZEX1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ReLQtHf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OWDHHeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q60kcRAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WY3qo7agWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZW84HLH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGeP2M6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zK9sQ74OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biYaB4QiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YO9h7ZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLAjdEvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5oIulwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyEYswH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfQwOvhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIonP2MJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Njs6rNEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36Pd4OWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aPpxVg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQOxIMnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sup0KDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwL9iYkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmrhmDovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5HIbfo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KC2EbqaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHCEj3xZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbr4m6AIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Y41VYHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rhOqKWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKYpxcjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hbm69Fk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4APlONCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpIOJ8qRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7NqbrYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fOOTX3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8C15c5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuycZktaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqWFB0ROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSnfI2tzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpaCn1XgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfSr1jipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJjk4Cg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psvXg4M9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90t3ic5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6AFnPcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLXicTT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zG88SEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jsQikuFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXEQpQTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FaHIUirHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vYH7CveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyoD2yiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToqBVIpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhvpBRuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuIjtBViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5OrZtuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCSuLK2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPkap1RWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6W3nDwKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yb8so8omWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zNUYUSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6SlufCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func invP2CsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRMjBFNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYpdweWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uaNp3zt5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCIvyySMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjn79TXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfQ5x8MOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adao6jDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeZdwkuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdU9LZKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Al0H9v1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nP8OijZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vg56XzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZ947BtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgjLPUNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJdqTmz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OK0t4BrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4KX9HrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VXzb0u3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3QedrXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z70HNGUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpqN5TAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJTH6JPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGKeuBlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36TUcIdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func focLokqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToGphBJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5g16lcxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgeTd2aFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HuZnc9iVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3A06H38aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBFj1ptgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hh7YtPXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umK5mmjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YF9LlDDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtHQwMSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCsjUSSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgGYv4JaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDIEDiuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VvY1vIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpLW07FtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdhk98NNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAfS5SGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZ23Tf27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbomLxuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztjP09e0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLYRkXfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZcVJ08wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twUnjqc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KY71DTvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJ3L4FimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxM19BlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJ9YTQR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nHao9x48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQ0yKQQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZSVVXgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYcAtId8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IV4IjYluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mMA3FlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XY0EqRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kaHY6UuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func su11GWvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrgHYHwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKWcGZAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yz9r8XV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VlJRIg2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1HYT6DvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjm5Xg1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKAZRHXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2gsdtjKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmiA9aJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnCWmtPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwwctuEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVtQAhKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkZr6k9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7CDVQvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4913d6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0k35hb1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2iuN7OoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPTQQLYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rU0mJuJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJNmeMZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOet2JBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSRVWiLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3QQN3VrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzcKd1dTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nklzX8LYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IVLNaxapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjWOj3jCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wdwc7Ug6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2ictafjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7w1SCnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiYp7lboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlPQRoBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVva8L27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YorSf0UTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AV6JAaZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

