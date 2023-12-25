package main

import (
	"fmt"
	"strconv"
	"strings"
)

func D19() []int {
	data := []string{
		"px{a<2006:qkq,m>2090:A,rfg}\npv{a>1716:R,A}\nlnx{m>1548:A,A}\nrfg{s<537:gd,x>2440:R,A}\nqs{s>3448:A,lnx}\nqkq{x<1416:A,crn}\ncrn{x>2662:A,R}\nin{s<1351:px,qqz}\nqqz{s>2770:qs,m<1801:hdj,R}\ngd{a>3333:R,R}\nhdj{m>838:A,pv}\n\n{x=787,m=2655,a=1222,s=2876}\n{x=1679,m=44,a=2067,s=496}\n{x=2036,m=264,a=79,s=2244}\n{x=2461,m=1339,a=466,s=291}\n{x=2127,m=1623,a=2188,s=1013}",
		"vvd{a<3062:hrj,x>1762:zgl,s<881:dlf,A}\nckk{s>862:R,a>3127:A,A}\nth{a<3804:jjp,a>3923:mfp,a<3866:ck,fgv}\nsm{m>453:A,s<2996:R,A}\ngj{x>1334:R,R}\nszs{m>3306:A,a<2977:A,A}\nhk{m<2759:tbr,s>3402:A,cn}\nhrj{s>993:R,R}\ngd{m<1654:pb,zk}\nzb{x>2892:R,R}\nqd{x>926:R,s>1347:A,s<1075:R,R}\ngb{a<3731:tt,x<2259:tdk,qxg}\nnh{a>1660:A,a<1559:A,A}\nqb{x<952:A,a<824:R,A}\nhb{x>958:R,R}\ntl{s>3339:glg,m>3104:R,bx}\ndv{a<3358:A,A}\nkgt{a<2746:R,s<590:A,x<961:R,A}\nfgv{m>1968:R,pcs}\nkzc{x>1543:R,x<1255:lrn,lb}\npz{m>3096:R,A}\nfbq{m<3544:fq,sdh}\nvp{m>1213:qgj,a>2889:dn,m>629:A,tk}\ncpv{s<1222:R,a<3006:A,m>485:R,R}\ncm{s>1201:vbl,x>1442:lv,ff}\nms{x>1035:R,A}\nnqf{a>3776:A,x<2712:R,A}\ntns{m<512:A,R}\nddj{x<982:A,R}\nlkm{s>1761:A,x<3406:R,A}\nrgv{x<1838:R,A}\nkk{a<3285:A,m>3883:A,a>3614:A,R}\nsb{a<2998:A,x>1425:R,m>399:A,A}\npsn{s<2782:R,a<2995:R,A}\nxn{s>2939:slq,m>2356:fgs,rj}\nflx{m>502:A,x<3413:A,s>1324:A,A}\nhqb{a>2426:tx,s<3349:A,R}\ntq{m>1617:dvq,x<2295:jsh,a<3368:rzc,qms}\ncn{s>3099:R,A}\nlft{m>726:R,m>598:A,R}\nkt{a<2349:tdv,x>1126:hnz,bps}\nqms{m<1499:lvq,a>3479:R,R}\nns{s>1283:R,A}\ncj{m>3441:ql,m<3391:A,R}\nrcm{x<1000:R,lrz}\ntqd{x>1200:A,A}\nlkc{m<1761:A,A}\ngzf{x<3613:A,a>2771:R,A}\nnzt{a>306:A,s<973:R,A}\nfth{m>355:R,s<3285:R,A}\ncmg{a>3790:R,m>1001:R,A}\nvhf{s>1126:A,a<3520:tsb,R}\nmpm{m>869:qh,a>3099:zq,dff}\nlq{x<2632:zpc,tct}\ndlf{m>1614:R,m<1126:A,A}\nbkk{s<3105:mqs,sg}\ntct{s<984:gnm,m>1073:jfz,s>1514:dc,xj}\nlv{a<239:A,A}\nfb{m>2007:R,A}\nfpj{s<1828:xz,zt}\nkl{s>2400:A,x>1193:A,x<671:A,A}\njb{s>796:qt,zn}\nqs{x>2260:A,a<3886:jxg,a<3903:pt,A}\ntc{a>2840:rqp,m>1035:vl,hq}\nfss{s>1221:fld,m>1250:hvc,rg}\nqkd{x<3277:A,R}\nscm{a>3604:A,x>2828:gtl,s<3136:nn,R}\nlxd{x>1459:R,s<2218:R,m<1488:A,R}\nqvl{a>2241:zsv,a>2125:R,R}\nttb{x<1466:bfr,R}\nzbk{m<464:A,s<3379:R,R}\nxv{s>494:A,R}\npld{m<2671:R,a>3092:R,A}\nzjm{a<3448:A,A}\npd{s<2953:rm,szc}\ntfn{s<719:dh,m>343:bgz,bn}\nktl{s<2776:dp,s>3093:mfv,a>2897:fm,pd}\npnh{a<239:A,x<3185:R,R}\nrrc{m>1995:A,s<1027:A,A}\nnzp{s<2226:A,a>2676:A,A}\ntzf{a<3801:A,m>841:cgp,R}\nzp{m>158:A,A}\npr{x>2833:txv,x<2596:jf,a<3047:fj,np}\nrjl{a<2122:R,m<1656:R,m>1865:R,A}\nzx{x<1497:A,A}\nvtf{m<1214:R,R}\ntzk{x<3013:R,s<3078:ss,s>3545:tb,tpm}\nffj{x>1221:R,A}\ntkx{a<2272:R,x>915:R,A}\ndg{a<3554:A,m<929:vsm,hqx}\nlf{s<1049:R,R}\nkpz{a<2829:dsg,m<3108:R,m<3701:bp,kk}\nqf{s<635:jd,a>1648:qmt,fd}\ngsn{s<3722:A,s>3839:km,vqb}\ntxv{m<3809:A,R}\njrm{a<292:R,m<840:R,x<3311:A,R}\nzbg{a>538:A,m>1824:pnh,jrm}\nlmk{s<809:A,a<3296:R,A}\nrv{s<2680:A,s>2808:sr,x<1735:hhm,R}\nxc{m<2865:xn,bd}\ndzj{s>1477:R,x<2248:A,a<3592:R,flx}\nzpc{a>608:pv,x<934:sx,m<1037:cl,dj}\npvz{a>2770:A,A}\ngr{a>462:R,m<2929:A,A}\nzq{x<1166:tfn,m<571:fjm,s<1013:jss,qx}\nkp{s>355:md,m<2930:cc,nj}\njz{x<987:gzx,rl}\npp{x<3704:A,s>3206:R,m<3115:A,A}\nbcn{s>2021:A,m<1423:A,R}\nqtk{x<1430:R,x<2500:A,A}\nvsm{x<2477:R,R}\nvmb{m<1207:tns,jrn}\nlvq{s<810:A,s<1122:A,s<1245:A,A}\nkhs{x>2121:xvq,A}\nngj{s<399:R,A}\nss{s>2610:A,x>3640:R,s>2396:R,A}\nkbz{x<1704:A,a>3333:R,a<3253:A,A}\nmks{x<3622:A,x<3764:R,a>3793:A,A}\nmx{a>1771:A,R}\nnf{a<1730:R,x<3361:R,s>1633:A,A}\nsx{s<1067:R,x<583:bm,x<742:mcm,A}\ngjk{a>2935:A,x>3630:A,s>1501:R,R}\ngk{x>1901:pcd,s>1977:fpf,lqr}\nhz{s>2381:R,s<2154:A,R}\nnht{s>1162:A,m>558:R,x>2991:A,R}\npll{m>1200:sgn,x>395:phg,a>2429:xd,A}\nbj{x<1510:ks,s<729:jlm,gfx}\nxr{m>488:R,A}\ntfh{m>864:R,m<433:R,m>662:A,A}\nhfd{x<3738:R,R}\ncjt{a<2997:A,s<1260:tvq,m>1403:xm,R}\nkb{s<1000:A,s>1542:A,A}\nlsf{x>1237:fqk,A}\nmmv{x>3617:R,s<1009:mq,ns}\nbfr{x>951:A,s>1265:A,R}\npb{a<1607:zgf,kg}\njts{a>2338:vsj,x>1160:A,a>2097:zp,A}\nzmz{x<1333:R,a>2919:R,a<2898:R,A}\nql{s<2765:A,s>3539:R,a>2246:R,R}\nfqz{m>313:dl,R}\nbdt{m<2707:pld,lrr}\nvrj{a>3214:scm,zv}\nhkq{s>2374:sz,A}\nlp{s<1125:thr,sfr}\nkbh{s>3110:jcg,s>2776:R,m<1890:A,R}\ntk{x>916:A,x<350:A,R}\nxkq{a>3056:R,psn}\nsgz{a>579:A,x>2261:pc,m>3185:kzq,A}\nvd{x<3081:fsx,s>2036:dth,gkg}\nsr{s>2908:A,x>2436:A,x>1169:R,A}\nzm{m<3770:gzf,m>3888:R,nvh}\nkcs{x<3636:pkt,a<2781:hlg,x>3835:xsj,xr}\ntg{s>2776:R,R}\ndp{a>2961:R,A}\ngh{a<2396:R,m>972:R,x>3211:gpb,xp}\nsk{m<1280:A,A}\nbc{x<2841:A,a>1557:A,s<385:R,R}\nthf{a>3106:A,x>3631:A,s>1037:A,A}\nnn{m<2491:A,R}\nvhc{a<3571:nmk,rjb}\npcz{s<275:R,a>2784:R,A}\njgv{x<1067:vhf,lp}\nvvf{a>2403:R,R}\njsh{a>3401:A,A}\nvbp{x<2776:R,s>3492:R,m<2655:R,A}\nctl{m>1709:R,x>1436:R,a>3348:R,R}\nxrb{a<2649:R,x<934:A,A}\ngs{m>2043:A,x<3011:R,x>3230:R,R}\nrg{s<670:A,s<957:R,A}\nbl{s<1524:vx,s>1654:A,s<1601:A,nf}\ngfx{a<3594:A,s<974:R,A}\nkxl{s<1519:R,s>1562:R,A}\nds{a<3168:A,A}\nfcp{s<2341:R,m>297:R,A}\nglb{s<2066:R,a<3736:A,a<3870:R,R}\nfvx{x>263:A,A}\nmt{s>625:R,a<1776:A,s<217:R,A}\nqzv{m>1080:R,A}\nthr{a<3487:A,s>505:R,a>3759:A,R}\ntsb{a<3152:A,m<2461:A,a<3391:R,R}\nfsm{m<561:R,a<2786:A,R}\njzs{a<3752:A,x<1950:xf,m<1782:R,fs}\nkxk{m<570:A,R}\nzv{a>2688:R,s>3118:vbp,R}\nxcp{m>644:R,s<2394:A,m<338:A,R}\nlt{a>3140:ds,thf}\ncxt{a<3625:scz,tzf}\nqp{a>2830:zz,m>1424:zbp,sk}\nvb{s<3582:A,s>3789:R,A}\nrh{x>799:R,s>3176:A,R}\nhzq{x<2680:R,A}\njmm{s<1119:bj,m<680:pq,x>1590:dkd,bgc}\nvtg{s>3299:gj,s>2431:zx,A}\ndfq{s>1634:lkm,a>3475:mks,kxl}\ndcr{s>1739:A,A}\nbnn{a<207:R,x>446:A,m>3421:xdp,gr}\nhqx{a>3731:A,R}\nxd{s<3275:A,a<2588:R,m>437:R,R}\ndth{m<702:fqz,xtr}\nvnj{a>3383:A,m<1784:A,s>1539:R,R}\ndl{a>2976:A,s<3191:R,A}\nfv{x<3335:vrj,qkm}\npg{x<1444:R,m>2004:R,A}\nck{m>1984:kcq,m>1942:qtk,R}\nzz{m<1305:R,A}\npbc{m>226:A,s<1134:R,A}\nkx{m>1355:vhc,gn}\ngrr{m>2230:lzz,x<3328:A,s<1205:hc,cr}\nnj{x>1430:R,s>185:A,A}\nltb{a>597:tj,x>970:cm,bnn}\nqnn{x>3342:R,x>2644:R,R}\nmfp{s>1123:R,m>1993:ft,R}\nzgf{s>2601:A,m<973:A,tzx}\nnvh{a>2945:A,R}\nqxg{x>2968:xcp,A}\nbn{m<149:R,x>731:R,A}\ndqk{s>1560:R,s>534:A,A}\nmg{s>2297:rv,pgf}\nqt{s>1334:A,s<1018:R,R}\nhc{s>459:R,m<812:R,a>1440:R,A}\nfld{x>3004:R,A}\njd{x>2634:bc,s>234:R,a>1560:A,A}\nndr{m<1486:R,s>1539:dcr,R}\njvl{a<3475:A,s>2038:A,A}\ndkd{x<3070:dg,s>1405:dfq,lfp}\nsg{m>1909:A,x<914:R,m<1778:R,R}\nkhh{x<1396:R,a>3708:A,s>383:R,R}\nmb{s>1878:ftc,kpz}\nin{a<1965:fpj,m>2060:vh,a<3206:pdl,kx}\ngn{s>1904:qkj,jmm}\ngnm{x<3102:A,m<1222:kjn,gp}\nps{m>1681:hg,s>2179:qp,m<1433:fss,sd}\nzn{a<975:A,a>1069:A,x<3304:R,R}\nhhm{m>1630:A,m>617:A,m>333:A,A}\ngz{x<1706:A,s<2485:R,R}\nzpx{m>725:zss,xkq}\nfsx{s>2295:skb,bkr}\nscz{s<3285:dv,A}\njsj{x<3085:A,x<3686:R,A}\nxsj{a<2827:R,x>3904:A,A}\ngzq{m>810:A,a>827:R,s<3297:R,A}\nbb{x<2959:A,dz}\nvbl{m>2965:R,A}\ndmt{m>1732:hd,s>2127:R,jvl}\nlrn{x>1164:R,R}\nsmp{x>3224:R,s>1268:A,m<2046:A,A}\npc{s>693:A,A}\nkpx{m>1878:A,m>1807:R,x>928:R,R}\nglj{a>1711:R,a>1631:A,A}\nqk{m>1185:R,xrb}\nbdp{a<2516:lsf,a<2625:xk,a<2678:qk,tfh}\nczn{a<2241:fxp,x>3789:R,s<3133:rsm,A}\nks{x<999:xv,x>1311:khh,R}\nbkr{x<2429:mc,m<484:A,m<850:lft,R}\nbrl{s>1349:xg,gh}\nbx{x>1197:A,a>1427:R,s<3095:A,A}\nfhz{a>2435:A,m>3386:A,A}\nrl{a<3051:nsq,R}\ntr{s>1784:A,s>1693:A,m>190:A,A}\nmkj{x<1486:A,m>3363:A,a>1667:R,A}\nnb{a>1464:R,a>1365:jth,R}\nkzq{m>3512:R,A}\nkz{x>553:R,m>1531:A,m>1423:A,A}\ngl{a<2776:kgt,R}\nxpt{m<249:R,m>335:A,a>3718:A,R}\nzbp{a>2789:A,m<1574:A,s>3020:R,A}\nmfv{a>3208:A,x<1280:fhz,rf}\nqpn{x>3117:A,s<1302:A,a>2683:fjl,xnt}\nqm{a<2927:ps,tn}\npt{a<3892:R,a<3898:R,R}\nhks{x>2521:A,s<483:R,A}\nmd{x<1323:R,R}\njct{x>1059:gpl,a<2740:qv,mj}\nbt{m<1666:R,s>1464:A,m>1696:R,R}\ndsg{x<3064:A,A}\nnsk{x>3173:R,spq}\nqr{m>521:R,s>377:R,R}\nqqq{x<378:fg,m>2497:R,A}\ntj{a<935:jj,m>3383:A,a>1076:R,R}\nglg{a<1551:A,m>3175:R,R}\ncc{x<939:R,s<224:A,R}\nrk{s<1079:A,x<2703:A,A}\ndc{x<3243:R,a>568:R,jc}\nbcp{x>686:hqb,pll}\npft{m<912:vnn,dm}\nkfr{s>3103:nsk,x>3254:zm,pr}\nxvq{a>3436:A,m<1515:A,m<1586:R,R}\nbd{s<3314:ktl,jz}\nkm{s<3941:A,m<418:R,A}\nsj{s>3320:R,m<735:A,R}\njrn{m>1587:R,A}\ntx{a<2594:R,m<987:R,R}\ndj{m>1535:rrc,x<1518:xh,R}\nzgl{s>843:R,a>3152:R,m>1357:A,A}\nmfq{m>378:R,R}\nszc{s<3003:A,m<3402:R,m>3709:A,R}\nvnn{x<979:fdn,a<2498:A,a>2632:A,A}\nxkc{m>3553:kfr,pk}\nrzk{m>386:R,A}\ntbl{a<3119:A,m<1094:A,a>3172:A,A}\nsgn{s>3270:R,s<3159:R,x>306:R,A}\nkdh{x>1392:A,s<2174:gsg,m>2235:R,nzp}\nfg{x>157:A,x>87:R,A}\ndn{x>1120:R,A}\nzss{a<3072:A,s>3213:A,vtf}\ngsg{m>2239:A,x<588:R,a>2666:R,A}\njfz{m<1831:R,x>3533:jh,s>1453:gs,smp}\nphg{s>3312:R,m<723:R,m>941:R,A}\ncb{a>2536:A,s>812:pm,ngj}\nmf{m<1593:A,x<529:R,R}\nfhr{x>1106:R,x<386:A,x<798:kz,pcz}\nml{s<1370:R,R}\nzd{s>1589:A,a>1448:R,a>1311:A,lrh}\nxm{s<1627:R,a<3035:A,s<1799:R,A}\nsgs{x<3304:mt,x<3393:A,glj}\nttv{x>3479:nb,a<1580:grr,s<1026:sgs,bl}\nxp{x>3128:R,a>2555:R,R}\njx{s<2054:A,A}\nbjd{s<818:A,R}\ntzx{a<1332:R,m<1279:A,a<1434:R,A}\nmj{s<2372:A,s>2416:R,A}\ndh{s<467:A,m<418:A,m>634:A,R}\nfj{x<2703:A,x<2772:R,m<3838:R,A}\njxg{x>909:A,s>2752:A,s<2236:R,A}\nbfs{m<1689:R,m>1872:A,A}\ndz{m>3210:A,x>3525:R,A}\nmcm{a>283:R,a>116:R,a>49:A,A}\ndqr{m<1356:ckk,s>1025:nmm,a<3125:ffj,A}\ntbr{m<2367:R,x<2937:R,A}\nrsm{x<3615:R,x>3713:R,m>909:A,R}\nrm{m<3439:R,A}\nfmb{s>2014:nr,a<2950:hm,mpm}\npq{m<400:knp,dzj}\nlsk{m<322:R,s>2264:R,R}\nzsv{m>713:R,R}\nmk{a<2676:R,R}\nknp{s>1623:tr,s<1355:xpt,ssk}\nbmv{a<3815:qnx,m<1902:vtg,a<3932:qs,dmk}\nhlg{a>2740:R,m>464:A,R}\npn{s>759:A,m>1519:R,A}\nqnx{x<1888:R,A}\nnmm{x<1288:R,s<1542:A,A}\nff{a<305:A,R}\nrt{m<1666:khs,x>1599:kbh,bkk}\nbgz{x<430:A,x<824:A,A}\nhm{a>2828:ksr,m>1357:mhl,rnt}\nfdn{x<359:R,A}\nhq{a<2800:fth,x>1082:A,s<2977:R,zbk}\nxh{a<267:R,m>1348:R,a<492:R,A}\nrf{a<2730:R,a<3039:R,R}\npl{m>2474:mkj,a<1456:A,nh}\nhrg{m<3582:A,x<1494:R,x>2002:R,A}\nkn{x<2826:sgz,a>638:jb,x<3390:ccz,mmv}\ncgp{m<1166:R,a>3899:R,x<1426:R,R}\nnvp{s>1649:A,s<765:mzn,zkb}\nskb{s>3129:R,s>2585:R,x>2580:zb,R}\nfs{s<971:A,s<1381:R,s<1589:A,A}\nfm{s<2965:A,s>3026:zjm,A}\nzvz{s>2991:bv,mg}\nfxp{m<759:A,R}\nmzn{s<436:R,x<2333:A,a>2292:A,A}\nhg{x>2695:rn,s>2386:R,x<2329:R,A}\nxf{a<3868:A,A}\ngx{a<2954:nnv,m>2815:fbq,jgv}\nnp{a>3593:A,m>3821:R,R}\nnct{m<1503:A,glb}\nmpl{s<1048:R,gjk}\ngpl{x>1584:R,s<2360:R,a<2875:R,A}\ndmk{a<3973:czg,s<3272:pg,A}\nlqr{a<2295:mnt,s>823:bdp,s>373:vmb,pft}\nsdh{s<743:R,a<3416:R,m<3797:fhq,tp}\nxg{m<1238:R,x>3346:A,m<1761:A,R}\nnss{s<3633:rh,A}\njth{x>3707:R,A}\nbv{x<2477:ljl,s>3516:zbg,gsq}\nct{a>3210:R,A}\nmh{m<3086:A,A}\ngnt{m>1076:R,R}\nmhl{m>1815:rcm,s>751:pvz,s>442:gl,fhr}\nfz{m>447:A,s>2628:A,m>293:A,R}\ngp{a>573:R,a<367:A,m<1733:A,R}\nqq{s>1754:czn,a<2335:qn,cb}\nksr{s<1161:vp,s<1671:kr,a<2871:rjj,lj}\njhs{a>3568:A,a>3361:R,R}\npgf{x<2090:R,m<1854:rqz,m<3102:R,qnn}\ndrk{m>2799:R,x>844:A,m<2423:R,A}\nlr{a<3110:A,x>3689:lxp,tg}\ntsf{x<906:A,s<3626:R,A}\nslq{x>1074:kzc,x<635:qqq,a<2692:nss,kj}\nxtr{m>851:R,m>780:sn,sj}\nsd{x>3112:lf,m<1593:pn,a>2806:mtv,A}\nmqs{s<2800:A,a<3429:A,a>3486:A,R}\ndm{a<2520:R,s>247:ms,x>887:fpz,A}\ntdk{s<2421:R,a<3901:R,A}\nqx{s>1541:R,nv}\nhd{a>3498:R,a>3424:A,R}\nfq{x<1299:jl,a<3628:A,A}\nqv{x<689:R,x>933:R,s<2373:R,A}\nnnv{a<2442:ttb,a>2628:kb,s<712:kp,zh}\nxdp{m<3623:R,m<3797:A,x>156:A,A}\nshr{m>595:bqs,m<393:R,qr}\npkt{s<1087:R,a<2786:R,x<3363:A,R}\nxz{a>1176:kdx,m>2309:fl,lq}\ntpm{m>3107:A,R}\nsh{s>2775:R,A}\npk{a>2798:bb,m>3345:cj,tzk}\ngkg{a>3040:lt,a<2878:kcs,m<499:bjd,mpl}\ncjc{s<623:A,s<1014:A,a>3256:A,R}\nkc{s>697:A,R}\nxnt{a>2374:R,x>2862:R,A}\nmtv{s<1181:R,A}\nftc{s>1991:A,mh}\nqgj{a<2874:A,A}\npcs{a>3897:A,s>903:R,R}\nspq{a>2688:A,s>3621:R,m<3800:R,R}\ncfh{a<2335:xrc,fql}\nlj{s<1815:zmz,R}\ndvq{m>1770:A,s>678:kbz,ctl}\ntn{m<1611:kjv,a>3081:ml,zxd}\nrxz{s>802:sb,R}\nnmk{s>2328:rt,s<1387:tq,hr}\nkjn{a<773:A,a<910:A,s<582:R,A}\nfh{m<317:jts,m<594:gsn,x<661:qvl,bq}\ntdv{m<1383:A,rjl}\njhr{a<2523:A,a<2728:R,A}\nbqs{s<416:R,s>574:R,a>2772:R,A}\nqnv{x>702:R,s<1521:kxk,R}\nfhq{s<1199:R,a>3693:R,R}\ndrc{s<3614:R,m<1189:R,a>2578:R,R}\nbs{x<2494:gx,njd}\nfjl{s>1410:A,R}\nqn{a>2138:R,m>1309:lkc,x>3708:A,R}\njj{x>921:R,s>1153:R,R}\npdl{a<2720:gk,x<2022:fmb,tm}\ngsq{a<574:qkd,x<3457:dbp,m>1541:pp,gzq}\ntp{x>1394:R,x>699:A,a>3692:R,A}\nqmt{a>1858:rk,s>1346:mx,s>1021:R,A}\nlnx{m>1720:kpx,s>2143:lxd,R}\nmnt{s<1107:hb,m>1084:ndr,qnv}\njlm{a>3607:nqf,s<268:mtx,hks}\nbq{x>1150:stt,m>747:A,tkx}\ntt{m>724:gz,a<3383:lsk,s<2421:A,fz}\njcl{a>3378:A,s<1567:R,A}\npm{m<977:R,A}\nft{a<3968:R,A}\nktp{m<1837:jzs,dxl}\nvsj{m>207:R,A}\npjg{a<3083:A,s<3749:R,x<491:R,R}\nrn{x>3266:A,x<3028:A,s>2454:R,A}\nrjj{x<694:lck,R}\ncpc{a>3153:R,A}\ndf{x>2484:nct,zhf}\nfpz{s<129:R,m>1620:A,a<2635:A,A}\nqxb{a<1692:mxv,A}\nrjb{m<1740:df,s>1900:bmv,m<1912:ktp,th}\nfql{a<2531:kl,m<1336:R,m<1729:R,ddj}\nmtx{m<574:A,a>3385:R,x>2832:A,A}\nbjp{m<302:hzq,x<2687:dqk,vvf}\nkjv{a>3075:jsj,a<3004:R,a<3042:A,A}\ncr{a>1387:A,s>1535:R,A}\nfjm{x<1530:A,x>1775:pbc,x>1626:mfq,A}\njcg{x>2441:R,R}\nnsq{m>3370:R,m>3177:R,R}\nzt{a>1008:gd,zvz}\nxs{x>3300:kc,s>754:cg,a<2947:jhr,pz}\nfpf{s<3090:cfh,s<3519:bcp,m<934:fh,kt}\ngtl{m>2494:A,s<2767:R,s>3241:A,A}\ncl{m<658:nzt,R}\nbm{s<1419:A,A}\nqkm{s<3147:lr,px}\njc{a>292:A,R}\nlrr{s<2464:R,s>2665:A,R}\nsfr{s>1594:A,s<1406:A,A}\njss{s<385:cpc,A}\ncks{a<3347:A,m>1497:A,m>1438:R,A}\nvqb{s<3794:A,A}\ntm{m<1090:vd,qm}\nxk{m<842:A,m<1321:sp,qd}\nzf{a<2394:fcp,a<2554:A,s<2158:A,R}\nlrh{s<1419:R,x<652:R,s<1477:A,A}\nkcq{a<3825:R,R}\nsc{a<2768:R,s>2918:R,s>2526:A,R}\nlpk{m>1742:R,R}\nlzz{x<3346:A,s>660:R,x<3417:R,R}\nfqk{s>1274:R,R}\nbgc{a<3564:jcl,kqc}\nvh{s<2107:bs,x<2417:xc,lm}\nlb{x<1423:A,m<2424:A,m<2603:A,R}\ndbp{x<2990:A,R}\njf{x<2529:A,m>3833:R,A}\nzk{s<2758:hkq,x<2198:tl,hk}\nkg{s>3042:R,s<2430:R,R}\nzh{a<2554:drk,m<3186:A,hrg}\nrj{s>2451:sh,s>2299:jct,kdh}\nmxv{m>2171:A,m>1286:R,s<1567:R,R}\nnv{x<1693:A,A}\nssk{s<1524:A,A}\nfgs{m>2627:bdt,ct}\nxrc{m<1236:hz,a<2134:jp,s>2439:tqd,R}\nkqc{x>611:cmg,s<1534:fvx,gnt}\npx{m>2369:A,a>3103:A,hfd}\npcd{x<3064:sf,x>3489:qq,brl}\nhvc{m<1363:R,m>1409:R,R}\nrzc{a>3309:cks,a>3269:lmk,a<3243:A,cjc}\nnjd{s<1006:vc,s<1480:mfl,mb}\nvc{s>404:xs,mk}\nmdp{a<2842:A,x<200:A,x<431:R,A}\nnr{a<2936:tc,zpx}\nstt{s<3742:A,a<2429:A,A}\nbp{s>1625:A,R}\njh{x>3726:R,x>3640:A,a>637:R,A}\nczg{m>1984:A,A}\ngpb{m<443:R,m>754:A,A}\nkj{x<843:R,x<942:tsf,A}\nrnt{s<858:shr,fsm}\npv{m<1137:qb,s<792:nsg,x<1386:A,lpk}\nljl{m<2075:vb,R}\nvkf{x>836:pl,s<1216:mf,x>426:zd,qxb}\nkdx{x<2015:vkf,x<3140:qf,ttv}\nhr{s<1980:zr,x>2467:smb,a>3368:dmt,lnx}\nsn{m>817:A,x<3461:A,A}\nvx{a>1758:A,s<1203:R,s<1355:A,R}\nkr{a<2895:A,s<1391:R,A}\nccz{s>920:R,R}\ntvq{s>734:A,x<1240:R,x<1314:A,A}\nqh{x<1010:qsc,x>1431:vvd,a>3080:dqr,cjt}\nmc{a<3014:A,R}\nmbh{x>3007:A,s<1172:A,s<1552:A,A}\nrqp{x>1114:A,m<946:sm,hf}\nsz{a<1451:A,A}\ngzx{s>3618:pjg,x<579:mdp,szs}\nbps{m>1506:R,a>2520:R,A}\nfl{x>1717:kn,ltb}\ntb{x>3663:R,x>3384:R,R}\nzhf{m<1536:bcn,bt}\nsf{m>776:nvp,x<2370:zf,bjp}\nlfp{s>1251:qzv,jhs}\ndxl{m>1885:A,s>771:A,A}\nqsc{m>1428:A,m<1232:tbl,x>399:R,A}\njjp{a<3653:A,fb}\nffg{x<3318:A,s>2102:A,m<1905:R,A}\nlrz{x>1539:A,R}\njp{s>2644:R,R}\nzr{s<1776:vnj,s>1888:R,rgv}\nlxp{m>2494:R,m>2241:A,a>3647:A,A}\nrqz{a>338:A,A}\nfd{x<2753:R,mbh}\nhp{a>3707:R,a>3524:A,m>3277:R,R}\nvl{m>1659:sc,A}\nhnz{m>1381:bfs,s<3698:drc,a>2479:A,A}\nnsg{a>843:A,a>744:A,s<406:A,A}\nlm{m<2985:fv,xkc}\ndff{a>3044:rzk,x<1040:cpv,rxz}\nmfl{a>3319:hp,qpn}\nzkb{x<2485:R,m>1493:A,R}\njl{a>3338:A,m<3078:A,A}\ncg{s<841:A,m<3221:A,s>923:R,A}\nmq{a>414:A,s<617:R,a>253:R,A}\nhf{a>2898:A,R}\nxj{x>3177:A,s<1294:nht,m<651:R,R}\nzxd{a<2981:A,m>1800:A,x>2930:jx,A}\nsmb{m<1780:R,ffg}\nqkj{s>2787:cxt,gb}\nsp{a<2561:R,m>1156:A,A}\nlck{s<1850:R,A}\n\n{x=302,m=140,a=650,s=1288}\n{x=93,m=1448,a=1836,s=292}\n{x=3311,m=758,a=62,s=1463}\n{x=576,m=1505,a=314,s=1038}\n{x=7,m=685,a=4,s=1132}\n{x=472,m=57,a=3605,s=9}\n{x=2369,m=83,a=412,s=382}\n{x=2834,m=573,a=875,s=1127}\n{x=97,m=539,a=90,s=1568}\n{x=100,m=836,a=263,s=179}\n{x=846,m=227,a=396,s=261}\n{x=492,m=1517,a=150,s=935}\n{x=291,m=3188,a=1419,s=942}\n{x=1384,m=1726,a=74,s=1873}\n{x=690,m=369,a=115,s=3123}\n{x=1399,m=589,a=58,s=135}\n{x=512,m=85,a=1572,s=313}\n{x=253,m=2433,a=480,s=2736}\n{x=2148,m=2367,a=1531,s=96}\n{x=1008,m=1373,a=685,s=250}\n{x=706,m=815,a=57,s=3476}\n{x=236,m=1404,a=652,s=587}\n{x=1093,m=2589,a=229,s=329}\n{x=656,m=892,a=1729,s=438}\n{x=2046,m=2759,a=1544,s=568}\n{x=985,m=229,a=2228,s=56}\n{x=248,m=1893,a=112,s=609}\n{x=2837,m=73,a=1216,s=338}\n{x=290,m=44,a=9,s=153}\n{x=1499,m=1634,a=2499,s=1240}\n{x=2266,m=227,a=1560,s=1034}\n{x=32,m=3306,a=744,s=403}\n{x=1187,m=171,a=17,s=33}\n{x=330,m=259,a=710,s=655}\n{x=294,m=325,a=383,s=2483}\n{x=2434,m=168,a=1967,s=785}\n{x=11,m=586,a=916,s=743}\n{x=1773,m=362,a=1553,s=70}\n{x=1130,m=1599,a=31,s=622}\n{x=1595,m=23,a=1268,s=1578}\n{x=3912,m=283,a=1407,s=842}\n{x=1081,m=2942,a=807,s=375}\n{x=1257,m=1551,a=1523,s=12}\n{x=69,m=3109,a=1305,s=2094}\n{x=2625,m=1929,a=16,s=759}\n{x=926,m=1384,a=629,s=291}\n{x=917,m=373,a=740,s=2138}\n{x=59,m=1757,a=1283,s=323}\n{x=2434,m=510,a=9,s=56}\n{x=210,m=395,a=54,s=1646}\n{x=3555,m=2225,a=1096,s=120}\n{x=462,m=932,a=316,s=13}\n{x=177,m=2607,a=570,s=1945}\n{x=10,m=3261,a=2065,s=2260}\n{x=1017,m=250,a=1022,s=1572}\n{x=682,m=169,a=782,s=3}\n{x=718,m=483,a=1875,s=1986}\n{x=208,m=1563,a=1441,s=1752}\n{x=1509,m=1408,a=852,s=1057}\n{x=64,m=310,a=1589,s=897}\n{x=2427,m=757,a=2034,s=1266}\n{x=526,m=45,a=841,s=429}\n{x=320,m=778,a=797,s=1253}\n{x=65,m=2044,a=354,s=191}\n{x=2245,m=163,a=892,s=856}\n{x=850,m=321,a=42,s=522}\n{x=626,m=410,a=201,s=844}\n{x=1130,m=1745,a=1196,s=1984}\n{x=89,m=513,a=225,s=1482}\n{x=180,m=35,a=198,s=1999}\n{x=367,m=1679,a=609,s=127}\n{x=32,m=28,a=2507,s=736}\n{x=2219,m=1881,a=255,s=199}\n{x=635,m=2381,a=2243,s=2395}\n{x=1442,m=495,a=2242,s=1557}\n{x=863,m=63,a=3459,s=522}\n{x=1353,m=493,a=312,s=2}\n{x=445,m=401,a=587,s=114}\n{x=356,m=2289,a=1520,s=1686}\n{x=425,m=568,a=191,s=527}\n{x=2595,m=758,a=3569,s=708}\n{x=1311,m=2144,a=945,s=771}\n{x=167,m=666,a=1899,s=454}\n{x=1623,m=281,a=1474,s=579}\n{x=102,m=56,a=1720,s=722}\n{x=356,m=16,a=602,s=334}\n{x=1149,m=212,a=791,s=245}\n{x=2357,m=622,a=527,s=220}\n{x=1965,m=5,a=1133,s=1332}\n{x=1652,m=169,a=21,s=977}\n{x=372,m=100,a=674,s=1224}\n{x=1774,m=873,a=10,s=221}\n{x=8,m=604,a=15,s=201}\n{x=779,m=304,a=1486,s=52}\n{x=44,m=1797,a=532,s=59}\n{x=102,m=1013,a=1916,s=101}\n{x=1007,m=621,a=80,s=2917}\n{x=825,m=1444,a=845,s=1813}\n{x=396,m=55,a=1851,s=570}\n{x=2119,m=872,a=1209,s=2087}\n{x=1896,m=59,a=2709,s=1492}\n{x=680,m=216,a=45,s=1607}\n{x=2180,m=2415,a=149,s=793}\n{x=955,m=1397,a=616,s=441}\n{x=1337,m=746,a=97,s=33}\n{x=1291,m=703,a=1680,s=344}\n{x=1890,m=50,a=414,s=1292}\n{x=207,m=303,a=1959,s=514}\n{x=2020,m=1104,a=653,s=3283}\n{x=274,m=2059,a=835,s=604}\n{x=1945,m=22,a=86,s=711}\n{x=1225,m=390,a=1795,s=1384}\n{x=1837,m=238,a=638,s=295}\n{x=966,m=568,a=2067,s=755}\n{x=1942,m=340,a=1419,s=818}\n{x=53,m=2013,a=1221,s=1871}\n{x=86,m=2633,a=196,s=202}\n{x=102,m=2322,a=711,s=3427}\n{x=1967,m=477,a=67,s=1255}\n{x=2742,m=493,a=1673,s=411}\n{x=546,m=252,a=43,s=429}\n{x=821,m=422,a=2094,s=1533}\n{x=78,m=116,a=184,s=947}\n{x=502,m=258,a=1337,s=82}\n{x=186,m=372,a=313,s=271}\n{x=2717,m=586,a=502,s=342}\n{x=490,m=2086,a=45,s=1424}\n{x=546,m=1048,a=411,s=663}\n{x=1387,m=2721,a=534,s=1500}\n{x=22,m=1,a=1231,s=2927}\n{x=584,m=1382,a=752,s=1516}\n{x=2467,m=1090,a=89,s=2937}\n{x=2783,m=126,a=1324,s=801}\n{x=1343,m=87,a=2232,s=2066}\n{x=846,m=980,a=2807,s=3698}\n{x=1260,m=20,a=128,s=199}\n{x=26,m=1734,a=53,s=1459}\n{x=713,m=1371,a=762,s=827}\n{x=1563,m=946,a=539,s=643}\n{x=1898,m=3022,a=3464,s=66}\n{x=508,m=113,a=2159,s=186}\n{x=818,m=1473,a=232,s=11}\n{x=140,m=1337,a=2485,s=577}\n{x=1196,m=1725,a=857,s=989}\n{x=374,m=291,a=276,s=120}\n{x=162,m=665,a=660,s=84}\n{x=1829,m=590,a=1732,s=551}\n{x=55,m=421,a=178,s=349}\n{x=51,m=2387,a=235,s=1420}\n{x=1795,m=3341,a=624,s=692}\n{x=923,m=641,a=3147,s=1686}\n{x=62,m=2162,a=953,s=520}\n{x=133,m=218,a=3130,s=346}\n{x=436,m=291,a=1366,s=1710}\n{x=341,m=2054,a=420,s=593}\n{x=1451,m=85,a=1296,s=355}\n{x=1128,m=372,a=411,s=264}\n{x=960,m=477,a=1139,s=444}\n{x=2073,m=1841,a=1447,s=734}\n{x=1808,m=3556,a=2231,s=80}\n{x=1004,m=2097,a=417,s=925}\n{x=1593,m=1820,a=1450,s=3391}\n{x=576,m=932,a=90,s=742}\n{x=295,m=2888,a=2296,s=338}\n{x=2693,m=190,a=630,s=1186}\n{x=1086,m=66,a=85,s=893}\n{x=1181,m=32,a=814,s=88}\n{x=1671,m=650,a=933,s=1900}\n{x=110,m=682,a=852,s=1431}\n{x=141,m=89,a=563,s=2260}\n{x=771,m=2014,a=216,s=753}\n{x=1485,m=3065,a=2418,s=26}\n{x=2175,m=828,a=1241,s=2374}\n{x=601,m=785,a=2903,s=707}\n{x=2635,m=865,a=1163,s=971}\n{x=560,m=2287,a=1469,s=2081}\n{x=246,m=4,a=26,s=2596}\n{x=1104,m=109,a=788,s=165}\n{x=1142,m=2257,a=1125,s=155}\n{x=3292,m=2635,a=511,s=744}\n{x=2205,m=395,a=197,s=553}\n{x=1168,m=1348,a=510,s=238}\n{x=1827,m=289,a=359,s=293}\n{x=513,m=352,a=1181,s=330}\n{x=164,m=2153,a=263,s=408}\n{x=363,m=158,a=275,s=2266}\n{x=544,m=404,a=2695,s=692}\n{x=7,m=730,a=2276,s=32}\n{x=988,m=340,a=1480,s=191}\n{x=690,m=85,a=1289,s=57}\n{x=368,m=813,a=762,s=1787}\n{x=2088,m=2053,a=1917,s=2756}\n{x=3331,m=712,a=1735,s=309}\n{x=334,m=490,a=268,s=1451}\n{x=631,m=70,a=371,s=577}\n{x=662,m=893,a=541,s=2204}\n{x=1147,m=354,a=1121,s=1455}\n{x=2517,m=2025,a=2242,s=674}\n{x=418,m=1394,a=1326,s=405}\n{x=1630,m=354,a=226,s=1078}",
	}

	var res []int
	for _, d := range data {
		input := strings.Split(d, "\n\n")
		workflows := strings.Split(input[0], "\n")
		ratings := strings.Split(input[1], "\n")
		res = append(res, d19p1(workflows, ratings))
		res = append(res, d19p2(workflows))
	}

	return res
}

