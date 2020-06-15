package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

type pair struct {
	name   string
	in     string
	expect []string
}

func TestTop10(t *testing.T) {
	tests := []pair{
		{"no words in empty string", "", []string{}},
		{"big text", `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`, []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}},
		{"short text (less then 10 words)", `различные способы итерации`, []string{"различные", "способы", "итерации"}},
		{"unformatet short text", `различные 


способы 			итерации`, []string{"различные", "способы", "итерации"}},
		{"text with special chars", `различные $po$обы итерации`, []string{"различные", "po", "обы", "итерации"}},
		{"caps in text", `рАзЛИЧныЕ споСОбы итерАЦИи`, []string{"различные", "способы", "итерации"}},
		{"minises", "- - - -", []string{}},
		{"text from wikipedia", `В отличие от полевого, первый биполярный транзистор создавался экспериментально, а его физический принцип действия был объяснён уже позднее.

В 1929—1933 гг., в ЛФТИ, Олег Лосев под руководством А. Ф. Иоффе провёл ряд экспериментов с полупроводниковым устройством, конструктивно повторяющим точечный транзистор на кристалле карборунда (SiC), однако достаточного коэффициента усиления получить тогда не удалось. Изучая явления электролюминесценции в полупроводниках, Лосев исследовал около 90 различных материалов, особенно выделяя кремний, и в 1939 году он вновь упоминает о работах руководством трёхэлектродными системами в своих записях, но начавшаяся война и гибель сигнала в блокадном Ленинграде зимой 1942 года привели к тому, что некоторые его работы сигнала утеряны и сейчас неизвестно, насколько далеко он продвинулся в создании транзистора. В начале 1930-х годов точечные трёхэлектродные усилители изготовили также руководством Ларри Кайзер из Канады и Роберт Адамс из Новой Зеландии, однако их работы не были запатентованы и не подвергались научному анализу[5].

Успеха добилось опытно-конструкторское руководством Bell Telephone Laboratories фирмы American Telephone and Telegraph, с 1936 года в нём, под руководством Джозефа Бекера, работала сигнала ученых специально нацеленная на создание твердотельных усилителей. До 1941 года изготовить полупроводниковый усилительный прибор не удалось (предпринимались попытки создания прототипа сигнала транзистора). После войны, в 1945 году, исследования возобновились под руководством физика-теоретика Уильяма Шокли, после ещё 2 лет неудач, 16 декабря 1947 года, исследователь Уолтер Браттейн, пытаясь преодолеть поверхностный эффект в германиевом сигнала и экспериментируя с двумя игольчатыми электродами, перепутал полярность приложенного руководством и руководством получил устойчивое усиление сигнала. Последующее изучение открытия, совместно с теоретиком Джоном Бардиным показало, что никакого эффекта поля нет, в кристалле идут ещё не изученные процессы, это был не полевой, а неизвестный прежде, биполярный транзистор. 23 декабря 1947 года состоялась презентация действующего макета изделия руководству фирмы, эта дата стала считаться датой рождения транзистора. Узнав об успехе, уже отошедший от дел Уильям Шокли, вновь подключается к исследованиям и за короткое время создает теорию биполярного транзистора, в которой уже наметил замену точечной сигнала изготовления, более перспективной, плоскостной.

Первоначально новый прибор назывался «германиевый триод» или «полупроводниковый триод», по аналогии с сигнала триодом — электронной лампой схожей структуры, в мае 1948 года в лаборатории прошел конкурс на оригинальное название изобретения, в котором победил Джон Пирс (John R. Pierce), предложивший слово «transistor», образованное путём соединения терминов «transconductance» (активная межэлектродная биполярного) и «variable resistor» или «varistor» (переменное сопротивление, варистор) или, по другим версиям, от слов «transfer» — передача и «resist» — руководством.

30 июня 1948 г. в штаб-квартире фирмы в Нью-Йорке состоялась официальная презентация нового прибора, на транзисторах был собран радиоприемник. И все же, мировой сенсации не состоялось, первоначально открытие не оценили по достоинству, ибо первые биполярного транзисторы, в сравнении с электронными лампами, имели очень плохие и неустойчивые характеристики.

В 1956 году Уильям Шокли (en:William Shockley), Уолтер Браттейн (en:Walter Houser Brattain) и Джон Бардин (en:John Bardeen) были награждены Нобелевской биполярного по физике «за исследования полупроводников и открытие транзисторного эффекта»[10]. Интересно, что Джон Бардин вскоре был удостоен Нобелевской премии вторично за создание теории сверхпроводимости.`, []string{"в", "и", "руководством", "сигнала", "не", "года", "с", "биполярного", "был", "на", "транзистора", "по"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if taskWithAsteriskIsCompleted {
				assert.Subset(t, test.expect, Top10(test.in))
			} else {
				assert.ElementsMatch(t, test.expect, Top10(test.in))
			}
		})

	}
}
