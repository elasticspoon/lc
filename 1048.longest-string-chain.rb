#
# @lc app=leetcode id=1048 lang=ruby
#
# [1048] Longest String Chain
#

# @lc code=start
# @param {String[]} words
# @return {Integer}

# Word = Struct.new(:word, :preds) do
#   def initialize(word)
#     super(word, [])
#   end
#
#   def potential_preds
#     return @pot_preds unless @pot_preds.nil?
#
#     @pot_preds = (0...word.length).map do |i|
#       word[0...i] + word[i + 1..]
#     end
#   end
# end
#
# def longest_str_chain(words)
#   words_mapped = Array.new(16)
#
#   words.each do |word|
#     (words_mapped[word.length] ||= []).push(Word.new(word))
#   end
#
#   words_mapped.reverse_each.with_index do |v, i|
#     v.each do |word|
#       puts word if words_mapped[i - 1]&.contain?(word)
#     end
#   end
# end
def longest_str_chain(words)
  words_mapped = []
  words.each do |word|
    (words_mapped[word.length] ||= {})[word] = 1
  end

  max = 1
  (words_mapped.length - 1).downto(1) do |i|
    puts i
    words_mapped[i]&.each do |word, num_preds|
      pot_preds = (0...word.length).map do |j|
        word[0...j] + word[j + 1..]
      end

      pot_preds.each do |w|
        if words_mapped[i - 1]&.has_key?(w)
          x = [words_mapped[i - 1][w], num_preds + 1].max
          words_mapped[i - 1][w] = x
          max = x if x > max
        end
      end
    end
  end

  max
end
# @lc code=end

