package noc

type FilterFunc func(next HandlerFunc) HandlerFunc
