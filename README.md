# Go Race Condition Example

This repository demonstrates a common race condition in Go that can occur when multiple goroutines access and modify a shared variable without proper synchronization. The program uses several goroutines that increment a shared counter; without synchronization, this leads to an incorrect result. The solution involves using a mutex to protect access to the shared variable.