# some more complex test cases
# puts longest_str_chain(["a", "b", "ba", "bca", "bda", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca"]) == 4
# puts longest_str_chain(["a", "b", "ba", "bca", "bda", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "bdca", "abcde", "abcde", "abcde", "abcde", "abcde", "abcde", "abcde", "abcde", "abcde", "abcde"]) == 4
#
# # some test cases from leetcode
# puts longest_str_chain(["a", "b", "ba", "bca", "bda", "bdca"]) == 4
# puts longest_str_chain(["xbc", "pcxbcf", "xb", "cxbc", "pcxbc"]) == 5
# puts longest_str_chain(["abcd", "dbqca"]) == 1
words = ["biltnzk", "jxwakrfxsifoj", "uzdwyaxvcsr", "sqqgkhwbf", "tnoftkolx", "ipmtvxcwe", "zsucxrqkhahuo", "qngglugvm", "kvohqyedig", "njoxacsnddwrg", "vwtnxw", "kjjourlrzpgeem", "xcs", "pfsgimurs", "lsifyg", "uzwyxcsr", "muzdwcyanxvcstr", "teqyrlhbvcv", "rkga", "tudezgzbnzb", "uzwyaxvcsr", "qvzkmgfulby", "x", "muzdwcyianxvcstr", "koqyig", "gl", "aqcacmy", "pmvwe", "eskofqduddkhykr", "pm", "saxxd", "ds", "iemm", "tudegzbz", "yipsawmxbp", "qyrlhbvcv", "yxuhwkzvoczoz", "zsucxqkahuo", "kga", "zwziivbijeiig", "wffaheemjnjahzdd", "zcxkahuo", "djjjsulms", "plxh", "ffpasoizwhtu", "zwziivijeii", "fyvpzegautteiv", "qszaitzfzv", "uwoghcy", "qqgkhwbf", "eteqyrllhbvcvg", "qknspkhngorof", "qwvzkmgfuljbyz", "grkte", "grikrnwezryi", "xjbpvekneaxn", "cy", "wnhnyqmpbsum", "m", "offqllgj", "plxhib", "omblqcoktkyf", "pasw", "prsngzx", "offlj", "rvvudgpixa", "djjjjsulmmrs", "gt", "mpfsgimurs", "cxkahuo", "ipmtvxcwue", "pqrbaoquxqemv", "prqqv", "tnoftfkolx", "jfzzaw", "rshquwmrboghccy", "ebqhvwewzzmqif", "rrd", "dvjjjjqsulmmrs", "pfsiurs", "crnruydj", "rvqgeqql", "djsums", "prbaquqemv", "bs", "dzytccvny", "kce", "llfv", "jfzaw", "qwvzkmgbfuljbyz", "kgieph", "hnympsum", "ewv", "vfgel", "rklga", "llzqbfv", "gte", "jckqurkg", "qngglugm", "tudgzbz", "ipmvcwe", "rr", "kkcev", "djjjjsulmrs", "llqbfv", "offqlgj", "paswu", "tlrlcnnrsrf", "jcckqurkg", "jjourlpgeem", "nvl", "shquwmrboghccy", "vncfgelm", "dgcdgjcksk", "vvhvmibflb", "juifgeqkaectlcj", "scvdl", "whcy", "yipswmbp", "wcy", "hbqq", "bsth", "etjurltvpsuy", "dzvytcccevnceyq", "apqrbaoquxqemv", "kvohuqyediyig", "lenybbukzftz", "ffpasoiuztwhtu", "lzlhzqibfv", "wfeemjnjahzdd", "djsulms", "xtudezgzbnzb", "eemjhzdd", "scavdil", "guchrvaqbe", "nvll", "sxzfpzjmxvu", "dytccvny", "grikrnjwezryi", "prng", "ntvmcwwpzo", "laqgcacyxmym", "mglosifyg", "nynvlqll", "vwtn", "lh", "zhhxducgelhy", "prg", "kghierph", "zsucxrqkhahuom", "kvohqydig", "eemjhzd", "offiqcdllgji", "dyc", "toflx", "dzvytccvney", "ghvb", "to", "guchrvab", "wyimthhfzndppwt", "elbqhvwewzzmqif", "hkghiyerph", "hkghiyejrph", "hlsioorugbsuu", "c", "kgierph", "bstbghj", "prbquqev", "mpfsdgimurs", "zfpjvu", "zfpvu", "yxuhwkzvoczfgoz", "gel", "ntvmcpzo", "ekofqduddkhykr", "ekofqdddhykr", "rqeql", "nhnympsum", "xhoqlfolk", "ipmtvxcwuje", "wgmhjhdmnqot", "bsh", "rvncfgelm", "hkahpbb", "lzlzqibfv", "xoqlfok", "tnoftfkogwgplx", "ekofqdddkhykr", "zwiieii", "ujfzzaw", "jfzw", "djsms", "scavdpilj", "tnoftfkoglx", "ps", "vwtnw", "scavhdpilj", "scayvhdpuilji", "pdrshqngzx", "crnrud", "wmhjhdmnqot", "wghmhjhdmnqot", "vbyipsawmxbp", "qknsapkhngorof", "wymthhfzndppwt", "wxcs", "dzvytccevney", "acacmy", "dycy", "teqyrllhbvcv", "uzwyxcs", "wmhjhdmnqt", "qvzkmgfulbyz", "qngglum", "zhhxgdyukcgelhy", "oj", "iljes", "bstbh", "laqcacxmy", "tofx", "ke", "yivkqoek", "djjjsulmrs", "lbirdzvttzze", "l", "zhhxgdukcgelhy", "grikvrnjwezryi", "bltz", "npynvlqll", "gvb", "okzrs", "urbarfkmnlxxn", "qsyzaixtzfazv", "dytcy", "h", "kohqyig", "hgri", "ojdxm", "ujfdfzzaw", "qyrhbvcv", "ebqhvwewzmqif", "uzwxcs", "lebzf", "ysijvkwqmoekromh", "wffaeemjnjahzdd", "crnrduyndj", "ujfdmfzzaw", "laqgcacyxmzgym", "jjourlrpgeem", "kvohqyediyig", "lebukzf", "zwiijeii", "guchrvb", "omoktkyf", "hpgt", "yikoek", "ysijvkwqoekromh", "tvpo", "ysijvkqoekromh", "xbgq", "d", "abmtk", "ors", "rnrd", "xzrugvlzduaxhzc", "njoxacjsnddwrg", "yipswmxbp", "xqsyzaixtzfazv", "urbrfknlxxn", "sxzfpjxvu", "prbaquxqemv", "dvjjjjsulmmrs", "kviahvqu", "urbfknx", "qvmgfulby", "yikqoek", "zsucxrqkhfahuomm", "koqyg", "djss", "moxpfsdgimlurs", "qeql", "urbrfknlxn", "kgieh", "qnspkhngorof", "plxyhib", "scyayvhdpuiljki", "vvhvmbflb", "lpzluhzqxibfv", "kkcbev", "hpzgty", "nyvlqll", "kvahvu", "rklgja", "ipmtavxcwuje", "lbirdzvvttzze", "psw", "fpasoiwhtu", "dgcdgjckk", "qknhsapkhngorof", "qszaixtzfazv", "tvp", "abmtvk", "uwrboghcy", "hbq", "crnruyd", "etjurltvsuy", "etjurltyvpsuy", "lenbukzf", "teqyrllhbvcvg", "ipmvwe", "o", "crnryduyndj", "lbirdzvvqfttzze", "tnoftfkowglx", "ipmtavxcwujre", "omlcoktkyf", "rnperyemtmqh", "bltnzk", "sxzfpzjxvu", "uzdwyaxvcstr", "bq", "rvvugpixa", "laqcacxmym", "wffeemjnjahzdd", "fpvu", "xjbpvekngeyaxbn", "dzvytccevncey", "qgly", "scavdl", "fw", "tox", "toftklx", "prbaoquxqemv", "ztrobzqiukdkcbv", "yivkqoekr", "feemjnjhzdd", "plxhi", "cp", "fyvpzgauttei", "prshqngzx", "kplxyrhib", "suwrboghcy", "kviahvu", "mvwe", "dzvytccvny", "hbqwq", "prbquqemv", "lzlhzqxibfv", "ll", "omblcoktkyf", "toftlx", "lpzlhzqxibfv", "tudegzbnz", "ddgcdgjcgkspk", "kgih", "xjbpvekneaxbn", "suwrboghccy", "zwiiijeii", "dytccy", "ympsum", "jxwakfxsifoj", "uwhcy", "yxuhwkzvoczfoz", "xzfpjvu", "lenybbukzft", "b", "llqfv", "laqgcacyxmgym", "xq", "scavdilj", "zwziivbijaeiig", "scyayvhdpuilji", "amvevfulhsd", "dss", "tlrlcnnrs", "uzwyaxcsr", "qspkhngorof", "etjurtvsuy", "wgqhmhjhhdmnqot", "tvmpo", "tnoftklx", "qgflby", "mlosifyg", "oqyg", "gchvb", "t", "offqcdllgj", "ziieii", "zwziivbijeii", "vp", "lpb", "fyvprzegauttejiv", "vtn", "amefulhsd", "llf", "muzdwyaxvcstr", "zucxqkahuo", "pfsgiurs", "obstbghj", "ipmqtavxcwuzjrbe", "djjsulms", "qvmgflby", "ljpzluhzqxibfv", "jjourlrzpgeem", "zrugvlduaxhzc", "xbpvkneaxn", "ljpzluhzgqyxibfv", "yivkqoekrh", "laqcacyxmym", "nyvll", "muzdwcyaxvcstr", "fyvpzegauttejiv", "offlgj", "vnfgelm", "eteiqyrllhbvcvg", "zsucxrqkhahuomm", "ibiltnzk", "rklgjae", "fpasoizwhtu", "t", "zhhxdukcgelhy", "fpasoiwu", "xzfpjxvu", "tlrlcnnrysrf", "ojx", "mpum", "lxh", "eturtvsuy", "rklgbjaae", "kahpbb", "qngglugmfvmp", "fielbqtcri", "xzruogvlzduaxhzc", "rshquwmrbtoghccy", "nyvlll", "lbirdzvvqttzze", "dgcdgjckspk", "vvhvmibfilb", "dzvytcccevncey", "g", "vwe", "zwxcs", "k", "jourlpgeem", "cpk", "cds", "tlrlcnnrsr", "ivemm", "fgel", "grktse", "urbfknlxn", "qwvzkmgfulbyz", "xjbpvekngeaxbn", "wphuutlgczfspyga", "xbq", "offqcdllgji", "vbyipsakwmxbp", "qyrhbvc", "ygzpztbno", "xhogqlfolk", "ujffzzaw", "xbnmgq", "uwohcy", "rnperyemqh", "prbqqev", "lenybukzf", "mxpfsdgimurs", "ga", "hpt", "moxpfsdgimurs", "vb", "offqcllgj", "rklgbjae", "lifg", "ztrobzzqiukdkcbv", "xoqok", "cs", "snaxxd", "cdds", "qknhsapkhngorohf", "rvqgeql", "rnperyemmqh", "scavhdpuilji", "urbfknlx", "rvvugixa", "ygzpztbndon", "zrugvlzduaxhzc", "shuwmrboghccy", "mlsifyg", "xhoqlfok", "wfeemjnjhzdd", "lbzf", "wythhfzndppwt", "mglqosifyg", "ojxm", "kvohuqyevdiyig", "grte", "prsngz", "eteeiqyrllhbvcvg", "dytccny", "qngglugfvmp", "kohqydig", "fu", "qgfly", "tvmcpzo", "tnoftfkowgplx", "zruglduaxzc", "yijvkqoekrh", "xqsyzaixtzfdazv", "ipmqtavxcwuzjre", "omloktkyf", "ympum", "lzlzqbfv", "pasowu", "rvqeql", "qngglugvmp", "hkghierph", "eemjhz", "feemnjhzdd", "c", "yxpuhwkzvoczfgoz", "dgcgjckk", "lbz", "yxuwkzvoczoz", "zrugvlduaxzc", "ntvmcwpzo", "fzw", "ygzpmztbndon", "rvncfgxelm", "mpm", "tudezgzbnz", "bltzk", "ffpasoiuzwhtu", "cd", "r", "okrs", "byipsawmxbp", "prsqngzx", "wnhnyqmpsum", "ipmqtavxcwujre", "w", "fpasoiwtu", "plxyrhib", "bstbhj", "xbnmrgq", "ipmtvcwe", "urbfkn", "nympsum", "qtngglugmfvmpt", "jckqurg", "hgr", "hpzgt", "rvvxudgpixa", "ysijvkqoekrh", "lebkzf", "guchvb", "kvohqyediyg", "amvefulhsd", "suwmrboghccy", "fvu", "ibdiltnzk", "rnrud", "iem", "urbarfknlxxn", "ygzpztbnon", "prsng", "zcxqkahuo", "ffpeasoiuztwhtu", "laqcacmy", "qszaitzfazv", "xbngq", "qvkmgfulby", "scavhdpuilj", "zsucxrqkahuo", "v", "qtngglugmfvmp", "ysijvkqoekrmh", "lfg", "prqqev", "pasoiwu", "p", "tvmcpo", "kcev", "im", "crnrduydj", "vfgelm", "ddgcdgjckspk", "ivqemm", "ljpzluhzgqxibfv", "lenybukzft", "nhnyqmpsum", "iljesr", "hp", "tqyrlhbvcv", "eemnjhzdd", "xbpvekneaxn", "wghmhjhhdmnqot", "uwboghcy", "guchrvabe", "xoqfok", "fyvpzgautteiv", "pg", "zwiivijeii", "qvgflby", "lsifg"]
puts longest_str_chain(words)
