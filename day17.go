package main

import (
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"strings"
)

func D17() []int {
	inputs := []string{
		"2413432311323\n3215453535623\n3255245654254\n3446585845452\n4546657867536\n1438598798454\n4457876987766\n3637877979653\n4654967986887\n4564679986453\n1224686865563\n2546548887735\n4322674655533",
		"111111111111\n999999999991\n999999999991\n999999999991\n999999999991",
		"112211323123313231221322424444241114332223124451424553135231525524415314331423354424151142512532322133331244112344412431332132122133112212211\n122223123113132112231214112341424142422134142231351412332345532123322123241131151432434144322125551312414133341314131233422322322132111232121\n221122112112311332233321231113341442341431433433123255542341432521344514231533554131332544154431243333341241244434322431231213223221323312122\n122311112332123113214214443411114431243245541442442314425124453113434525224551243115454313325131424435543442211434142111214432312313213111231\n221131232332113113124214432212132144441514344335332253221455141515253422334534111522444433122415241343353534244432442321124124232132131123331\n123131223112232321314232231314122241534352424515445222251215213441353512511514442325231214135114551233123312331211414122342233321331312323112\n122321323122232121231114233242221411123453351115411115521555542134333144112252155243544542352114554341151155313414313131212242123231133131131\n122323331323211212211232114341143323514135243431225323132111341433662544454663353424334313125555454251424133451423232441231332132331131212231\n222312113122222242431121412432545314512223332313345534235644563566325643465425426634331434523532122242223524444413123334241424332411231113231\n122111111332321341242222431344331155241453421311115451244562222343566245662244532443532552311451453221331553512532222443312121122131221331131\n211112131321313111213243243115514145143125122211515433365535622426654345624525525234333246354433551121522213135512524142431133421223331222332\n332221332241423422413324423251414212133134121523644263544446664256365263643446536566245642362633154125433441135234311423242224312232212322313\n322323122114334111322321112133243524454254222442253662263536542553533455632423423662562465342346232155314152412341311443243332433123212111222\n223123214141212134414231144425414224513354165255424433646234654556464344256336654362234463335364654652233143453553512525243111124344232322323\n213211243341111323431154353411543152544114566356254543323665266522655544252426533335633356252633326332225111121341255122114112331311114431132\n321111133133423124333523224122255151152362664366345662554442255265462634225563346243426326635555345552621535241451231321214222432412334233122\n322131122243433432334245413421243332514636532444226456652223656624236423653254266363654464224665633425665432331215551141241132331333234321313\n322142131414211441452131422143253154442444526635535325223465364425322255626656553255652356556532442226522643425453345144534132232413243413232\n323313332322212432215521523322422333566252326463452625452252265477455474744736554443552364566256366243652242243433452552113321114441412444132\n211241313133333345251225251344243266223462444332326222455444336333345357774335456467633423225336224664332536563552244433523421214314321222312\n331131434221313455323134414424214223363223442425266545535435664754373537566677646563343633245645342326343263243515333332211422522413423213342\n141423344331424434352345424543126636633355225645544377767764545645563546446436677657467757725226534444623542456614522433111114322324412223333\n211143233223222215413143353434536626545433623235365343334536533643743767466434567767457566466463445632565435222345522144324535233222233213434\n134331321433255133242124133522423443624562335224465333374447475665356644645334736456367773363437356365455332264662335132313344324222334442124\n113431331144311133155213431642246556524244635646437466637656636743753537333345464443734355565447766443455623653426452442155134534142322433113\n144321334314413325242124246436432632623464223355644544536344567743756353746447773774344745746545655424465655566542336311535153443112232331112\n442113343112442234143334142456454633225564436443353363374763377664567374467766334757563464567736666766242546524645246622112155431244114311323\n334121314222313521523432434232425625426254744464373777736653635463447756663433364565635635543755463465242452522335542563532412215323133113212\n332243241225152422452326666555664342644734757467474663356536756745456385576374375575673565365765474675546462234555662246554132131252431313424\n241434234133454125134424432562436334325737464475635676335437655754644488676447657743633644654664776446763552252322355365453131115532513323143\n232412224422454121522224455353232363277737555633364564334748574665565857774556865588766774746675573473745522345444343326515425342113535242213\n212212431113325533545245444652334426337553477677437353436874785658686576746456678685858774733363344333766363645462245346265115142215543321242\n342231211352533113535232636332643364434746636346544648784588756665757845678656754477847668655577567734377375735652334322264441331234152322223\n413221534254544111345265534543323644734777344646547874685858478467768454666848475886678848843474373544536547446645244563355452341212121311214\n442324445331341151356453424233425446347566777443738764877674687688477845668857688866556456644753457374537655767435644454424441513455332152331\n213342321411224345463464324633427347473336336767665777486586858885774576444574757446766465465477343535457435465554566635553552513153454453112\n123315124111152146326446662353377776474755676666686876566554875884745788588656678884784486657658733334366577334445336436542232532322552354323\n441113215251343153426426355634374537755334545776467446786774455866885557577884656487867578876758653433747433376574454365556365112245152344444\n211255332155135545632653625633463755635755457578648685888484885588755888477675457788665447658445576756663537345535643662623424655345213323433\n331353445441324556525523564467444337544653767465556784487474846556898587558767778878655456644666488437677667435555444526244334231424451445221\n123333154152451252534543252573566753575635545458687876848857466786858675656567767488888844674688657546536753636757726232236245343151233513334\n142423552143344564523236363644376756674538658774548874866855686597579656756556966678676578867845878565456553335376553233423364463352251415113\n335421224353356425542355567644454434343347888886454557764659788777578578588798675757867585778576587657745656475645636265226556256135234222443\n235215225545235654446222343773445476466874647586557558596989765968898988568669669867867865654576667467476637433566656535452243344623134515423\n414452445555465634242634666746365536577848667457766745575596958689779988579598689878668986466746687557858537446677363535246652563433411122352\n135212532214534364622663665744565654456844586754747598698866959597756766968585679786589577977665755868878664567336466552236556343421141214335\n444325512232245254452324463436535374558485686757754785756767557765779666699675577786699659766667465658485565663476674445422633535423525414331\n424342423225534235262333733646577777478846477467659555887858957756878889585955786785695768976687578786765775346545645435522535654632234232535\n245353153523564566552333477364357554777646776448785965576789868677979559556996898895796865676958787678686747674356434473352554632245434523444\n344345313213624544332334765537566355648844668574958989597776999558797656775875875799668596975574567546884646536436456563644253623633252111222\n121445142165245656355263636557463464646886547689666556677887878587987767989977779595857889896757875456645574847657764756456564635625215123433\n314313354554255363333554565377367877777885554759755859657958868999787997686896868595598785699585984555864648787645536545645625546266541445542\n531232355253524565246657747555433867475457478965959876997886999897788689978678898999785968779798665748455846753565754367333656434253453314411\n511314543156362655424674334736635885475887786787686667957676869689766969999696998898559776756787877468477877846376467663346534642263241214111\n313522211662634443664765565675634566656646756988767768567787898686697687768687667687688759668785865844658844855656676563577536252656525533242\n414225114645642253554576663335367644765867686566666955966798899896999699986878768988869989799776686965675846647647656776767636424636222433535\n214434515256533563364667747733476876475687495668595977766697679678866968886769966877887777878858889586548666656464476545444325453342531314313\n224543454656536553343777767656458566655464859566965669579779868986998796669779868686867856786758786784686556886544573473457556426433626215424\n334411416624456532666353546663685777878766996696969657766769767696868668969688789988886897886889855954647667676453364455446365456524654552354\n132542354442364222464337356374664556466576685659859769788879899979788678878886976667887786667879856896568448876576634333767562242335446522333\n345543315626333265573536777536455856686877896786669858996878897789867788789988776666997866557676668978784854568674637445436722234622665545515\n145132224446323562245565734463864447486547796578577796977786696986878779798988767697677697965678899559555776576853343473474345456534424145554\n441343216322662453437635733647755656847877968558855789786679998787979988778778796768986999879656587657475765445785344543366662452345232222444\n552253142442342424355557676445587448565486656769586677679977699687987979888888966678877698667577759688977455456775657634546665333543426433312\n112153336343643463353474737736567475458895678887967996776866799798799997788778887997686897695978675768558665775867676446564464433426632345155\n353433555265552343563746656548788557764576879775586798786696667998799978777889896686997899775699766556757777857885433367737743245656535653131\n221321126366552432647653655558566564884886557959897699879967687999799779887997897886988966995868867966765846666878573673365544644436334351351\n232532436265523532546744767758477885654797999896666879987666767798777798789778998986869889668959875958866855775774766557746436464365653633333\n414411446354244635556474336377766678746599685896597668678968798778879799979888978797767789988797579597676477878477556544467436635225646452355\n114213465232332525737676445445675747745888567659566989797979778987899997877999888899966676996679795795565555756887643777745774225365464344154\n352154442653636535653537747467465675778867985597897877867688798878899778889789978879969767887589567577858457678845476547656536353534253334111\n352443126524636554745637345445847845645877856665679986967799678879977898898778798987979769868858688968784586754445743436463365256622435451243\n252322546342242363373346764376674475766779755789868788978986697797799779787878897986877987696666766655975888855478576534774476456253325441443\n143143144533345326575754566455886447775757758788676889779688689989877977878978899999698667878859767959567685776486547743737375435236424232522\n331432126625563465644337534674565647546858578695577898878989797878877877889878898768966768996669665795654466884484445777337754436264553413543\n244112533533634355775437376565475754548599666667557686976889867787899998779777899979668978787997969969574686676747457674666545232242233222441\n214541243424626343573546465465756585454797785889989988889687798779778898998989889999879979988595757897946657864875365356564366626556624645413\n522121434542522663645546346367586645865568956766685977999697977989878778879997776788788788968667686678855657854765757644457633223325535341451\n154255452664246252753354676566685657854469789987968776689788698879988979899797777776978979668956689897945586585548774434574644245634665445112\n222512466225335446335467647544876884476685569759978768968696969887999797787999896978898967677767887875547854687775755437676732433354525643143\n524552212452546532667656346446658585674579585857786887868686998896989987778898976686899766765775567898885674564477654645435436245444655234222\n442121455443625662475466664444665654585549686889885867897998668787879788997997766696989778968665768765487576444475467735773354233352524425215\n531525326234354266667366737535576648756876959668756789676868799767986678887696868988796886859965695877785784884555776745346433662365443454113\n352334515336352543363455645347586744887554687798569555866677697869987897977678998899896976779686986558458548868775766665377444345663345224453\n231332234556653334256766635664776647565586878877959595698787976776999878988678866867989798675987677878846645576563366637754722253233646454442\n553552145444452646336457677653554755664555579569865699976889996986698889969678777779878986985698778654568584866673333376665232233452244421134\n213522351235464554255367674744445657575585779697776789977977976766789769689897997786987698696575565944587644558875765474375446333632232541321\n352241334544262333325476574573337754585765499657858895785878996779668686996689986676998569859668987845446484868357544366675664663442351324222\n514534332326346534266533337563667585755748446756667958857768998668688896966968767666669875579687859688676665648463444444655543653332324131414\n254524451566362624264564375565767478686858467669698865865575898766676897869866776978655896585895587577464548666357647444773363626463543525243\n442535342553232452324766737733766654766586446599998558759986696789999986888979678786958675767587564885768644685773746647744236523263621123125\n132353435163463253362437347555447775466875668878957699895866896696797688776976755665688979767558946855664655664647735673646363642224332234215\n544411434126645644532365346776547476665575484447675858959657969557988666686658885687985675696697687857447885537753747464443555663264411513321\n522453242522424234536564676343337437448477676655965589955877786885599977869999665667978558896778875578745485357756766665654466243252135245544\n311335321514544526266363346373737468584456447588878588685665585898775698968888765776865665965888544465544676463675653766545262336263313434353\n133114532551333246364267774777533335788744547488889779897798969997956687865758995966597998666867484456657666377665464544455365522251421334541\n543235341241342556255363464437454673848747446558646966696685696976969767677589995959577868565644886875665443466763455644344563523234552524551\n124253533423343324646465344644577745458576568445854899956987975689676666958758576897696678575874846744586776577557457564345354245232215514132\n123143222524245222633563334735767547575674654484855864999669757599766759857965656665985687448677855645867737777776457354422332365545121333544\n423115121112156626223535444734444774644747477768864568759997695997867587779776568686866655585748544565864665554445445436443243222421555334113\n114345315331322553245525633333373437445658478555657477547688959889675898769769878567964654785448784678557655545377433534366336533152321333352\n311443411311433243325364245765545746633764647646756486667774868867685966669656888787654468488788775874573777563346754326542566466122433122511\n422123135533133553456244366573345664356457566685766674576746644799659898679788957654544874654678577675363664574364725332664243435455345311154\n441112341151122365536244345353477645756577787867666748456648444888558766678677848578778864764845674756665556547533365266255225464531323122554\n421225214342535246262445462245765754636566578574845768874576474465746868576468556774664867757766568564665363465675453563455536622122214423113\n211124134315442145434426353534573777474747637754887466754544854648854655877548467866875455745586786656463737565434663332553464253311444551311\n431231522213423432666636546645345667667457576646647845646466877447487867868654568488575447757888633355455676356533342226426463135342132232122\n344243535312553554443366352244624563434477356763555684665765665665848758657685574545666448587576336737653566343242432654566244125323551452123\n413311554351521241146336335653527357444467566664344664887755576574868664646875566474877578746467445675633656455366365353545515141153325531321\n141214542545313421156663465566426547465676337347737478888676885866665848556558585874585485686667473676667645573236365346443553151232552252223\n241431141325215323114446553553536575354367375343756535886854864675444767488847886874488688757477335774536743446335662645654352352232152224112\n124221344354514445151456322242526533465366443367436753335674866464677676476774564485466633475566774535665345262245346356444221534212523223221\n214322331253522525123243526253225233676444553736454745676756847655776777748657565455634667735737743367653732242362362546325541512135524412122\n412323211532344125133335623443325453244655577737666343377347356486484665665554868374347735564643443677374626536624433542223454132545131444142\n434231124424232443512526324336442433264335557347457634645666653644364484674567677455465746536647777334545636656256644566213512215152432344324\n243242421244311354452343336566646463264466634646344465377454457777767455337336334334764766764544375566352356626663655324435242314231312443322\n113213443233242433321325343642243354654623563375554363434333555373363377757337663544734573443647757474243343445242536552145322235531344233312\n123332214422533544122215455432323635324654423776733766656475344773367444673343634345563335564654373424435355525252265224454311451252314123121\n124313431134252125531353523523245644652634365564553377577637465354777365334336443466765665773334564336326623334552442254153411121524342241413\n234324143433241545353443255346223552364424226666356436536574356757435656577764476764767655344767565446532332225665214342552442555433211341334\n312323144344311435133532221324563665456456432563447547763577536643763673653447437573345475675425245254323343426355554132351435222144312134143\n121422144423412521223324215122265633635454566364565366535446467655573535564477446465776433626445633352226455435321523344433455334333411441331\n314242331324342255423255531333125246365225326434654462357545643457646364647344546457633444443566652235426422623514332245521311342141442241144\n232121214231313214141313552531551622645436432453333462365646543644544575564345743543646232564226455462445552332451224145554321413222422244223\n321134421122441142331552453113433245545536443344642626462242456346465377456744344364433365625322665533563423135253323142212454233342323411323\n323113142331114413254321222442532254465336243522542465624423346344232362254433243566535333235255455332222464445122544135354412214323124144123\n323224442342133212211313422153542511212223664434646665244366332354456255443634334664354626453543454553562315512455242352344244241323224212223\n333332321114241114232313332525141254343352435424534465566442432343346324266622262544442366366252544222624112152112221532453242132424111221312\n331223141441442234231351151535553122255526435235552466364462556554426336544663344333632324566254654234533214425312432431111324123422311132312\n311212323441413234411113254244521541242425323664645246422464562665623562334624345445342423555662444633532122253435515321141333322143221211132\n131133332423142222222424415341342113531412455456522426335443654265423455265256553532642356464446455235143343544414251434424231431443112113313\n312322333111142211432122312512144535525154421251523223243462524664363534663652224565534434532261322212114514555142512413112214342234122211323\n312231123221323441121422312345231312323511521515234233634432652526242262364425625265355633444445555532444122531325211422221412224242232112332\n121221233233443134434313212343514235445545412523315111355342546665463562226345365663452363315453325253332242122143141112144432112112122331112\n122312212122221312121144212412252142515541124225512344425352646354324644536324534363624535325155523143223355344253144132434413211333332131322\n331331223321132222224243414122234453551313424323342524314344151224542263353555352114221425341344144112114234133143411324124441234231312111122\n311221211232312141311314241344442343511534523112424323233522342355244225144235541441233133323142435355554122222431124331123323443223111223323\n112232332313233332344411211212423432212423143212112153351353325435413125241111535125131514325342342215225323143134422244142323223211322321113\n121332333321322114144121241334314444222424354411224234121555221211444345144332155324325554514442234213151142132114321321231324212212332232331\n212332332122311233232144222314232411222234524534114233421351225511124345212225254141552443333145115143533131314323233434442412333333132123221\n211212322131311231132322134223411213222122134314131213545541433115144333341122311253343145513421442314131434121434213141443133322212211323122",
	}
	res := make([]int, 0)
	for _, input := range inputs {
		data := strings.Split(input, "\n")
		for _, mm := range [][]int{{0, 3}, {4, 10}} {
			res = append(res, d17main(data, mm[0], mm[1]))
		}
	}
	return res
}

