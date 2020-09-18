package main

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	s := "anagram"
	t := "nagaram"
	isAnagram(s, t)
}

func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)     // 初始化一个名为alphabet的int类型切片, 长度为26
	sBytes := []byte(s)             // s 转换为byte的切片
	tBytes := []byte(t)             // t 转换为byte的切片
	if len(sBytes) != len(tBytes) { // 判断两个切片的长度是否相同,不相等返回false
		return false
	}
	for i := 0; i < len(sBytes); i++ { // 循环遍历每个字母,然后通过sBytes[i]-'a'得到字母位置,然后记录进alphabet切片
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {// 循环遍历每个字母,然后通过tBytes[i]-'a'得到字母位置,然后对alphabet切片进行减一
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 { // 判断alphabet的每个字母的value是否为0,不为了返回false
			return false
		}
	}
	return true // 两个字符串是字母异位词
}
