package main

import (
	"strings"
)

func D25() []int {
	inputs := []string{
		"jqt: rhn xhk nvd\nrsh: frs pzl lsr\nxhk: hfx\ncmg: qnr nvd lhk bvb\nrhn: xhk bvb hfx\nbvb: xhk hfx\npzl: lsr hfx nvd\nqnr: nvd\nntq: jqt hfx bvb xhk\nnvd: lhk\nlsr: lhk\nrzs: qnr cmg lsr rsh\nfrs: qnr lhk lsr",
		"ccq: htb nhm srv\nqzd: cmb pfc vxt\nnrd: gkm phj hpg\nkmj: tkd dgn crq mjf\nsdf: zld klp\nrrl: qbr\ntgr: vcs fvr drf qvk kzc\nzdq: kvv\nfvd: czc shs\ncxr: qdk kgr pdr phl\nltl: sfr hhd bfr rpj\nldq: trr xgg gdz\nrcm: kqk pvk hpg\nzfm: rsq zps zzh pdd sbq ktm\nrxj: lhm\nxdr: kcn zkm mpx cvh\nhdd: qsc dfx vcn\nqln: mgp vfh nct ssp\nsnn: tdv fkj vmh lps\nmgx: nqt ptz\ntqz: znz rrx sns\nzkk: nrt xhr nch\nbff: bqc\nkqh: xds vmd\njfp: jgc gmx dzq\nxln: tnf gnm\nqfs: xxb nsk lzt zjt\nklj: nnn gsn mrr hsg\ntrm: rsd zqg fjh qjm\nvnz: lcr bbx vsd\nlkg: bqq ngf mmt\nshp: fxb jlq cvp\npgv: svq kkv\nvpj: xnj xjq knz\ndbf: zkk gth ckf\ndvn: rzl pnx\nvnj: tgz qrv rdx cpv rrc kps\nctr: qdx pmc jfg qbk\nncz: fxv qsc\ndfm: mgp jzp rrv dqj\nkrl: zhs pmc nxh mll\ncrv: cpc rgg rxj kxn\nzvb: fpg mrh\nrrx: zps\nrvp: sxx rkg xvq\nnkj: vrs qtt llp\nbls: lcv mgh fhr xqs\nqjb: mrz sjf kqg cpc\nxhh: xcf tsp mtq qzc\nvzm: nbr mjs lmq hjz\nffj: tkd vpj pcg qvg\nqsl: lgp rvd zhs\nvxc: ljj ckf zdq mxx\nrxm: tzf sjn clp mhh\nsng: fvj ssp fxx qgh lzs\nsrt: hpg qqs qmp\nzsr: hbj zvb\nlrc: kpb\nlcq: lhm zst jnn rgb\nccl: phc lsq ztf\ntjm: bfg rbm xcs\ndrx: fmq mgr xgc\nnpf: mzf\ndng: klp pvn xfx nxb\nvdl: phr\nkkg: qbh fxl zmg\npjc: hlk fjx frx\nqxx: mkz rsj bjz qtl\nskk: djc\nkvl: zng\nlzg: dcg kfx fjg xsr\ngzv: thd kzx hkn\nphg: vhg jgc kmf qxq\nzft: nrd fxv rdm mll\nmjf: ctn zqg vjh dpt\nqst: xrk hrf\ncrq: hvt\nrsq: nht\nqct: cnn tgz prm rdl\nrtb: kps vhg rxl\nfkg: nms\nvqb: jnh gqg ffp rgp\nqsx: fnb xfk\nsdt: rkg rlk qrs nbp zkz\nplp: nzm gvd mtl prj\nffp: cgd xxb\nftj: vms\ndkh: xjb vpj zcd dfg\nqpr: gnr qld ffd zpq\nsnr: djk jfv kpx\nhkn: ncx gql\nnhc: jsg mfl\njnd: sxq dpg rzz\nfmr: cbb hlk qrv ctn xjq\nqnj: mxk\nprf: jnn\nvht: zmq zbc fgf\ncnm: rrz nrf vdl nbd\nggd: gth vpk\nhks: sqv\nfmm: cjt hrf xjb ddc\ndpg: vkx cln\nqgn: bks jfq vzh qvs\nzfp: jht qkg qlx\nnch: rvj rcv\nkgb: gdm lnp dfx hmz\ngqn: mtm htb vxp vtt\nxdf: lvj tps\nmsf: kjx njs djg\nkxh: rzl\nvxp: mlr\nvpr: gjp zhb\nfzc: hqq lhb zbg xcp\nbqr: mvt sgd mth tgv\nzbf: vvm dcg rht xqs\nctn: mzb rmf sdk hmn\njcs: pdd sjn clp rkc\ndxl: vcx\njgc: cnr cbb dvv vmp bvk\nghv: bzv hfk njh\nxrb: vdk\nvmf: tns kcp xvq pgv qlb hzp\nrzl: rgb\nttc: qjz gdv prm vzx dxf\ntmj: tfs rzv rcq dxf\nxfb: tzf spn fvr hzg qrd\nbks: dnp\nrfk: lpb fjx nrx\nvzf: vvv hdr\nhtb: sxl\nfmk: bmp csx tnz kmv pnd\nmvs: vdr\nrkg: slp\nlpn: xgp vsk zmk rrg scf\nnzd: fjj zpp lzt hlx\nnfl: ghk dff\nhgx: nms pjz rcg vzx\nbsz: snq gvd nht jzh\ntfn: zgj\nhgt: qzv dnl rrg vxp\npsb: czj jnh hzp bsb\nrrc: drh pzn sxl\ntmd: ghn hzj jhh\ncjd: cnk\nhvx: rbp qqm\nqcv: ghn rmf pfq\ncct: msc bph\nzbc: zkm rdl\nfjh: dvv\nkcx: xgx zpr\nclx: pdn zst mkp pkr\nxbq: rrr mcf lvf zsq lgv\njfv: xpn stt pph\nmnk: njs qtv\nrgd: jst chx\nvbq: dpt chq\nmcs: sbr hmn bdf\nqkm: xdq dfr qqm\nkzr: kkr hvt vpk flb nch\nksp: nqt tdx\nscz: hgz lgk tbx hds cpr xqj pxf\nqvn: gsd sjn djd\ndqh: xdq fxb mdd dlh\npnz: qkg\ncrg: sbb\nxnk: xkp rms fxb ntj\ndsq: msr\nnjs: fll\nxlf: mft pfn zfk svs\nnnv: jqz fvg\npcn: vbq kcx khn vzx\njfq: rkb\ndsl: bxs\nzrz: hmr cpj shg qsl\ndqb: msq ldf gzv xkj\nccf: gdv rrl qgp\nfcc: dcr gqg htb\nksz: gcx ldq\nvnl: sjs mlc llc rhx pnz\nmcn: rrr jlf zps\nqtd: pkl lbg bxs khp\nszj: qhd mpv gjb nkj\nbph: stz cxf\nqqc: qdk xcx sfr bgl\nppb: jjn mfh msm\ngnt: xgx\nhjq: zhs npc ksb jsg\nljk: gdz mzk phj lvx sqv\nncm: nbd fqr klr cvp\ntns: dxf lvj xlx\nrrg: pns fgf\nfhs: mjm\nxsr: bqq\nsbq: tdx\nlmr: zth\nhbj: qgh\npxf: snf\ncgn: dbv pvf jcr\nzgm: drz dff\ngxn: fpg vbq svf\njzr: rrl zbc flm vkx\nqqs: dfx\nfhk: qfh knv sdk\njxg: lxj djl cxq hcp\nvfz: ssp mjs rjc mgz\nkrf: pzr pfj\ntxp: hxz\nnnn: kdm\nkkr: lvj\nnms: zrv\npcf: gfx bgm hfp cpc\nhzg: vpv drz\nlcr: zpp dqj rkb jtf\nlfz: jnj bbq msc nxh ppb sqs\npjz: ltc ffp rlk czs\nrhm: zht dpt\nmzj: hcb tzj bff nkr xpc\nhnd: vjh qtn\nccj: vcx\ntzs: rbk\ngdv: zpr\nphj: rnh\nqrb: drh ljn\ntbm: scg cnr\nkfq: qqk lmr hhd\nbfg: mqd\nqtq: zcm sqv dqg kmb\nlbm: nhm\nfnl: vlp\nflv: nbv mgk pxp\nlxt: rhm vdr qtv ccm\ntfl: gzj gsn kql\nqrd: ntq hqq\ntfs: jkb\nqdx: fnb\nkfl: gnn\nckg: fgm bqf\ngmd: dnl qnq vfq hnd\nxkj: dln hvx fjx\nlfl: rmp jmv vvm phb\nmth: jnl\nhbl: jcq hcb\npph: cbj\nchq: bjv\nsxm: kdm rtx\nqgp: xxb kcx tbv\nfgb: zmk jkq\nxnc: pxp snf bqc\npdd: bcr kkh\nfrx: fgm\npzv: brd\nlgv: rtl qdk kzc\ngmv: lvd mrr\nzzd: scg gnt lcz hdq\nndx: sbd zdq\nxxk: zth xnh\ndgv: gvt hqr zdj\nzst: rtl\ncnn: drh nkr bqn\ntfk: dsg kgp nkj dhf\nqlp: nxh rnh psm\nbvr: bzl hrk zsq\nmrh: kmd\nzqz: qpg jlj qgr cvh\nmzv: kkz mkv\ngjd: qvc\nkzc: lvd\nsgf: njh gzv psp mbs flg\nmbd: fdp jxs\ncgp: sbd\nhpd: fjj ffm knz xrt\nxrp: lvx pqp glj kkg\nlmn: bqq pjk\nskl: vrl qnb ftj\ngrd: bcs hfh\nhsg: pdn zmf rct\nsxx: ptx hvt\nrkc: sxc fqd nkz snq\nblm: mgg tkd vxp\nnqb: scl zkm nfv slp jkq\ncbg: qxq vht vsk fzl qvg hkm\nfcl: tsp pzf ccj\ndzz: tmd bpj kps qjz\nbtj: rsx cqp vjh rhj\nbvb: jnd tvp tpc rvp\nklr: qpg sph\ncvh: vmp pnb hrf\nltf: hzj dlh\nbvz: vtt knz ppx vhg vzg ddc\nzfj: kgq fjx xqj vjp pxm ckg\nnzm: qbh\nmxk: hlk\nhrf: klb\nrjh: djg tvp sdk\nsjc: rvk vkt\nllf: srl nhc\nfjq: xkp\nffd: qtt glg mlk\nngd: mhv cjk xnj mrd\npkr: brr tgb djk\nfbr: sgd cpc drq zgm\nmss: rdj ftc lbm\nmrf: qgm bqc fjq hzj rcg\njzc: nsz gzj mlk szn sdl\nqkh: ngz qxr zcg nhc\nbjz: xvr\ngpr: bdt jqp gkj jfp\ncdd: xsx\nqmr: qhg vvm hdz xqv\npch: dsl kqb\nhlx: zqg\nvrc: ngf jfv xhx\nbqq: sbn nsh bxb\nmnz: shs lkg tgr vnl\nxrm: fjq jlh\nhqg: rmp xsr xcx crr ktm\nrnt: pzf trr\nxqv: nqc gkm rnh bsg\nbhk: dln vkx\nmvt: mpl fjl\ndgr: sbb nsz xbn dxl\nhcq: dzl ckg rgp lzl rjh\ndxq: fvg zmg pmj\nmps: sjf xcg qnx glj\nvsd: scf ppv rdv\nkzj: qxh fjx zjt csl\nzbv: zgj llt\nfmq: mmq\nxcx: cpf hqr\njht: djc\nnqt: vcv\nbcb: vxt llf zgm qqc\nmhh: pfc tvj stt\nmnn: rbp rrh knp lvj khn\nkng: pvf djc cvl\nqgt: bcp dnp xnj lrx\nhrk: vjx xnh\nnmz: mcr dxg\nhqv: xtl mlt vqc mqr\ngcz: tzm vzx mxk\nzxn: pjm hdz shj crx\npkk: lpz nht fbl kgp\ngns: vvt sdx kgt\nkpx: tzf kxh\nbjt: snf psp knx rdx\nvph: ggr\nfph: dkf bfq qnx jcp rhx\nsrc: fxx pzn rxx\nlmq: mcp\nptk: mrs hxz jxs fph\ncpf: phl vcv\nxvr: bzb\nvsn: kzx hgf xxb\njkx: gdm qbz\nphr: vtt msq\nbgm: dhf vpv\nccm: fqr nrt qcv\nqhd: gvt zmt dxs khp\ndjd: jfg sfp\npcr: hgz dpt hfk ffm ljf ncx\ndmc: vqc kjx\nzmt: hlz mkz\ngsl: frx cps\nlvd: jnl\nxzk: ldq glg mmq\nmcp: dvv\nbzv: jqp fgb hhr\nxdc: sxb dlh ddf\nfhr: vcv xhx\nbrr: fqn msm zgj\ndgg: xfx vzm hbl dsq\nmrz: xgl qqs xfb\nkfx: crx vjl cpm\ncft: mrp jjs scq\njtz: qxr kms kvk\nhgf: jjt zqx\nrcc: rjz sjg xjm cvp\nqxh: bdf\nbrn: msr qgh ggd mlt\nnxb: bqn htb\nqpl: bqf\nxlm: rsj vrs\nfjg: rtx fzm\nvcc: rjc\njpn: fvn nch\nkps: jjt\nztl: bks dnp\nlkh: xrt prm pfq cbv\nthd: xdc jlh\nxpc: fjh jlq scf\nddm: gdz xfk\nllz: zqx cbb mcp ntj zfb\nqtf: trr fpd fqn csq\ntvj: jkx pmj vxn\nqlh: llt\nnrq: nlg\nlfr: hjb mtm\njhh: fnl\nljd: mqd\nknv: vfq xrr\ndsb: skk gpv xkb fnb\ndsm: fgq fgg frf gcv\nxcg: kkz\ntct: jzc ddm rht hqg\nnpr: vxb nnv nmz xrx vmd\njxp: jdq\nblz: mlc lkd kpb mzf\nkmv: fdf\nppx: knx cdd sld\nntj: fvn\njfh: vmz zsn\nddf: ppv\ndhq: bbx vbq kmv\nmcr: vjx kkz qxr kpv gnr\ntnf: nqc\nxgp: mqt\nknj: blm lrl rrz qpl ncx\nssk: rdx\nfjl: rpd ntq xmd\ngsj: rmf\ncpm: sxc sbb\nxzl: slp qlb mtz fjx\ngsn: qqp qtt\nzzv: sxb\nkbp: xhh bzl rbm gvb qdk\njzp: knp kgq fvn\nxlk: fpj fhs ccq\nrgp: ttz cnn\ncxq: gnn ptn bzg\nrxf: lfz qbh xlm vjl\nkms: zng fvg\ngnc: vph lrc fdp\nvcs: ddm vzq bcr\nnsx: gcv jmt mxk\nrpt: vzq\nfzk: tqv xgt cln\nlrx: lpb\nhfj: xrk gsk dfg\npdr: zcm gsm\nzhl: ftl qqp pph ngf\nmxx: dcr fps\nlhj: kkr mpx rjq dmc\nnrx: czj dxf\nqsc: brd sjn\npzr: mgk\nzvq: zhs ntc pjk mmq tff vqg\nhzj: fjj\npnp: dkj\nkvh: kkh gpv dxq nsz\nnsc: dcr njv bsd xjm rxx\nprm: bjv\nfvj: vsn hcp czs tqv\nmpn: nrx tfs hcb shp\ndjk: mkx gpf tff\nfgp: nlj lmn\nxsx: dsq\nlhn: mvs smx\nzqg: nhb\nqqj: dvv lzs fvj vnz qkk\ngvm: hds fqr dzl scg\nrsd: zvb mff\npvf: zng\njnj: bfg\nshj: phb lcv dxg rpd\npqg: msm gfx lsq\nqvs: rcg fgq\ngpv: xkb\nnjr: rdv zzd dnl rjq\nqjz: rhl gtd\nkjg: lkg rsq rfz\nldf: src klp hlx\nrxc: rrx fkx\ngmx: ccm vdr\nlkm: stz ccj ccl mfl\nzlx: ggd nbd\nmlq: zst lhb jxs cgn\nqhg: ftl\ndrq: zcg zbf plg\nxbg: cmn xzl jfj msr\nvpv: dkf fkx\nqbj: crg mhd mkz\nmxd: tzs rsj zkp\nsld: mff czj\ncvl: gkm jht\nllc: ncd xvr mqd\nvbg: pvf qtl bjz bgm\nfxx: mqs qkk\ncps: knp kmv\nmkm: rmf rsx smz bzp\nvgg: sbc gmv ggr zgm mzk\nlrl: jtf xxz\nntl: hmn lzt\nqnq: rbp flg fhs njv jdq fdf\nknz: ssk\ntmb: zdq kvv xjp tvg\nvxm: kkp hqr crx\nfsn: bbn hfh pzt jrk\ndfx: hdr lsq\npns: ffn kcp\nhgr: mvs xrm nfv njv\nlvr: rpj gxx sxm hmz\nkss: gzd knz fkg nnl\nqbq: kng gnm phz pjk\nkxp: klp htv fdf jlh\nhxj: dqj zsd rxl cnm\nprl: tgv gvf bjz\nrjz: bsb ckf kps gnt htv mqr\nqlb: zfb ghn zkz\nznm: qrz rfk fkg rqc xjm hbj\nbcs: sxl htb\ncsq: jql vmd\nznz: vrj\nqtl: xcf cpf\njrp: kkp qgb nkg cxd\nvtv: nnn rrr fvd vzf fnb\ngtd: jnh hzj\nrvk: zps zfr\nbfq: kpv\ntdt: fhl tvp hlk jlj\nflm: tsj\njsl: xhr\nnqn: xtv jqp flm mss nbr\nssv: nms vsk pzr nsk\ndjl: ffn flm smz\nnbv: cdd gqx\ngsd: gqb\nrzj: sqv jql\nbsd: trl msh flv\nlhg: lft xjm\nktz: jfq fcc vxp hrs\nzkz: lhg gmc\nkgr: pzv dgs dkf\ngqg: mqt rbp\ngfx: gjp ncz\nvzq: mlc\nmtq: nnn mth kkg\nnxj: dkv qmp xxk\nbjr: lmn xln vpr\nbhz: mqr qbr rgd\nghn: fpg\nczs: vjp rjq\nzsq: qvc kqk\ndzq: sxl\ntkp: nxh ngf nhc txp\nsqs: nrq\nfzp: qbr thd htv kht\nmjs: rsx csl\ngjb: lrc qpj\nhds: fvn\npqp: nkz fkj\nqnn: rdl xhr\nqmp: sjf glk\nspn: lbn mhd qgb\npgt: zfb trl qst lpb\nvkv: pnz llf stz tff sqs\nmtn: mrh xdf\ngjx: pzn cjk zzv pkt fhs lhn xrt\ntdb: hdq dln hzp\nbqn: ghn\njxs: qxd\nlpt: zrv xjb fzk nbp zrr\nvrz: cnm srb rrh\nlvv: vms rff sbr\nxgc: dsl\nnsz: krc\nrlk: njh snj\nmlv: gjp trr prl hvv\nqrz: drh lrx zcd\ntzm: bqf mcp kmd\nvpc: cnk mqs jkb fll\ncmg: mvt hpq znz rct\nksg: mrj ntq bqz mms qzc\ntlv: gqg xdc dft\nzcs: zmq fgg rhj\nfhl: rsx vlp dpg jfh\nmhq: glk kpv jnn cpm\nxvx: fml ccj fjg cxf\nsss: dxl gvf zpc pnx pzr\nfsq: ksp bjr\nvxb: cxd djd qsx\nnrc: qxd xxk qzd xln fqn rjk\nbsg: zcg\nsdk: xgp\nhpg: gkm\nzlf: mlc ghk bgl glg\nfxl: qlx ftl nqc mtl\ndrf: mrj sbb\nzpc: vvv flq pfc hdd\nsvs: jsg znz\ncgj: flb rms\nfht: mtm\npxk: srl ghk rsq pzx brc\npnx: vbj\nzvk: jlf lhm sxx sbn\njcq: zrv\nztf: tff vzf\ndzd: jfj gcz\nxmg: pns pnb fht svv hrs\nvbh: lbm psp\nrhj: fht\nzfg: pqp gjp sbq npf\nkmb: sxc\nqzc: xcf rjk\nqcb: pnz nkz sjc jtz lxz\nlnp: mft hmd\nxgr: ttz rrv tpc gcz\nlps: npc gcx vcx nrq\nzkp: vbj\nhjb: mtm\nrgg: mgr nhx mbd\nrbm: xcn\nrfd: hrk zfk phm\ndrc: jsl lnl zcd\nfss: jtf scl pdm dvh\nnfv: ncx dvv\nspl: xnk zqx smz nls\nmpx: qbr qxq\nmpl: prf lmr\nhtl: rqc chq vbm xgt\nkpb: vkt\nsns: srl tjm kfq\nsnj: ppv rqg\nrxh: psp lzs\nnsh: kkz\nzln: kkz dcl vjl\ntnc: cnk fhl kbs csl\nzsc: jvj sfg xvr qbj\ncbv: fps vml mvs\nxjq: chx ntl\nzhb: tfn\nvbj: mfl\nvmz: lcp drc nhb fgm\ngrx: cxf kpb zth fzm\nxcn: czq fqd zhb\njns: mkv\nckr: hbl fdz mtn rsd\ntvg: trl fdz qtn\ndxg: vjl\npvn: rcg hzr zsn\nhdq: mxx vml\ncpj: nsh\nrld: rjz qxh kzh pmn\nnrt: bqn\nbzb: ftl\nlcv: kpv jnl rxc vph\npxh: hqq fls xcf rvk pmj\nhvt: bsb\nnrf: rcq zxv ptx\ndrz: msc\nfdc: bph gxx kqh lcq smm\nqjm: vms pnp jfp\nzfk: rhx pbx\nmrd: dft zmq\nrcq: pvc dzq\nsfp: sxc\nhxc: ktz jxp qbl cgj zcs\ndbh: qnb bhk trl\nqkk: tps vmz csl hcp\nnct: zdx vzg dmc\njlj: chq vpk\ntxh: dvn sjc gmv zts\nqbg: rhx mcn mcr\nlpz: tsp rhx mgr\nmhf: mnv krf nsx rkg\nqrv: nbt bjv\npxp: rzv ptn\nfpj: nct\njqd: rlp tpg hzd fzv kgt\nnfm: thd vcc qgr klb jzr\nvhg: djg\nsfd: mnk nsk pns fps fjh\nsxj: qst vbh smx\nfnk: sph xtl tqv xgx\ncbb: vlp\npzx: stj rbk zcm\ngvf: shs\npnm: lzx gnt\nhgz: fht vdr rsx\nsjh: cpf tfl vph hdr\ntbx: vbh kbs\nrzz: ssp sbd\ntpc: sdf xrt\nkjx: bqf\nfzv: mrr rvk sxc\nhdz: zfr qbz\nqxr: vdk dvk\ntdx: hzd\ntlb: djl mss sxq qkm\ntmc: mgg bdt nct zsr\nnzc: rlj zgj qkg rtl\nddd: kkg nlg zfr lnj\npfn: rgb\ntmn: lbm zzv qgm ctn tps\nnkr: rcv\nszn: vrs nlj fmq\ngvd: ksp dcg\njqz: gqb\nqnl: sxm pqp llt nxj\nspq: mcp qqm lhg jgz\nsfb: ngf qbk zsq vcn\nzzh: sbn sjf kmb\nxkl: vms zpr vxg svv\ntjk: rnt mkq gjd glk shv\nctb: xgc jht\nhfn: zpq qbh stj fvr\npjm: zbf kqh qpf\npmn: bqn csl mtm\nfnb: lkd\nqkt: sxb grd sxx svf mrd\nxns: qpj pqp kkp mlc\njhf: sdf kcp dzl\nkhf: rsj\nzpp: mlr\nnbt: pfq jgz rkb\nzkx: xnp xhx fqd\nrff: jlq dzq pxm\nfgf: fpg\nmgk: hbj cgj\nbfr: gjp fqn mms\nxrr: nbr\nhcp: svq\nxpn: pmc sbc fgp\nmnv: zlx klr zcp\nmmh: bsg fpl lvf snr vrc\nxms: nrq bzb nsh\ndtz: nmt xgt dnl mjm mcs\nscq: kqb rbk\nthl: nqt hlz kdl\nqpj: vdk pbx bzb qvk\ngqc: jns fmq lnq sbq\nhlp: bxb vvv bfg tnf npf\ncrx: mth txp rpd kvl\nrkq: hnd bkk vmz pnm\nxpb: rjq\nbnb: zbf qxr vcd jnj\nhfk: jtf\nssq: qrb nrt xsx\nvqc: zkm\nmcm: ldq zbv srt tfn gnc\nmms: lpz mkv\nksb: sfp\nnkg: zmf mfl\ndnb: mzf jff shv llp sbb\nvhv: ssk rxx cgp\nhqs: bjv qnn vfq\nbsx: kht lsj czs zcd\npkl: gvt stt lnp\nxrk: jxp gsj\ncgd: pfj xdf\ntvt: czj fpj\njmv: bsg nht\njsj: vvp fbl xcp qlp\nfdm: cmb qsx nzm vzf pch\nbxb: dvk\nnjz: rsf phb\nzfb: qnj ljn hfj\nxnj: nbr\ngvc: xrb bvr qxd gjh jzh tsp\nqkg: stz tdx rtx\njnc: qtn ffn svq rlk jcq\nlhb: dff xbn rtl mqd\ndkf: ljd\nvdr: lft fgm\nnsk: dzq\nxhx: tfn\ngvr: dbf gqg sld vjh\njjm: fqd xrb zng fph\nsbr: cqp\ncpr: crq dzd\nhfp: ptz jns gjd\nglm: qzc jzh msm rgg\nglj: fdp ggr\nqtv: kkv\nnjx: hkn pfq pbx htv lzt hmn\nxnp: hmd dsl\nhqj: jhh gxn gzd\nhkp: kvl jxv czq lsq\nqpf: hmz llc fhr\ngth: hkm xgx\nvrl: zxv\nfrp: gdz brd\nphb: ggj\nflq: ghk zln rvk\ndrh: hfh\nlgk: fxb\nvgh: vdl vqc qrv cgd\nlrd: crq knx jhh\nrqg: bzg rkb\nsxb: vzh\ncmn: qcz ltc btj\nzql: vfh xnc bsx qfh jsz\nvtm: cvl ffl\nrdj: knv chx klb\nvjh: zxv\nkhp: pfc\nfcm: knx rqg\npcg: flg xxz\nrtr: hmd mfh drf brd\nclp: jnj ksz\nrsm: qlh hmr zvs xlm\nlfd: zkp gpf mcf kqk\nmdd: pgv rrz\nbvs: hrf bvk lzs xnj\nzpq: mrj xds\nfkz: jhg\nsst: flq fnb rsf rsn\nphz: pnx drp pnh xgl xmd\nrct: vrj zfr\njgz: qxq sxb zlk\nvfh: tfs dfr rgp pnb bvk\nbrc: rnt nmz ggr qbk\nvvm: rgb jvj\nztm: vpr xnp qdx qlh fgp zkx\npzj: skk zbv dkt vxn zng gnc\nsxq: tzj\nkht: vzh gmc nxb\nsjs: xkb fls\ndfg: bjv pkt\nhdc: rdj brn dng lmq\nqbf: qnn\njjn: qvc lrc\nbvc: xpc ptc xgp jlj\nrxl: gqx zqg zht\nkcq: lkh pnp vfq knj jnh\nrht: tvl\ntzj: kvv\njff: zbg sxc kxh\nttz: zxv rzv ntl mqt xtl srv\nmcb: bzl ggj gjb mtl mkq fsq\ntlq: jlf bls sbn gsn\npmj: pnh\nlmd: mbz fsq cbj\nknp: vrl\nrpj: qtt vjx\nscf: tps\nhmd: srl\nnlj: sxc hrk\ntxx: hqr gpf kkz gmf\npnv: bqn rrh jsz\nvxg: jqp msq\nrxk: qnq dkj mmp qrb\nxrx: zcg rmp kpx mkp\ntpd: fpl xsr xcs mgx xbn\nsmm: njz gvt nkj\ntrk: mzf kzc nqc xpt ctb vxt qxd\ndkv: gqb nhc\nptz: pnh qbz zmt\nngn: bmp hlk fhk qpg gkj hzj\nkpq: fps rrz scl cpl\nbbn: xxz bbx\nfvr: zmf\nfdz: jpn mqt xsx\ncmb: gvf bbq\nqgr: bzg frx\nmkx: bbq zbv\nvxn: hjv hks lvx qlx\nzmk: kbs fnl\ncpv: rtb zsd hzr\njlh: xpb\nxjp: jsl mnk slp mtn\nvvt: bzl tvl bqz\ndvj: nhx czq qlh qhg\njxq: qxz qqp xfk mlk\nnxh: khf\nxxb: lzs\nptn: lxj\nzld: smz kmf\nlvf: vcn hjq jff\ndcg: xcf kxh\nllt: jlf\nfgg: qcz mgg\nbmp: qpl\nkdl: kvk bxb phb cbj\ndzl: smz\ncqg: lfr djg qjj fkz\nlxz: phj rpt\ngsk: vjp snj ddf\nstr: mkx jns tqz vvp\nzts: pzx fvg sqs\nqnx: jfg hhd dfx\ntqn: ztf msc cxf\ndgn: zcd\nktm: vmh kqb\nffl: xxk kgp\nxqs: vcn dxl\nrcg: zpr\nmhd: stj khf zzh\ngnm: xgg hxz\nfkx: dsg\nsdl: rsn fxv\nmzk: zcm vrs\nlnq: gjp njz zfp\nhrs: ddf qpl\npzt: kbs pnv\nzcp: ckf zsn phr\nmkp: dgs jkx\nlsj: flk jst xxz ptc\nckf: sbd\nhsj: psm gjh tnf gkm cct\nqls: fls llt\nlnl: srv xpb gsj\nzdx: ggd mlr jsl\nfml: rpt psm tzs\nlgp: dvk llp\nzqx: chx\nklp: srv\nrlj: xcn mlc tvl\nmcf: txp\nxpz: htv pxm flb svq\nvvb: dhq zvb rrl lft\nhjk: pjc fkz tsj knz\nbqz: fdp\nzmf: kqk\npfd: qhg lnj rzj gnm\nskp: cpj dhz fdm\nncd: lgp bxb\ncsh: gvf dvn hjf vkt qtq\nglk: pfn\nrvj: bzg\nzds: tzf ksz rsn hzd\nxkb: khf\nxtv: jfh vht nbd\nnmt: svq\nfjn: mlt dnp xgt\nvcd: hbh drp xgl gxx\nkql: vcv lmr\nphc: qlx jmv lhm\ngpf: ggj\njrk: hqs vqn vxp\nqvk: mfl\njcp: ffl rrx\npvc: zdq\nnls: jdq\nkdm: vbj\nqqk: qkg gqb\nmll: kql dqg\nqvg: tkd hjb\ntbv: vdl rff vqc bdf\nffh: qgr tzv xhr ckf\nfpl: pbx jqz\nxqj: fjj mtm\nxsq: vkx khn rcq kmf\nkzh: ptc lrd\nmbs: bhk lzx qfh nkr xdq\nbcp: bzg smx rjc\nzdj: mfl\nvqg: gjh rxj\nxkf: vjp htb zrr mff\nkmd: ffn\nxzt: hks gjd fvg lvd\nghh: pzt ptc vmz mzb cnr\nqnb: bff kkr hjb\nmgj: mtl tgv rfz cvl\nnvv: pqp nhc ldq jfg\nmzf: ksb\nftc: pxf bdf pxm\nbkk: tzj hvx\nqzm: xrr nbt jkb\nfpd: sst ggr xbn snq\nltc: ljj qnj bff\nmgz: cvp vkx vmp mrh\nfbl: gzj\nrrv: lzx dgn\nxhr: ljf\nplg: szn vzq pzv dcl\njkq: ljf kgq\njzh: jsg\nhlz: mft\nhgs: svs gsd glj skp\ncjs: mgx gpv thl qls\ndxs: lkd\nsjg: kcx xtl pfj lrl\nbgl: gdm\nscl: tzx\npnd: qzm gvr\nfbz: tvl gvj ncd zdj\nmkv: gxx shv\nzgp: sbd mpx xsx pvc\ndmd: qcz xjp vxg ndx\ndhf: zkp gdz\nhmr: prf\nstt: fls dqg\nvbm: bks qgh mlr\npbx: lsq\ntzv: jmt pnm jcq\nqgb: pdr rsn\nrcv: jjt dcr\nnjj: gqx hqj ppx gmx\nvqd: mcf vdk zfk\nmgp: zlk pdm\npjd: kzr dgn skl kzh\nvzg: rdl\nklq: nnv kdm bxs pzf\nhbh: lkg glk nzm\nhjf: fkj xsr fvd\nzsj: mfh frp vqd dhz\nrrr: djc\nzbr: xlk mnk rzv rjc\nxpj: ptx cbv hcb fks jpn\nqxz: kvl gcx\nkhn: hzr\nhjz: csx vdr\nxdn: grd ltf vsk\nmkq: shs\nrpd: xcs\nfzm: lvd rzl\nxfx: cjd rrh\ngtk: mhv nbv fkz ljj\npqr: kkp qpj zdj xms\ncsx: fps nhm fll jxp cqp\ngvb: djc qqs bgl\nlcp: ljj pkt hkm\nzjt: jmt zlk zld\nhjv: xnp hxz mkv vmh mfl\nxbf: rcm glg csq dsg\nrsf: xcg dsg zng dkt\nfll: sph rms\nmqs: lvj bdt\nmmt: gqb phl\nxfk: dkt\nhxr: rzh dfr spq dfg bcs\nnpc: kvk xfk\nhhr: sbr bqf hzp\nkqg: rht xlm rtx\ncxk: gjp llp kxn fzm\ntvp: ghn\nfxs: xkp hzr kkv\nmzp: fjn vpj dbh btj ggd\nzcz: cvh gsl zzv snf ttc\nmgh: njz ljd mzv\nqbz: vmd\nlrq: mpx psb kzx cbb\nqjp: xvq qgm krf hfk\nkqb: mrj\nrnh: lsq\ngcv: smx rdx\nftf: vpv mbd sxd crn\nkxn: mxd prj\nhtj: qvs fgq bmp\nqbl: dkj rxh tvg\nkmt: vcc bdf mdd\ncjk: qst vht xnk\nhrv: ljd qqk rrx rxc fmq\ngkb: vpc gsj tdb xdn ftj flk\njcr: bxp bfr hmr dxs lgv\nmdt: nhb rjc hzj bbn fdf\nmmc: fgf fcm zlk qbf\nddc: flk jfj fgb\nffk: ksb fkx sjh rfd zfr\nrtx: jnn\nlnj: ljd zhb\nvkx: fps\ngdm: xvr bzb\npdm: msq sxq rgd\nsqc: xgt cdd ppv nct\nmmp: ffp lnl qnj\nphm: sjh ctb npf nht fhr\nlbn: prj sfp jmv psm\nhhd: xds\nhvv: gsm kgt vxm\nvxt: pfn\nntq: hdr\nbbx: smx\ndql: nrq zgm nqt kqb\ngsm: zhb\nmsh: kjx zqg jjt\ncpl: msq zsd rxx\ngmc: pvn nms cgp jkb\nnkz: pjk\nmtz: xnj fxs qxh\nxcp: cbj ncz\nrzh: mlt xnk rhm dln\ncqp: rvj\nlzv: jql drz dkv vjl\nrjk: cxd rxj\nsdx: bfq\nrmp: rbk ggj\nqhb: zqx ztl kmt fjq\nnxt: pcg ftj flb svv cjd kfl\nvqn: hcp lxj vtt\nlzb: pph snq dxs qqk\nbxx: hcb gzd rxh lhg\ncvc: qpg htj cnr nms\ndhz: qbz zcm\nptc: xrr zht lgk\nzsd: qjz vxp\nzht: pnb\ndlh: gzd\nzvs: hzg sjs bqz mkq\nzmg: mmt kpb\nnhx: vrj\njdq: kmd\nlcz: htb pnb\nxbr: rpt fbl qvk nkg fcl\nsnf: rms\ntpg: lnj tzs jvj\njtj: ljn kzh lmq vzg nmt\nfqr: qgm bks\nzmm: sdl krc qbq qbk\nsxd: kjg lgp xgg gmf rsj\njsz: njh vbm\nqld: crg dgv xnh\nlzx: tbx ffn\nfqm: qqm pnd djg ztl mqr jst\ntgb: jvj hzg nkg\nlzf: qvc rzl gpv\ngnr: nfl gvt\nfks: hgf lpb dcr\nljn: dnp\nngz: xfk pph vmh\nsgd: mbz vtm xds\nkkh: fdp qqp shv\nqzv: qtv xnj hvx\nntc: gmf xzt kgp\ndvh: zrr dqj vlp\nzzx: hzd pqp drp gsd\ndqg: phl\ndbv: sdx kmb pnh\ngzj: gvj prj\nfrf: rcv vpk kcp\ncln: flg qcz\ndkj: vhv gdv\nfnd: kpx bjr dxg fjg\nzbg: kgt\nljb: nfl vmh rlp gmf\nsdz: pvd jfq jhf kvv\nfgm: nhb\nmfh: mlc vkt\nrxs: fcm fnl nls jlq vzh cpr\ncjx: drp pch zbg plp frp xrb\ntsj: fvn\nrfz: qls\nnbp: fks rmf\nflk: bqf\ndxj: gqg sxj hlx xlx\nmpv: dcl mft ntq\nffm: fxb zmk\nptx: fgb\ntgz: rrh lxj\nbvk: dqj\nxpt: qbg fxl jzh\nthn: dgs crr fcl jql jjn mmq\nqrs: xkp gql jst\nnjv: lvv\nbpj: smz xsq gtd\nhfq: ljf ccf cgp ssq lrx\nbcr: jht zfk\nvjp: tbx\npdn: stt fkx bbq\nlsb: hlz tqn fkx xcs lvx dvk\ncrn: qdx jqz zfp mzv\ndfz: svf gsk fpj tvt xrm\nqfd: mkz xgc hkp gvj\nzcg: vmd\nqjj: lgk csl xgp\ngvj: xnh\ncrr: xzk fgp\nmbz: pmc lkd nlg dsg\nlbg: vcx czq vcs\ntfh: rlm tlv mgg pnp\nvjl: nlg\nglf: fxl plg gsn mfh gsm\ndfr: ndx\nlzl: bhz pzn kmf ghv\nkfg: zlx mrd zdx svf\nmgr: fxv\nbxp: rsj gcx mpl\nhms: khp xpn ffl gqc\ntdv: hjv djk dkt\nbbb: gnn fzl xtl xjb\ndgs: gdm\nshg: vqg cft kms\nmlk: ljd\njhg: ddf vrz xlx zkk sxl\npvk: xzt jnl gjh\nzsn: drh\nrjs: crn cpj vtm rfz\nsbc: bfq cct\ncjt: jfj ptn bzp rvj\nrdv: zrr\nmzb: dsq hkm\nxmd: pzv xcg nhx\npvd: tqv vcc mrd kgq\nvml: rrz tzx xgp\ngqx: bsb\nbkc: zmq qgt lft hjz\nrlp: prf vcv\nsfr: scq fzc xgl\nstj: tdx\nlzz: tzx msr cjd lfr xvq\nxfp: tvt hnd kpq lhn\ngql: fll\ntlk: fps fgq gsl zsr\ncnk: hrf\nscg: vxp\nmrp: dxl llt lvd\nrvd: kvk mlc cft\ndft: zrv tzx\nkrc: xnh\nhpq: rzj fqn gns lxz hfn\nmhv: qgh rhm\njpb: qrd ktm hks xgg\nrhl: njs tbm bhz sph\njnt: dzd gql ltf nls rzz cps\nknx: zpp\njjs: rxj drx fkj\nsfg: qxz vvv lzv\ntgv: skk\nczc: hqq xkb\nbzp: rdv xpb\njxv: phm rtl dff\nsvv: gkj ptc\npfj: gsj bdt\ntnz: ssk qxh hds trl krf\nnnl: msf fkg qfh\nmff: gth hlk\ndcf: bqr lzf djd crg zfp\nqds: ntj zmk kfl tsj\nsrb: fxb vrl pkt\ngnn: hfh\nmjm: tbm vmp qbf\nvvp: gsd phl\nmrr: cxd\nqtn: xjb\nkzx: bqc xlx\nfzl: zmk kkv\njvj: pmc\nrqc: rhj ffn\nvjx: zth\nrdm: dcl rbm czc\nrbj: lmd qvn pkk xcg\nrlm: jkb pxf xdq\nphx: jmt gkj ckg fgq\nflj: pnx pqg sdx vcx pzf\nhmz: vrj\nkcn: kfl jxp nhm kmv pvc\nghk: krc\nlpb: vtt\nggm: jcp cpc msc mzf\nglg: ghk\nnhp: klb nmt qbf bkk\nssl: ssq gkj lcz mmp\nmrs: bxs rpt mrr",
	}

	var res []int
	for _, input := range inputs {
		data := strings.Split(input, "\n")
		res = append(res, d25p1(data))
	}

	return res
}