type d17qi struct {
	x, y   int
	dx, dy int
	value  int
	line   int
}
type d17ck = [5]int

func (e *d17qi) getKey() d17ck {
	return d17ck{e.x, e.y, e.dx, e.dy, e.line}
}

func (e *d17qi) move() *d17qi {
	return &d17qi{x: e.x + e.dx, y: e.y + e.dy, dx: e.dx, dy: e.dy, value: e.value, line: e.line + 1}
}

func (e *d17qi) turn(d int) *d17qi {
	return &d17qi{x: e.x + e.dy*d, y: e.y + e.dx*d, dx: e.dy * d, dy: e.dx * d, value: e.value, line: 1}
}

func d17main(data []string, from, to int) int {
	visited := make(map[d17ck]int)
	queue := pq.NewWith(func(a, b any) int {
		return a.(*d17qi).value - b.(*d17qi).value
	})

	queue.Enqueue(&d17qi{x: 1, y: 0, dx: 1, dy: 0, value: 0, line: 1})
	queue.Enqueue(&d17qi{x: 0, y: 1, dx: 0, dy: 1, value: 0, line: 1})

	for !queue.Empty() {
		q, _ := queue.Dequeue()
		node := q.(*d17qi)

		if node.x < 0 || node.y < 0 || node.x >= len(data[0]) || node.y >= len(data) {
			continue
		}

		node.value += int(data[node.y][node.x] - '0')

		if val, ok := visited[node.getKey()]; ok && val <= node.value {
			continue
		}

		visited[node.getKey()] = node.value

		if node.x == len(data[0])-1 && node.y == len(data)-1 {
			if node.line < from {
				continue
			}
			return node.value
		}

		if node.line < to {
			queue.Enqueue(node.move())
		}
		if node.line >= from {
			queue.Enqueue(node.turn(1))
			queue.Enqueue(node.turn(-1))
		}
	}

	panic("no response")
}