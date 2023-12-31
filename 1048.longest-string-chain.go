/*
 * @lc app=leetcode id=1048 lang=golang
 *
 * [1048] Longest String Chain
 */

package main

// @lc code=start
func longestStrChain(words []string) int {
	wordArr := [17]map[string]int{}
	for _, word := range words {
		wordLen := len(word)
		if wordArr[wordLen] == nil {
			wordArr[wordLen] = map[string]int{}
		}
		wordArr[wordLen][word] = 1
	}

	max := 1

	for i := 16; i > 0; i-- {
		if wordArr[i] == nil || i-1 < 1 || wordArr[i-1] == nil {
			continue
		}
		for word, numPredecessor := range wordArr[i] {
			potentialPred := []string{}
			for i := 0; i < len(word); i++ {
				w := word[:i] + word[i+1:]
				potentialPred = append(potentialPred, w)
			}

			for _, pred := range potentialPred {
				if predChain, ok := wordArr[i-1][pred]; ok {
					var v int
					if numPredecessor+1 > predChain {
						v = numPredecessor + 1
					} else {
						v = predChain
					}
					wordArr[i-1][pred] = v
					if v > max {
						max = v
					}
				}
			}
		}
	}

	return max
}

// @lc code=end
// func main() {
// 	words := []string{"a", "b", "ba", "bca", "bda", "bdca"}
// 	fmt.Println(longestStrChain(words))
// 	words = []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}
// 	fmt.Println(longestStrChain(words))
// 	words = []string{"abcd", "dbqca"}
// 	fmt.Println(longestStrChain(words))
//
// 	words = []string{"biltnzk", "jxwakrfxsifoj", "uzdwyaxvcsr", "sqqgkhwbf", "tnoftkolx", "ipmtvxcwe", "zsucxrqkhahuo", "qngglugvm", "kvohqyedig", "njoxacsnddwrg", "vwtnxw", "kjjourlrzpgeem", "xcs", "pfsgimurs", "lsifyg", "uzwyxcsr", "muzdwcyanxvcstr", "teqyrlhbvcv", "rkga", "tudezgzbnzb", "uzwyaxvcsr", "qvzkmgfulby", "x", "muzdwcyianxvcstr", "koqyig", "gl", "aqcacmy", "pmvwe", "eskofqduddkhykr", "pm", "saxxd", "ds", "iemm", "tudegzbz", "yipsawmxbp", "qyrlhbvcv", "yxuhwkzvoczoz", "zsucxqkahuo", "kga", "zwziivbijeiig", "wffaheemjnjahzdd", "zcxkahuo", "djjjsulms", "plxh", "ffpasoizwhtu", "zwziivijeii", "fyvpzegautteiv", "qszaitzfzv", "uwoghcy", "qqgkhwbf", "eteqyrllhbvcvg", "qknspkhngorof", "qwvzkmgfuljbyz", "grkte", "grikrnwezryi", "xjbpvekneaxn", "cy", "wnhnyqmpbsum", "m", "offqllgj", "plxhib", "omblqcoktkyf", "pasw", "prsngzx", "offlj", "rvvudgpixa", "djjjjsulmmrs", "gt", "mpfsgimurs", "cxkahuo", "ipmtvxcwue", "pqrbaoquxqemv", "prqqv", "tnoftfkolx", "jfzzaw", "rshquwmrboghccy", "ebqhvwewzzmqif", "rrd", "dvjjjjqsulmmrs", "pfsiurs", "crnruydj", "rvqgeqql", "djsums", "prbaquqemv", "bs", "dzytccvny", "kce", "llfv", "jfzaw", "qwvzkmgbfuljbyz", "kgieph", "hnympsum", "ewv", "vfgel", "rklga", "llzqbfv", "gte", "jckqurkg", "qngglugm", "tudgzbz", "ipmvcwe", "rr", "kkcev", "djjjjsulmrs", "llqbfv", "offqlgj", "paswu", "tlrlcnnrsrf", "jcckqurkg", "jjourlpgeem", "nvl", "shquwmrboghccy", "vncfgelm", "dgcdgjcksk", "vvhvmibflb", "juifgeqkaectlcj", "scvdl", "whcy", "yipswmbp", "wcy", "hbqq", "bsth", "etjurltvpsuy", "dzvytcccevnceyq", "apqrbaoquxqemv", "kvohuqyediyig", "lenybbukzftz", "ffpasoiuztwhtu", "lzlhzqibfv", "wfeemjnjahzdd", "djsulms", "xtudezgzbnzb", "eemjhzdd", "scavdil", "guchrvaqbe", "nvll", "sxzfpzjmxvu", "dytccvny", "grikrnjwezryi", "prng", "ntvmcwwpzo", "laqgcacyxmym", "mglosifyg", "nynvlqll", "vwtn", "lh", "zhhxducgelhy", "prg", "kghierph", "zsucxrqkhahuom", "kvohqydig", "eemjhzd", "offiqcdllgji", "dyc", "toflx", "dzvytccvney", "ghvb", "to", "guchrvab", "wyimthhfzndppwt", "elbqhvwewzzmqif", "hkghiyerph", "hkghiyejrph", "hlsioorugbsuu", "c", "kgierph", "bstbghj", "prbquqev", "mpfsdgimurs", "zfpjvu", "zfpvu", "yxuhwkzvoczfgoz", "gel", "ntvmcpzo", "ekofqduddkhykr", "ekofqdddhykr", "rqeql", "nhnympsum", "xhoqlfolk", "ipmtvxcwuje", "wgmhjhdmnqot", "bsh", "rvncfgelm", "hkahpbb", "lzlzqibfv", "xoqlfok", "tnoftfkogwgplx", "ekofqdddkhykr", "zwiieii", "ujfzzaw", "jfzw", "djsms", "scavdpilj", "tnoftfkoglx", "ps", "vwtnw", "scavhdpilj", "scayvhdpuilji", "pdrshqngzx", "crnrud", "wmhjhdmnqot", "wghmhjhdmnqot", "vbyipsawmxbp", "qknsapkhngorof", "wymthhfzndppwt", "wxcs", "dzvytccevney", "acacmy", "dycy", "teqyrllhbvcv", "uzwyxcs", "wmhjhdmnqt", "qvzkmgfulbyz", "qngglum", "zhhxgdyukcgelhy", "oj", "iljes", "bstbh", "laqcacxmy", "tofx", "ke", "yivkqoek", "djjjsulmrs", "lbirdzvttzze", "l", "zhhxgdukcgelhy", "grikvrnjwezryi", "bltz", "npynvlqll", "gvb", "okzrs", "urbarfkmnlxxn", "qsyzaixtzfazv", "dytcy", "h", "kohqyig", "hgri", "ojdxm", "ujfdfzzaw", "qyrhbvcv", "ebqhvwewzmqif", "uzwxcs", "lebzf", "ysijvkwqmoekromh", "wffaeemjnjahzdd", "crnrduyndj", "ujfdmfzzaw", "laqgcacyxmzgym", "jjourlrpgeem", "kvohqyediyig", "lebukzf", "zwiijeii", "guchrvb", "omoktkyf", "hpgt", "yikoek", "ysijvkwqoekromh", "tvpo", "ysijvkqoekromh", "xbgq", "d", "abmtk", "ors", "rnrd", "xzrugvlzduaxhzc", "njoxacjsnddwrg", "yipswmxbp", "xqsyzaixtzfazv", "urbrfknlxxn", "sxzfpjxvu", "prbaquxqemv", "dvjjjjsulmmrs", "kviahvqu", "urbfknx", "qvmgfulby", "yikqoek", "zsucxrqkhfahuomm", "koqyg", "djss", "moxpfsdgimlurs", "qeql", "urbrfknlxn", "kgieh", "qnspkhngorof", "plxyhib", "scyayvhdpuiljki", "vvhvmbflb", "lpzluhzqxibfv", "kkcbev", "hpzgty", "nyvlqll", "kvahvu", "rklgja", "ipmtavxcwuje", "lbirdzvvttzze", "psw", "fpasoiwhtu", "dgcdgjckk", "qknhsapkhngorof", "qszaixtzfazv", "tvp", "abmtvk", "uwrboghcy", "hbq", "crnruyd", "etjurltvsuy", "etjurltyvpsuy", "lenbukzf", "teqyrllhbvcvg", "ipmvwe", "o", "crnryduyndj", "lbirdzvvqfttzze", "tnoftfkowglx", "ipmtavxcwujre", "omlcoktkyf", "rnperyemtmqh", "bltnzk", "sxzfpzjxvu", "uzdwyaxvcstr", "bq", "rvvugpixa", "laqcacxmym", "wffeemjnjahzdd", "fpvu", "xjbpvekngeyaxbn", "dzvytccevncey", "qgly", "scavdl", "fw", "tox", "toftklx", "prbaoquxqemv", "ztrobzqiukdkcbv", "yivkqoekr", "feemjnjhzdd", "plxhi", "cp", "fyvpzgauttei", "prshqngzx", "kplxyrhib", "suwrboghcy", "kviahvu", "mvwe", "dzvytccvny", "hbqwq", "prbquqemv", "lzlhzqxibfv", "ll", "omblcoktkyf", "toftlx", "lpzlhzqxibfv", "tudegzbnz", "ddgcdgjcgkspk", "kgih", "xjbpvekneaxbn", "suwrboghccy", "zwiiijeii", "dytccy", "ympsum", "jxwakfxsifoj", "uwhcy", "yxuhwkzvoczfoz", "xzfpjvu", "lenybbukzft", "b", "llqfv", "laqgcacyxmgym", "xq", "scavdilj", "zwziivbijaeiig", "scyayvhdpuilji", "amvevfulhsd", "dss", "tlrlcnnrs", "uzwyaxcsr", "qspkhngorof", "etjurtvsuy", "wgqhmhjhhdmnqot", "tvmpo", "tnoftklx", "qgflby", "mlosifyg", "oqyg", "gchvb", "t", "offqcdllgj", "ziieii", "zwziivbijeii", "vp", "lpb", "fyvprzegauttejiv", "vtn", "amefulhsd", "llf", "muzdwyaxvcstr", "zucxqkahuo", "pfsgiurs", "obstbghj", "ipmqtavxcwuzjrbe", "djjsulms", "qvmgflby", "ljpzluhzqxibfv", "jjourlrzpgeem", "zrugvlduaxhzc", "xbpvkneaxn", "ljpzluhzgqyxibfv", "yivkqoekrh", "laqcacyxmym", "nyvll", "muzdwcyaxvcstr", "fyvpzegauttejiv", "offlgj", "vnfgelm", "eteiqyrllhbvcvg", "zsucxrqkhahuomm", "ibiltnzk", "rklgjae", "fpasoizwhtu", "t", "zhhxdukcgelhy", "fpasoiwu", "xzfpjxvu", "tlrlcnnrysrf", "ojx", "mpum", "lxh", "eturtvsuy", "rklgbjaae", "kahpbb", "qngglugmfvmp", "fielbqtcri", "xzruogvlzduaxhzc", "rshquwmrbtoghccy", "nyvlll", "lbirdzvvqttzze", "dgcdgjckspk", "vvhvmibfilb", "dzvytcccevncey", "g", "vwe", "zwxcs", "k", "jourlpgeem", "cpk", "cds", "tlrlcnnrsr", "ivemm", "fgel", "grktse", "urbfknlxn", "qwvzkmgfulbyz", "xjbpvekngeaxbn", "wphuutlgczfspyga", "xbq", "offqcdllgji", "vbyipsakwmxbp", "qyrhbvc", "ygzpztbno", "xhogqlfolk", "ujffzzaw", "xbnmgq", "uwohcy", "rnperyemqh", "prbqqev", "lenybukzf", "mxpfsdgimurs", "ga", "hpt", "moxpfsdgimurs", "vb", "offqcllgj", "rklgbjae", "lifg", "ztrobzzqiukdkcbv", "xoqok", "cs", "snaxxd", "cdds", "qknhsapkhngorohf", "rvqgeql", "rnperyemmqh", "scavhdpuilji", "urbfknlx", "rvvugixa", "ygzpztbndon", "zrugvlzduaxhzc", "shuwmrboghccy", "mlsifyg", "xhoqlfok", "wfeemjnjhzdd", "lbzf", "wythhfzndppwt", "mglqosifyg", "ojxm", "kvohuqyevdiyig", "grte", "prsngz", "eteeiqyrllhbvcvg", "dytccny", "qngglugfvmp", "kohqydig", "fu", "qgfly", "tvmcpzo", "tnoftfkowgplx", "zruglduaxzc", "yijvkqoekrh", "xqsyzaixtzfdazv", "ipmqtavxcwuzjre", "omloktkyf", "ympum", "lzlzqbfv", "pasowu", "rvqeql", "qngglugvmp", "hkghierph", "eemjhz", "feemnjhzdd", "c", "yxpuhwkzvoczfgoz", "dgcgjckk", "lbz", "yxuwkzvoczoz", "zrugvlduaxzc", "ntvmcwpzo", "fzw", "ygzpmztbndon", "rvncfgxelm", "mpm", "tudezgzbnz", "bltzk", "ffpasoiuzwhtu", "cd", "r", "okrs", "byipsawmxbp", "prsqngzx", "wnhnyqmpsum", "ipmqtavxcwujre", "w", "fpasoiwtu", "plxyrhib", "bstbhj", "xbnmrgq", "ipmtvcwe", "urbfkn", "nympsum", "qtngglugmfvmpt", "jckqurg", "hgr", "hpzgt", "rvvxudgpixa", "ysijvkqoekrh", "lebkzf", "guchvb", "kvohqyediyg", "amvefulhsd", "suwmrboghccy", "fvu", "ibdiltnzk", "rnrud", "iem", "urbarfknlxxn", "ygzpztbnon", "prsng", "zcxqkahuo", "ffpeasoiuztwhtu", "laqcacmy", "qszaitzfazv", "xbngq", "qvkmgfulby", "scavhdpuilj", "zsucxrqkahuo", "v", "qtngglugmfvmp", "ysijvkqoekrmh", "lfg", "prqqev", "pasoiwu", "p", "tvmcpo", "kcev", "im", "crnrduydj", "vfgelm", "ddgcdgjckspk", "ivqemm", "ljpzluhzgqxibfv", "lenybukzft", "nhnyqmpsum", "iljesr", "hp", "tqyrlhbvcv", "eemnjhzdd", "xbpvekneaxn", "wghmhjhhdmnqot", "uwboghcy", "guchrvabe", "xoqfok", "fyvpzgautteiv", "pg", "zwiivijeii", "qvgflby", "lsifg"}
// 	fmt.Println(longestStrChain(words))
// }
