/*
 * Bamboo - A Vietnamese Input method editor
 * Copyright (C) Luong Thanh Lam <ltlam93@gmail.com>
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * This software is licensed under the MIT license. For more information,
 * see <https://github.com/BambooEngine/bamboo-core/blob/master/LISENCE>.
 */
package bamboo

type inputMethodDefinition map[rune]string

// todo: move to input_method.json
var inputMethodDefinitions = map[string]inputMethodDefinition{
	"Telex": {
		'z': "XoaDauThanh",
		's': "DauSac",
		'f': "DauHuyen",
		'r': "DauHoi",
		'x': "DauNga",
		'j': "DauNang",
		'a': "A_Â",
		'e': "E_Ê",
		'o': "O_Ô",
		'w': "UOA_ƯƠĂ",
		'd': "D_Đ",
	},
	"VNI": {
		'0': "XoaDauThanh",
		'1': "DauSac",
		'2': "DauHuyen",
		'3': "DauHoi",
		'4': "DauNga",
		'5': "DauNang",
		'6': "AEO_ÂÊÔ",
		'7': "UO_ƯƠ",
		'8': "A_Ă",
		'9': "D_Đ",
	},

	"VIQR": {
		'0':  "XoaDauThanh",
		'\'': "DauSac",
		'`':  "DauHuyen",
		'?':  "DauHoi",
		'~':  "DauNga",
		'.':  "DauNang",
		'^':  "AEO_ÂÊÔ",
		'+':  "UO_ƯƠ",
		'*':  "UO_ƯƠ",
		'(':  "A_Ă",
		'\\': "D_Đ",
	},
	"Microsoft layout": {
		'8': "DauSac",
		'5': "DauHuyen",
		'6': "DauHoi",
		'7': "DauNga",
		'9': "DauNang",
		'1': "__ă",
		'!': "_Ă",
		'2': "__â",
		'@': "_Â",
		'3': "__ê",
		'#': "_Ê",
		'4': "__ô",
		'$': "_Ô",
		'0': "__đ",
		')': "_Đ",
		'[': "__ư",
		'{': "_Ư",
		']': "__ơ",
		'}': "_Ơ",
	},
	"Telex 2": {
		'z': "XoaDauThanh",
		's': "DauSac",
		'f': "DauHuyen",
		'r': "DauHoi",
		'x': "DauNga",
		'j': "DauNang",
		'a': "A_Â",
		'e': "E_Ê",
		'o': "O_Ô",
		'w': "UOA_ƯƠĂ__Ư",
		'd': "D_Đ",
		']': "__ư",
		'[': "__ơ",
		'}': "_Ư",
		'{': "_Ơ",
	},
	"Telex + VNI": {
		'z': "XoaDauThanh",
		's': "DauSac",
		'f': "DauHuyen",
		'r': "DauHoi",
		'x': "DauNga",
		'j': "DauNang",
		'a': "A_Â",
		'e': "E_Ê",
		'o': "O_Ô",
		'w': "UOA_ƯƠĂ",
		'd': "D_Đ",
		'0': "XoaDauThanh",
		'1': "DauSac",
		'2': "DauHuyen",
		'3': "DauHoi",
		'4': "DauNga",
		'5': "DauNang",
		'6': "AEO_ÂÊÔ",
		'7': "UO_ƯƠ",
		'8': "A_Ă",
		'9': "D_Đ",
	},
	"Telex + VNI + VIQR": {
		'z':  "XoaDauThanh",
		's':  "DauSac",
		'f':  "DauHuyen",
		'r':  "DauHoi",
		'x':  "DauNga",
		'j':  "DauNang",
		'a':  "A_Â",
		'e':  "E_Ê",
		'o':  "O_Ô",
		'w':  "UOA_ƯƠĂ",
		'd':  "D_Đ",
		'0':  "XoaDauThanh",
		'1':  "DauSac",
		'2':  "DauHuyen",
		'3':  "DauHoi",
		'4':  "DauNga",
		'5':  "DauNang",
		'6':  "AEO_ÂÊÔ",
		'7':  "UO_ƯƠ",
		'8':  "A_Ă",
		'9':  "D_Đ",
		'\'': "DauSac",
		'`':  "DauHuyen",
		'?':  "DauHoi",
		'~':  "DauNga",
		'.':  "DauNang",
		'^':  "AEO_ÂÊÔ",
		'+':  "UO_ƯƠ",
		'*':  "UO_ƯƠ",
		'(':  "A_Ă",
		'\\': "D_Đ",
	},
	"VNI Bàn phím tiếng Pháp": {
		'&':  "XoaDauThanh",
		'é':  "DauSac",
		'"':  "DauHuyen",
		'\'': "DauHoi",
		'(':  "DauNga",
		'-':  "DauNang",
		'è':  "AEO_ÂÊÔ",
		'_':  "UO_ƯƠ",
		'ç':  "A_Ă",
		'à':  "D_Đ",
	},
	"Telex 3": {
		'z': "XoaDauThanh",
		's': "DauSac",
		'f': "DauHuyen",
		'r': "DauHoi",
		'x': "DauNga",
		'j': "DauNang",
		'a': "A_Â",
		'e': "E_Ê",
		'o': "O_Ô",
		'w': "UOA_ƯƠĂ",
		'd': "D_Đ",
		'[': "__ươ",
		'{': "_ƯƠ",
	},
}