type d19rule struct {
	isAlways bool
	varName  string
	ge       bool
	value    int
	workflow string
}

func d19p2(workflows []string) int {
	wf := d19parseWf(workflows)
	type p2Process struct {
		wfName     string
		x, m, a, s [2]int
	}
	processes := []*p2Process{{wfName: "in", x: [2]int{0, 4001}, m: [2]int{0, 4001}, a: [2]int{0, 4001}, s: [2]int{0, 4001}}}
	res := 0
	for len(processes) > 0 {
		p := processes[0]
		processes = processes[1:]
		if p.wfName == "A" {
			res += (p.x[1] - p.x[0] - 1) * (p.m[1] - p.m[0] - 1) * (p.a[1] - p.a[0] - 1) * (p.s[1] - p.s[0] - 1)
			continue
		}
		w := wf[p.wfName]
		for _, rule := range w {
			if rule.isAlways {
				processes = append(processes, &p2Process{wfName: rule.workflow, x: p.x, m: p.m, a: p.a, s: p.s})
				continue
			}
			x, px, ok1 := d19processRule(rule, p.x, "x", rule.varName)
			m, pm, ok2 := d19processRule(rule, p.m, "m", rule.varName)
			a, pa, ok3 := d19processRule(rule, p.a, "a", rule.varName)
			s, ps, ok4 := d19processRule(rule, p.s, "s", rule.varName)
			p.x, p.m, p.a, p.s = px, pm, pa, ps

			if !ok1 || !ok2 || !ok3 || !ok4 {
				continue
			}
			processes = append(processes, &p2Process{wfName: rule.workflow, x: x, m: m, a: a, s: s})
		}
	}

	return res
}

