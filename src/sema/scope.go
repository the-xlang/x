package sema

import "github.com/julelang/jule/ast"

// Statement type.
type St = any

// Scope.
type Scope struct {
	Parent   *Scope
	Unsafety bool
	Deferred bool
	Stmts    []St
}

// Scope checker.
type _ScopeChecker struct {
	s      *_Sema
	parent *_ScopeChecker
	table  *SymbolTable
	scope  *Scope
	tree   *ast.ScopeTree
}

// Returns package by identifier.
// Returns nil if not exist any package in this identifier.
//
// Lookups:
//  - Sema.
func (sc *_ScopeChecker) Find_package(ident string) *Package {
	return sc.s.Find_package(ident)
}

// Returns package by selector.
// Returns nil if selector returns false for all packages.
// Returns nil if selector is nil.
//
// Lookups:
//  - Sema.
func (sc *_ScopeChecker) Select_package(selector func(*Package) bool) *Package {
	return sc.s.Select_package(selector)
}

// Returns variable by identifier and cpp linked state.
// Returns nil if not exist any variable in this identifier.
//
// Lookups:
//  - Current scope.
//  - Parent scopes.
//  - Sema.
func (sc *_ScopeChecker) Find_var(ident string, cpp_linked bool) *Var {
	v := sc.table.Find_var(ident, cpp_linked)
	if v != nil {
		return v
	}

	parent := sc.parent
	for parent != nil {
		v := parent.table.Find_var(ident, cpp_linked)
		if v != nil {
			return v
		}
		parent = parent.parent
	}

	return sc.s.Find_var(ident, cpp_linked)
}

// Returns type alias by identifier and cpp linked state.
// Returns nil if not exist any type alias in this identifier.
//
// Lookups:
//  - Current scope.
//  - Parent scopes.
//  - Sema.
func (sc *_ScopeChecker) Find_type_alias(ident string, cpp_linked bool) *TypeAlias {
	ta := sc.table.Find_type_alias(ident, cpp_linked)
	if ta != nil {
		return ta
	}

	parent := sc.parent
	for parent != nil {
		ta := parent.table.Find_type_alias(ident, cpp_linked)
		if ta != nil {
			return ta
		}
		parent = parent.parent
	}

	return sc.s.Find_type_alias(ident, cpp_linked)
}

// Returns struct by identifier and cpp linked state.
// Returns nil if not exist any struct in this identifier.
//
// Lookups:
//  - Sema.
func (sc *_ScopeChecker) Find_struct(ident string, cpp_linked bool) *Struct {
	return sc.s.Find_struct(ident, cpp_linked)
}

// Returns function by identifier and cpp linked state.
// Returns nil if not exist any function in this identifier.
//
// Lookups:
//  - Sema.
func (sc *_ScopeChecker) Find_fn(ident string, cpp_linked bool) *Fn {
	return sc.s.Find_fn(ident, cpp_linked)
}

// Returns trait by identifier.
// Returns nil if not exist any trait in this identifier.
//
// Lookups:
//  - Sema.
func (sc *_ScopeChecker) Find_trait(ident string) *Trait {
	return sc.s.Find_trait(ident)
}

// Returns enum by identifier.
// Returns nil if not exist any enum in this identifier.
//
// Lookups:
//  - Sema.
func (sc *_ScopeChecker) Find_enum(ident string) *Enum {
	return sc.s.Find_enum(ident)
}

// Checks scope tree.
func (sc *_ScopeChecker) check(tree *ast.ScopeTree, s *Scope) {
	sc.tree = tree
	sc.scope = s
}

func new_scope_checker(s *_Sema) *_ScopeChecker {
	return &_ScopeChecker{
		s: s,
	}
}
