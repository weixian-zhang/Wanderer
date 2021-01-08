package main

type UserCaseRunner interface {
	RunUseCase(done chan bool)
}