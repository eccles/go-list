// .*@mycompany\.com MY COMPANY 2025
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package list

// Conflate removes consecutive duplicate elements from a slice if the number
// of consecutive duplicates exceeds the specified limit. Elements are kept
// if their consecutive count is less than or equal to the limit.
func Conflate[T comparable](slice []T, limit int) []T {
	if len(slice) == 0 {
		return nil
	}

	result := make([]T, 0, len(slice))

	for i := 0; i < len(slice); {
		count := countConsecutive(slice, i)

		if count <= limit {
			result = append(result, slice[i:i+count]...)
		}

		i += count
	}

	return result
}

func countConsecutive[T comparable](slice []T, start int) int {
	count := 1
	for j := start + 1; j < len(slice) && slice[start] == slice[j]; j++ {
		count++
	}
	return count
}
