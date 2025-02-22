// Copyright 2023-2025 The Jule Programming Language.
// Use of this source code is governed by a BSD 3-Clause
// license that can be found in the LICENSE file.

#ifndef __JULE_EXCEPTIONAL_HPP
#define __JULE_EXCEPTIONAL_HPP

#include "any.hpp"

namespace jule
{
    // Wrapper structure for Jule's void exceptionals.
    class VoidExceptional
    {
    public:
        jule::Any error;

        // Reports whether no exception.
        bool ok(void) const noexcept
        {
            return this->error == nullptr;
        }
    };

    // Wrapper structure for Jule's exceptionals.
    template <typename T>
    class Exceptional
    {
    public:
        jule::Any error;
        T result;

        // Reports whether no exception.
        bool ok(void) const noexcept
        {
            return this->error == nullptr;
        }
    };
} // namespace jule

#endif // __JULE_EXCEPTIONAL_HPP
