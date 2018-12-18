# gobisect package

[![Build Status](https://travis-ci.com/UrNas/gobisect.svg?branch=master)](https://travis-ci.com/UrNas/gobisect) [![Go Report Card](https://goreportcard.com/badge/github.com/UrNas/gobisect)](https://goreportcard.com/report/github.com/UrNas/gobisect)

### Bisect Algorithm Functions in Go

The purpose of Bisect algorithm is to find a position in slice where an element needs to be inserted to keep the slice sorted.

gobisect package provides the bisect algorithms which allows to keep the slice in sorted order after insertion of each element. This is essential as this reduces overhead time required to sort the slice again and again after insertion of each element.


