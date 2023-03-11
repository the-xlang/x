package sema

type _Lookup interface {
	find_package(ident string) *Package
	find_var(ident string, cpp_linked bool) *Var
	find_type_alias(ident string, cpp_linked bool) *TypeAlias
	find__struct(ident string, cpp_linked bool) *Struct
	find_fn(ident string, cpp_linked bool)
	find_trait(ident string) *Trait
	find_enums(ident string) *Enum
}
