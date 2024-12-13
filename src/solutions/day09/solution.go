package day09

import (
	"aoc/src/utils"
	"slices"
	"strconv"
	"strings"
)

type File struct {
	id   string
	size int
}

func DiskMapToString(diskMap []File) string {
	result := ""
	for _, file := range diskMap {
		for i := 0; i < file.size; i++ {
			result += file.id
		}
	}

	return result
}

func convertToDiskMap(diskMapRepresentation string, compactMode bool) []File {
	diskMap := []File{}

	for i, elm := range strings.Split(diskMapRepresentation, "") {
		block, _ := strconv.Atoi(elm)

		if block == 0 {
			continue
		}

		var char string
		if i%2 == 0 {
			char = strconv.Itoa(i / 2)
		} else {
			char = "."
		}

		if compactMode {
			diskMap = append(diskMap, File{id: char, size: block})
		} else {
			for j := 0; j < block; j++ {
				diskMap = append(diskMap, File{id: char, size: 1})
			}
		}
	}

	return diskMap
}

func defragment(diskMap []File) {
	for i, j := 0, len(diskMap)-1; i <= j; {
		if diskMap[i].id == "." && diskMap[j].id != "." {
			diskMap[i], diskMap[j] = diskMap[j], diskMap[i]
			i++
			j--
		} else {
			if diskMap[i].id != "." {
				i++
			} else if diskMap[j].id != "." {
				j--
			} else if diskMap[i].id == "." && diskMap[j].id == "." {
				j--
			}
		}
	}
}

func defragmentV2(diskMap []File) []File {
	i := len(diskMap) - 1
	for i >= 0 {
		if diskMap[i].id == "." {
			i--
		}

		for j := 0; j <= i; j++ {
			if diskMap[j].id != "." || diskMap[j].size < diskMap[i].size {
				continue
			}

			fileToInsert := File{id: diskMap[i].id, size: diskMap[i].size}
			diskMap[j].size -= diskMap[i].size
			diskMap[i].id = "."
			diskMap = slices.Insert(diskMap, j, fileToInsert)

			k := 0
			for k < i {
				if diskMap[k].size == 0 {
					diskMap = slices.Delete(diskMap, k, k+1)

					if k < i {
						i--
					}
					continue
				}

				if diskMap[k].id == "." && diskMap[k+1].id == "." {
					diskMap[k].size += diskMap[k+1].size
					diskMap = slices.Delete(diskMap, k+1, k+2)
					if k < i {
						i--
					}
					continue
				}
				k++
			}
			i++
			break
		}

		i--
	}

	return diskMap
}

func calculateChecksum(diskMap []File, compactMode bool) int {
	checksum := 0
	x := 0
	for i, file := range diskMap {
		if file.id == "." {
			x += file.size
			continue
		}

		if compactMode {
			for j := 0; j < file.size; j++ {
				num, _ := strconv.Atoi(file.id)
				checksum += (x * num)
				x++
			}
		} else {
			num, _ := strconv.Atoi(file.id)
			checksum += (i * num)
		}
	}
	return checksum
}

func SolveProblem1() (string, error) {
	diskMapRepresentation := utils.GetInput(9)

	diskMap := convertToDiskMap(diskMapRepresentation, false)

	defragment(diskMap)

	result := calculateChecksum(diskMap, false)

	return strconv.Itoa(result), nil
}

func SolveProblem2() (string, error) {
	diskMapRepresentation := utils.GetInput(9)

	diskMap := convertToDiskMap(diskMapRepresentation, true)

	diskMap = defragmentV2(diskMap)

	result := calculateChecksum(diskMap, true)

	return strconv.Itoa(result), nil
}