func d19processRule(rule d19rule, value [2]int, v1, v2 string) ([2]int, [2]int, bool) {
	var res = [2]int{value[0], value[1]}
	if v1 != v2 {
		return res, value, true
	}
	if rule.ge {
		res[0] = max(value[0], rule.value)
		value[1] = min(value[1], rule.value+1)
	} else {
		res[1] = min(value[1], rule.value)
		value[0] = max(value[0], rule.value-1)
	}
	return res, value, res[0] < res[1]
}

func d19p1(workflows, ratings []string) int {
	wf := d19parseWf(workflows)
	res := 0
	for _, rt := range ratings {
		var x, m, a, s int
		fmt.Sscanf(rt, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		var r = map[string]int{"x": x, "m": m, "a": a, "s": s}
		if d19p1ia(wf, r) {
			res += x + m + a + s
		}
	}

	return res
}

func d19p1ia(wf map[string][]d19rule, r map[string]int) bool {
	wfName := "in"
	for wfName != "A" && wfName != "R" {
		for _, rule := range wf[wfName] {
			if rule.isAlways {
				wfName = rule.workflow
				break
			}
			val := r[rule.varName]
			if val > rule.value == rule.ge {
				wfName = rule.workflow
				break
			}
		}
	}

	return wfName == "A"
}

func d19parseWf(workflows []string) map[string][]d19rule {
	wf := make(map[string][]d19rule)
	for _, w := range workflows {
		w_ := strings.Split(w, "{")
		name := w_[0]
		rulesRaw := strings.Split(w_[1][:len(w_[1])-1], ",")
		rules := make([]d19rule, 0)
		for _, rr := range rulesRaw {
			r := strings.Split(rr, ":")
			if len(r) == 1 {
				rules = append(rules, d19rule{isAlways: true, workflow: r[0]})
				continue
			}
			value, _ := strconv.Atoi(r[0][2:])
			rl := d19rule{workflow: r[1], varName: r[0][:1], ge: r[0][1] == '>', value: value}
			rules = append(rules, rl)
		}
		wf[name] = rules
	}
	return wf
}
