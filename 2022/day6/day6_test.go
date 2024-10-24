package main

import (
	"testing"
)

func BenchmarkIndexUniqCharRun(b *testing.B) {
	type test struct {
		in string
		l int
		want int
	}
	
	tests := []test {
		{
			in: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			l: 14,
			want: 5,
		},
		{
			in: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			l: 14,
			want: 9,
		},
		{
			in: "nppdvjthqldpwncqszvftbrmjlhg",
			l: 14,
			want: 9,
		},
		{
			in: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			l: 14,
			want: 15,
		},
		{
			in: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			l: 14,
			want: 12,
		},
		{
			in: "mzrzqrzqrrmcmgmrggvjvcczbczccdwdcwcpwwclwwbttdntnsnllpggqhgqgzqggvjjjfqfhhbbvzvmzvvfccnznwntnptpgpcplpwwvgvzvvcgvvbcvcqqhnqqbsbdsbbbgtgvtgvtgvvrfvvjbjcbcchnhrhvvrmmwzwzmwmggjwwwwzrrmbmbmbrrdjrjfrrrmrffzgghtggwddtllhchdhsshjjwfwsfssnbbnjncczsczzvpvssljsjgsjgjcgjggspsfsfhfppsqqzlzjllblsldsldlndnhnffflwfwmwssvzzvqzzblzbbmrrfnrrqnnhfnnhphbhrrlvrlvrvcrrltrtmrrcscchdhqhfffrddzsszwwrbrfftddbwwtdtnndrrghgcgfgrglllcdcsswvwsvszsddjwdwwcvwvmvhhlqlddbhbrrnddpqpttrtctmcmjjtvjjfvvlcldltdtdqdllnrncrnrznzfnfsfnndqqbhhlnhhhsmhhfghgppflllrhhvdddvrvnrnpplwplldttswsbwwtbbvwbwnnhcczmczmczclcttwcchllmsmbbrvbbsjsswbwwbtwtgtrgtrgrnggqwgwmwrmrqmmghhtbbrppjmpppjtjtffvqvbbwsspzsstlstltrrtmtrmrzmmgqqgmmpzmpzpmmgnnrjjgrjjhgjjsccjllwzzsddznnrsrttdllcnntvnttqjjqljqqwlqlmmfdfttzzvgvtvjjwvwsvvhbhrbhbwbzbppbplphpjpzppljjnttnvnggqbqffqgqrrjwjcjllcblbqqglqgqjqrrffjqqqcvqccgrgqgrqrnqqhphccvwwsgspprmprrwdrwrdwdrwrwswbsbcctscttwvvvstvthtrhhddtppgwwqpwwzpzgzszgzvggzqqmbqqrqsrqrmmjjfggztzdzvvqdqzddrrnmmtccsstdsdrrmrnnzzpggbmbqbcbzczppbjbssfzfcfvcvlclgcglcgllvclcwwwpmwwjrrgtrtfrrnznssvvrjvrjrgjrrcsrrstrsrhrpplfljjwrwbbffdpfpzffrhhjrjdrrgpggbhhschcdcdcqcnncjcvjcvjvbvcbvbssnhhsddclcrcrjrdrtdtqdqrddddlcdcttzzlslblbqqcmmdgmmsnmnmccqmmljjncnhnchcdchddwsdssdrsscfcnffcqfqtfqqjcjrrvbrvvhsvvlqlwwsccjbbgpgzgcgbgzgccpvcpvpddqrrqfqtqrqqcjjwvwttrwwblwlrwllnsshdhqdhqqjllzlflvlgllhnllstsbbdqqpnpwpmwwzhzvzcclcvcfvcvssdcdvdpphcphcppffndfnfmfbfwwgzwwwvhwwmcwwwtmtgtvggvbggzzthzzlbzbzmmhlllvgllmfmpmbbldbdnnzzjcctvvsdvssrpsslglqggcchvcvllzhhbcczjjgllvqlvlflglwwscctfcbmdjhgwtfhjrgvvvrmdcpsrtsnvhwnnznnnmrhcnlnmjvdbqztspwbdwlttdtwlwdvqjpgqdzbnssglqczqwvfgdlbdmsbbmggjdhddzvzlqhrggvvccmcmtqdmmpqvvstmqgvntdsmjbzbwstdrmjjmzfgmczjrftwrwnfbdmlddvzdfwwztldqfvdswfdwrfcgptmmjnzngwnflzlvtwpdsvllfwqjnjjjbfcwmgrvlhdvpgprnjthlqlbtlhflzwqjwtqtdzcvqbgtjlnmwchpjrgfndrnzwctjwvwsvflmgnqmplhhwljqcshqqldcnfrdbdcbslrtbnqrvhhjrddmjbtmcjzlqcdcqbfftsmpdpnfrjjwmdpmrvjjmppbnmcbvnzbphjcjlldnpdcwsqfgpmbszvjbpzngvncbfncpmghfrsssrnbbmhnjjvgmjzwtwlpszphgwvtjzdsmvvhnrplmnrllvqvphtbbrnlwmffdmmbhvzfcnbnlbspfwbcbjwgpnbsfpfbdqbnpfcqpngfqwlwcgrzlvfqgfsbppgmfzgjdntdvzvzclbrfwpqmvgmtmppsjzjhqbgjnntdnwvwljndqbtfgrjcfnlpsgsmbdljdzqfclgwhdgndzwbcthmrgzcjbhrrltlsznbsvswtsjbhwqzwchwzbtqmnrqfrzspwrqwwqwvczmntwggltjmvwdftplrgtrltmpvvgslvhctnwwcgtblmqtnsmqdgcbvzlcbrrsstzgzrgzgtmqtvsnzqqwlfgrjbdsvpdjpgmmczddbptwvpthvdjvrpqsbbgctrpqsdnczzbttdrbflrllrnvjswslnghdqfqlnbcctwbnpvspfmrcnprpjzgwsdwlszzpdcgvzjsjtdgjnpwgtqmvsmsvzddcmfwtqjghrrrmvcjbccgrngcvvfvrrmmqmzsvhdqsfqpqnnhhpffrtblnzhwqmsgvvzvgpbczsznzlnvsgjjzqbpjmvpjzqqpzpvplwsvgtwvrhlvggrpvztvchbcwflwhwmvdslvhsgpcqhdhcpvclgmzdngsmbplzrqgwflltzgsglflhqhpnnfszvtgprqhnjlrvwfsjfpcjzbznsprtldvwjlnhqsdlfnwtlpzldfpnjpnnqfgqqrjrnszznlrjgwwlnbbznzcshtfqcprphgsldtgzzgqjvmjrtfhmgptrnqzmtfclqwgprvjptntgwhbqwhdjcfjtrmdwbslhqwgtcqwhjhhmltrlsqbvzqsjqsczvsqlsjdgmnjnpvmmwbsfnfnshmlrdgslbfrnhmdgcjgsqgwlbmlhhnfmcfnswmwgrzmpdmtrcmwgjlgrhzchmhmbjgjrfgdvplvngjnfdrscwqtnsspdszznnnjrdmvwbvnzmmddbcvsphfqsbcmfmfgrcdfjptzqfsjsmjtbswtqzvtfpwbvqsgjzftfssdmcngtgbfmrrsjtllnpsztvgbpgpcrmqsgvcjrpprhqzjwlcdjpvgrqtjcglwnvzvqtncmndzbgpgfbprpqqjfbnzhfzlfrgnfnbpjfrbtjgrgvmczthqsbfplnzvhqjmrrssnzjdhfvlngwptzzdshzhzfpzhmrcjmdvlnlqtsnlbhqgplsltgznshthcwnpzpcgjjwhjstfcdtmmnvbgqzrjpsqwlmlbjnqtrfmbrdsbdlrrwqqvzdnmzhzvrhdftzldzjrtfgzlrwczbzzzqtthdcjtqpcfdwpqlzqmrnsnhwbzhjnqjzgfdcdcrqsjfhpjcdtwnvwzzbwfgdzcmmfvdvdpjsltcrbfqczvvbjscptjsvzndgwhzfjcqljndrcqzhssnqtmjmwjrsqdpqnhwsntqnmrwhvsnvrvpvwbndpmgsnchtnzpjgshcvjgwlwdhqbqhvcwqmwllhcgvbslqvwgswvnsqvvcphbglrvlsczcscznlwzrvjptmbmrgjlhvbrlghcwjdqjlgrcbrhfmbcwgclqzllqtmshbjnhwpgsgtljbwhgcsznvslmglcnmmgjlsptgdqlbtclfbfvjpmqgwwbcvwtdhgjlpvgggjjvmltdgzjdsvtswgblgfcvdvtwrrfljtqjhdflhbtwmcdqpmdqrtjsvdhzstfnqjrzztnwslvwtvgvzfqlzfrhnthjfpvmwmdgrtzhbtwdfscdpwmwwrhgbmqrdftvjvrgzhbpqtqvvlwbmbvlszzqjwhtvwnsjcdcgdlwlrtvjsdqzngcdtpvsddsbqbhtrrpwqdhvdmnnncgccszqtbcgwbdbwrnwrhpwprslbmhrmwpmqzvssfzvrmwrmrzmhrcwvbdtvdflgmrghqngwsgrnctsnhpnmcmfmrtszttqtvvlhdjgplvgnjrtgnfgmdtvwzmzbtzvmhdvpcjqvgpsmdcfrbmqbrlsnccldrfdldqfnsfzznqtvsgwbljgrvbdmggdmhvvdzjfllzwzpddcnvrfggsddqmczfnnfvrwsmvfctctjqdrhvlntflccqgzg",
			l: 4,
			want: 1527,
		},
		{
			in: "mzrzqrzqrrmcmgmrggvjvcczbczccdwdcwcpwwclwwbttdntnsnllpggqhgqgzqggvjjjfqfhhbbvzvmzvvfccnznwntnptpgpcplpwwvgvzvvcgvvbcvcqqhnqqbsbdsbbbgtgvtgvtgvvrfvvjbjcbcchnhrhvvrmmwzwzmwmggjwwwwzrrmbmbmbrrdjrjfrrrmrffzgghtggwddtllhchdhsshjjwfwsfssnbbnjncczsczzvpvssljsjgsjgjcgjggspsfsfhfppsqqzlzjllblsldsldlndnhnffflwfwmwssvzzvqzzblzbbmrrfnrrqnnhfnnhphbhrrlvrlvrvcrrltrtmrrcscchdhqhfffrddzsszwwrbrfftddbwwtdtnndrrghgcgfgrglllcdcsswvwsvszsddjwdwwcvwvmvhhlqlddbhbrrnddpqpttrtctmcmjjtvjjfvvlcldltdtdqdllnrncrnrznzfnfsfnndqqbhhlnhhhsmhhfghgppflllrhhvdddvrvnrnpplwplldttswsbwwtbbvwbwnnhcczmczmczclcttwcchllmsmbbrvbbsjsswbwwbtwtgtrgtrgrnggqwgwmwrmrqmmghhtbbrppjmpppjtjtffvqvbbwsspzsstlstltrrtmtrmrzmmgqqgmmpzmpzpmmgnnrjjgrjjhgjjsccjllwzzsddznnrsrttdllcnntvnttqjjqljqqwlqlmmfdfttzzvgvtvjjwvwsvvhbhrbhbwbzbppbplphpjpzppljjnttnvnggqbqffqgqrrjwjcjllcblbqqglqgqjqrrffjqqqcvqccgrgqgrqrnqqhphccvwwsgspprmprrwdrwrdwdrwrwswbsbcctscttwvvvstvthtrhhddtppgwwqpwwzpzgzszgzvggzqqmbqqrqsrqrmmjjfggztzdzvvqdqzddrrnmmtccsstdsdrrmrnnzzpggbmbqbcbzczppbjbssfzfcfvcvlclgcglcgllvclcwwwpmwwjrrgtrtfrrnznssvvrjvrjrgjrrcsrrstrsrhrpplfljjwrwbbffdpfpzffrhhjrjdrrgpggbhhschcdcdcqcnncjcvjcvjvbvcbvbssnhhsddclcrcrjrdrtdtqdqrddddlcdcttzzlslblbqqcmmdgmmsnmnmccqmmljjncnhnchcdchddwsdssdrsscfcnffcqfqtfqqjcjrrvbrvvhsvvlqlwwsccjbbgpgzgcgbgzgccpvcpvpddqrrqfqtqrqqcjjwvwttrwwblwlrwllnsshdhqdhqqjllzlflvlgllhnllstsbbdqqpnpwpmwwzhzvzcclcvcfvcvssdcdvdpphcphcppffndfnfmfbfwwgzwwwvhwwmcwwwtmtgtvggvbggzzthzzlbzbzmmhlllvgllmfmpmbbldbdnnzzjcctvvsdvssrpsslglqggcchvcvllzhhbcczjjgllvqlvlflglwwscctfcbmdjhgwtfhjrgvvvrmdcpsrtsnvhwnnznnnmrhcnlnmjvdbqztspwbdwlttdtwlwdvqjpgqdzbnssglqczqwvfgdlbdmsbbmggjdhddzvzlqhrggvvccmcmtqdmmpqvvstmqgvntdsmjbzbwstdrmjjmzfgmczjrftwrwnfbdmlddvzdfwwztldqfvdswfdwrfcgptmmjnzngwnflzlvtwpdsvllfwqjnjjjbfcwmgrvlhdvpgprnjthlqlbtlhflzwqjwtqtdzcvqbgtjlnmwchpjrgfndrnzwctjwvwsvflmgnqmplhhwljqcshqqldcnfrdbdcbslrtbnqrvhhjrddmjbtmcjzlqcdcqbfftsmpdpnfrjjwmdpmrvjjmppbnmcbvnzbphjcjlldnpdcwsqfgpmbszvjbpzngvncbfncpmghfrsssrnbbmhnjjvgmjzwtwlpszphgwvtjzdsmvvhnrplmnrllvqvphtbbrnlwmffdmmbhvzfcnbnlbspfwbcbjwgpnbsfpfbdqbnpfcqpngfqwlwcgrzlvfqgfsbppgmfzgjdntdvzvzclbrfwpqmvgmtmppsjzjhqbgjnntdnwvwljndqbtfgrjcfnlpsgsmbdljdzqfclgwhdgndzwbcthmrgzcjbhrrltlsznbsvswtsjbhwqzwchwzbtqmnrqfrzspwrqwwqwvczmntwggltjmvwdftplrgtrltmpvvgslvhctnwwcgtblmqtnsmqdgcbvzlcbrrsstzgzrgzgtmqtvsnzqqwlfgrjbdsvpdjpgmmczddbptwvpthvdjvrpqsbbgctrpqsdnczzbttdrbflrllrnvjswslnghdqfqlnbcctwbnpvspfmrcnprpjzgwsdwlszzpdcgvzjsjtdgjnpwgtqmvsmsvzddcmfwtqjghrrrmvcjbccgrngcvvfvrrmmqmzsvhdqsfqpqnnhhpffrtblnzhwqmsgvvzvgpbczsznzlnvsgjjzqbpjmvpjzqqpzpvplwsvgtwvrhlvggrpvztvchbcwflwhwmvdslvhsgpcqhdhcpvclgmzdngsmbplzrqgwflltzgsglflhqhpnnfszvtgprqhnjlrvwfsjfpcjzbznsprtldvwjlnhqsdlfnwtlpzldfpnjpnnqfgqqrjrnszznlrjgwwlnbbznzcshtfqcprphgsldtgzzgqjvmjrtfhmgptrnqzmtfclqwgprvjptntgwhbqwhdjcfjtrmdwbslhqwgtcqwhjhhmltrlsqbvzqsjqsczvsqlsjdgmnjnpvmmwbsfnfnshmlrdgslbfrnhmdgcjgsqgwlbmlhhnfmcfnswmwgrzmpdmtrcmwgjlgrhzchmhmbjgjrfgdvplvngjnfdrscwqtnsspdszznnnjrdmvwbvnzmmddbcvsphfqsbcmfmfgrcdfjptzqfsjsmjtbswtqzvtfpwbvqsgjzftfssdmcngtgbfmrrsjtllnpsztvgbpgpcrmqsgvcjrpprhqzjwlcdjpvgrqtjcglwnvzvqtncmndzbgpgfbprpqqjfbnzhfzlfrgnfnbpjfrbtjgrgvmczthqsbfplnzvhqjmrrssnzjdhfvlngwptzzdshzhzfpzhmrcjmdvlnlqtsnlbhqgplsltgznshthcwnpzpcgjjwhjstfcdtmmnvbgqzrjpsqwlmlbjnqtrfmbrdsbdlrrwqqvzdnmzhzvrhdftzldzjrtfgzlrwczbzzzqtthdcjtqpcfdwpqlzqmrnsnhwbzhjnqjzgfdcdcrqsjfhpjcdtwnvwzzbwfgdzcmmfvdvdpjsltcrbfqczvvbjscptjsvzndgwhzfjcqljndrcqzhssnqtmjmwjrsqdpqnhwsntqnmrwhvsnvrvpvwbndpmgsnchtnzpjgshcvjgwlwdhqbqhvcwqmwllhcgvbslqvwgswvnsqvvcphbglrvlsczcscznlwzrvjptmbmrgjlhvbrlghcwjdqjlgrcbrhfmbcwgclqzllqtmshbjnhwpgsgtljbwhgcsznvslmglcnmmgjlsptgdqlbtclfbfvjpmqgwwbcvwtdhgjlpvgggjjvmltdgzjdsvtswgblgfcvdvtwrrfljtqjhdflhbtwmcdqpmdqrtjsvdhzstfnqjrzztnwslvwtvgvzfqlzfrhnthjfpvmwmdgrtzhbtwdfscdpwmwwrhgbmqrdftvjvrgzhbpqtqvvlwbmbvlszzqjwhtvwnsjcdcgdlwlrtvjsdqzngcdtpvsddsbqbhtrrpwqdhvdmnnncgccszqtbcgwbdbwrnwrhpwprslbmhrmwpmqzvssfzvrmwrmrzmhrcwvbdtvdflgmrghqngwsgrnctsnhpnmcmfmrtszttqtvvlhdjgplvgnjrtgnfgmdtvwzmzbtzvmhdvpcjqvgpsmdcfrbmqbrlsnccldrfdldqfnsfzznqtvsgwbljgrvbdmggdmhvvdzjfllzwzpddcnvrfggsddqmczfnnfvrwsmvfctctjqdrhvlntflccqgzg",
			l: 14,
			want: 2504,
		},
	}
	
	
	doTest := func(name string, indexUniqRunFunc func(in string, l int) int) {
		var testIdx int
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				t := tests[testIdx]
				testIdx = (testIdx + 1) % len(tests)
				
				got := indexUniqRunFunc(t.in, t.l)
				
				if got != t.want {
					b.Fatalf("input %s len %d got %d want %d", t.in, t.l, got, t.want)
				}
			}
		})
	}
	
	doTest("indexUniqRunMap", indexUniqRunMap)

	doTest("indexUniqRunSlice", indexUniqRunSlice)
}