package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.9" )

func 1ONqeYPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5P1F8dADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMbX9BFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CynAZsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWR9s28mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1kjm4v8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAgtE4ngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stNmwZMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2yFMQ3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MOP5GUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gv6INrdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLA0y6wYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYVlMtNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7NYMnrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjrwoeYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCITfhHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNgrGdsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func avASj26SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rv0o6kF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4Jgui5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4oFCqRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEzS2aTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpuljM3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdfrZBnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZs4jcPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vdFnpZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RpaniFM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbDsN1VxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4QfChNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pBjV2SJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnr54ifwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aILge1YLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hV1SpTw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEllACsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCs9V1j3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iw1MwSnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UC9RErH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9xqtLxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IxUTNycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6ntYttTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcyQ5WPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8LquZmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sI7ajKk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXpMDgQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocpfZDH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AwhHDgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auLty4fdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAmfP8XWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jsx5fKn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cna9ianyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPdjAFhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1ACglVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3nrEiICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1xu3fwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qb1xH8rMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPfoyWmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FIxUuw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5kRwCBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oy1Z6Q4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGqZg1egWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ri9LyLIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oes59lnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0liNvBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E76nszdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H137sHdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1W6f7kQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diLQKLxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wcbPXeDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func am4t37iwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QkpMPBQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLG0EHvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8Tf1XftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCRYT0WBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsWgWY2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNTRtjlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jr0Et0YiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CaREJ72MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7MLrKjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMOloDUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFe68O7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8v2qqshuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VevH96epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnYMBuqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jy7KYKb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPVDdW15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99EbaRM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCXiIZnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6PTJ9mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQFfAwHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uaTRDFN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckrbWxdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5YPinX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEmnqIOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCTBl9EHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fw90ToVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xqj2qwttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APw7qXbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIUbyUffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1ZhZymjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JQdD3BxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjh02JyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ovr2z4OcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSmBr2duWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPc7ckEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZQZNGkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hz3qGgNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lp571wjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9rF28M1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbsGqc1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pjsp5NrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isPPs0VEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7KHmpb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mf6ol0ClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05jN6o6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFJuohAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53aEpCAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYqy9BG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUDJGOG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExsigMUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ajwMtUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50HXqynAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EDVCV3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNHhGAgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6RVbxOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ie1P7tnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PN5KmBLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AppKg9fDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVlI3k1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGbTEoBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbHV7JmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zmg150FWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPUdRNjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKce8G7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ej8J2opOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2nmb0RFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Q5mVjN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZmzAJKDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1mHmZHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvPa38nmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Ubak40KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GG9ASfmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDjgmijUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9w9hKhGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7OFAAQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5Lji4LWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29hDMKZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFSI9LGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nU92N0epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aiug5hwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNU3SJSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hl0DbQ2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmOnpVSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJpvmnh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsSxiF3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPB0iqjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yM4FewvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KggXeujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzybH4L4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 42tFSiPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z50HnaOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkXMFHbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhvgDZTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdihDC75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrWaM48AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaM2b5FsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4i1exwUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ew0qPM9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRCnSTDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3vaoC4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ax2NABxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X97HVhlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JHkhOBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JO1gzWVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3eC8dPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhEehR2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svW0FHKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iz5gtQ3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fXjDFpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IP9xocdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hs41e3AfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0nYrxAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B626DKYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZf7pcEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5InJlIW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYKUS0loWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqgT6p81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6GNBao5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYLiicqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BV4hUOSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mzyq26yQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qe4f8i2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNj2wYyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHSrqKyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGow6lIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thZqcRE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zl4EfEluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T63PWC9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GoZmxq6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvfyC090Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuxp5fXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0no6K0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFTtGS1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUoojOiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgqI658XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNhP9GvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DFy2ZEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hO91tnKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kc8yQyFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d05enstyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSwffZm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjN6aqGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKpDoyYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgbX1mCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlQtKx7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lH9wcUZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbkXZJmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1zgxGsmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rj5pAyoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHWF3JSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOeo3nqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mFBc35FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0crx78dKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gf65x46hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGSli6inWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIA0bb51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15HI3rScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eM7W2ZSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ckMXu4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAMhUgbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5S79kCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljJAlUvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJvv0hIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wabSdROaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugoQTUJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzTqJfUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqQSSNycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkXtQCFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BU4e8ZzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJoZ7OvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzuAJTvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsm4QIQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYsULwzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k50nx0ojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0LuLb0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KnoD2cBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFT34EJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lwgV2HaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDsNI7ChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ita8uWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBvQoChGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCLUL7vSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bN206qDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIYTvc96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXhkfPCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHXlxJYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZfsQbb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LL4ztFHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzBcTa0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhsezdKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K6zXLBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCF6zkjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3hU0ZGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfJZDDerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQEF2TGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XCGry1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFg1jRdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKZy77vbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XxqMUeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OU7E0ByJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOfcW4KBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeU7XuDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpdyR9aPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEOze4xxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nGeDZDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUOKe9mwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2zMyqz7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mjs2MvyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySPrzORhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARD30R44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0DVr4bKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j34JB6rQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhzgjOHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALnVmgPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbiBNN1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVkvo45xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z40PpOP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIKAGpYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZWaiIj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EdpZBhdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAA703PNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OQmeh2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vh9pVnNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjvQ2pRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYgAt06JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkN32AGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLQQjP5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmxOyVWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uez8JCweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPNiHtSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3LrDlUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqVBfHuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwEItbiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjQoKaxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIEgD36qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmJCrog5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDLWT0RPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzhkGEImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjCkmD3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1tb3PRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAsjy6p5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gt2umMVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dvdddm5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53Ip7L3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwQdVzaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5durgyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYmQL91cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUodkC8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zj47o5soWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAMyDma3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImK4JbKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buWOmyOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sj9G7FiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjnxLPBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLz2oPhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNnNybCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHePWcCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3QCbObxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEAaZHujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcaI1sMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDvPS5evWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9MBPK1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gO5I750Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDPp5DpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmhYtfWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGlrfFyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSAZ5XyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jovn6iwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UHSjRf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NqKFmBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CjoypvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JKiGedEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ux9PcRWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bnN43XIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTPcFCu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkZ3l4SoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4B5tG8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUWjqh01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q69Jy252Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifR6FZvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zy5UGKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLEkR5xoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJ3UCxfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmpcROiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYXugXlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gh8beYfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CjXURo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21P6GgyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRDcnM0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpPx4yHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfyLNf8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJqNEJT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gt8taqxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhCpM3lJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0mqYAqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMdc67hGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WwTGVC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rE1BLmtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eaH0MdknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCfldZvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fc12CNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4r4gT5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJzAQuslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqWQV8RqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItcaJH8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeFegjc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adxM9bLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YMwu7oTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djSgraheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rbj0kYadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUWSvYPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJodmYZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbc38eDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uf8oj7MHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hEUN1G9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cpe5terUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3Pxar2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bU1SrYDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4WdpgmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhkcHXwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6RPynBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7WTwq0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ER9Qyah2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNicsuGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AMwl7n6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqKgIlN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooXR64weWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnByPlNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0o9LBruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eyL8BVTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xafp0JI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrYwO4H3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4BY2atYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4Wj0q3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUj4B7EXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLGAOPghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XkWBIK6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkJjaGuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgckoKEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CKcgPo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFOj2URdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gN44I4pzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wxcO7krWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbcyXiCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HwLXGyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnO7WchvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmVr8BvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHwK13r6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CabZx6XdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wSqCxDxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORMGKGlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8sC60jEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZmmMoj79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5xOOLbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbtzCXa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZ8949bRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QS4HPNypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THy8fCUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEf87ht2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAXE0DybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YiLmi2aUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGc1Qea5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBmuheGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2QifL1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtxLa54wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjyDvdk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTnscNZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kOLxp5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ih5g3JUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func phJf77ufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnMzZ749Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jEpAWrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ey6HJU3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MIUzI82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPyzJgSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Ees3cJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrToBeqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIw0aHtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfjpbvAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZSAOkhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89BV330GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NF6CKajbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71ubyX4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dq1Own5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNf3szZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u73jazsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEOQYPpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFiHPOHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnHcJVvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUPKthePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5boMJTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoIInHpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5bDpdTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFVuzZmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmb6y264Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epN0QkbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GoZ4kayoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjIMx4TyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func of59sHFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRzUTislWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfLUMcCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6d4kLp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfejEKmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xCx7BRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEeceHViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZW0Sk9xmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynWaG7BQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGFI0ePaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUKKezT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAP6VCjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONBWhiwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YBGzlAPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfkTpoLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFOZT5fqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pjjKO5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1s4BQzDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFxQH8ILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AG2qm2tUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ALQ57LuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSFqmsYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0gtBUKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdpansw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEafVgUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlYemN1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w70yF5MgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugOhxD4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itXakBblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfv1bEpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3SJyGYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wnVaBV0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ig5Taq9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tKVCZd3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3Z58JUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4rJ5US2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdBcXz0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIEqwPjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Nn43BGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tMmDcwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYkx0xjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQFkpnVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EXuIYsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VW1rOyqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tl6OBz0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQgushkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hr88QtcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzbwqBipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AKRrkMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7V79Jk62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7U9NED0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOUHSzQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func No1jhDrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYtybLCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iarrfUP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func El5eFyFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLd95gbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeLr4QoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5awNfNtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQtryDDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxpvVSMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtAYaaCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcgtg2sBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JvAa3yW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5q2F8trpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UB10o8EvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DuWUM6gPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGtgAyzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nt3WqOu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UbrGyYORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZ3zd3gpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJdfJfTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ibiKtKE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func votio60XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfFhhEh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrKs4lnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5gh3blaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Y2BVnu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWRah4KqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nM2PS0URWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btml4CCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aW4EtHRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90Bomdq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Zoabp0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9X6IrXS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eL4RNypnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKUJtDdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12jjPHjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJf1Y4cSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsCptwniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1BHOYjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbi3qjurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXzmmL6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flKLqgCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvM1Ee6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAaPm4egWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAPExMkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYqMpjauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSnsVOWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFndBjv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