func d25p1(input []string) int {
	m := d25parse(input)
	pr, partA, partB := d25getPriceToPartition(m, 3)
	for len(partA) < len(m) {
		var toAdd string
		price := 35000
		for node := range partB {
			partA[node] = true
			p2 := d25getPrice(m, partA)
			delete(partA, node)
			if p2 < price {
				price = p2
				toAdd = node
			}
		}
		partA[toAdd] = true
		delete(partB, toAdd)
		pr = price
		if pr == 3 {
			return len(partA) * len(partB)
		}
	}
	panic("failed to find")
}

func d25getPriceToPartition(m map[string]map[string]bool, limit int) (int, map[string]bool, map[string]bool) {
	partA := make(map[string]bool, len(m))
	partB := make(map[string]bool, len(m))
	visited := make(map[string]bool, len(m))
	i := 0
	for node := range m {
		if i < limit {
			partA[node] = true
		} else {
			partB[node] = true
		}
		i++
	}
	minPrice := d25getPrice(m, partA)

	swapped := true
	for swapped {
		swapped = false
		var toSwap [2]string
		for node1 := range partA {
			if _, ok := visited[node1]; ok {
				continue
			}
			for node2 := range partB {
				if _, ok := visited[node2]; ok {
					continue
				}
				delete(partA, node1)
				partA[node2] = true
				price := d25getPrice(m, partA)
				if price < minPrice {
					minPrice = price
					toSwap = [2]string{node1, node2}
					swapped = true
				}
				partA[node1] = true
				delete(partA, node2)
			}
		}
		if swapped {
			delete(partA, toSwap[0])
			delete(partB, toSwap[1])
			partA[toSwap[1]] = true
			partB[toSwap[0]] = true
			visited[toSwap[1]] = true
			visited[toSwap[0]] = true
		}
	}

	return minPrice, partA, partB
}

func d25getPrice(m map[string]map[string]bool, a map[string]bool) int {
	cnt := 0
	for node := range a {
		for node2 := range m[node] {
			if !a[node2] {
				cnt++
			}
		}
	}
	return cnt
}

func d25parse(input []string) map[string]map[string]bool {
	m := make(map[string]map[string]bool)

	for _, line := range input {
		q := strings.Split(line, " ")
		q[0] = q[0][:len(q[0])-1]
		if _, ok := m[q[0]]; !ok {
			m[q[0]] = make(map[string]bool)
		}
		for i := 1; i < len(q); i++ {
			if _, ok := m[q[i]]; !ok {
				m[q[i]] = make(map[string]bool)
			}
			m[q[0]][q[i]] = true
			m[q[i]][q[0]] = true
		}
	}
	return m
}
