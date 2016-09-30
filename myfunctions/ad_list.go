package myfunctions


func InitVariables() (string, string)  {
	
	help := `
	Usage:
	
	To fetch metar and tafs, use either ICAO or IATA codes (case insensitive)

		$ metar bru ebci tls tlv LFPG edDH
	
	The output gives the latest 4 metars and the latest taf. For every metar,
	the Wind Chill factor, Heat Factor and Relative Humidity are calculated.
	
	To search any airport based on its name instead of IATA/ICAO code, just free search using:
	
		$ metar -s Ushua
	
	For help (this screen)
	
		$ metar -h
	
	Data retrieved from aviationweather.gov
	
	Coded in Go (Golang) from Google's new compiled system language created by 
	Robert Pike, Ken Thomson (himself!) and  Robert Griesemer

	Go rocks!

`
	
	adlist :=`GKA;AYGA;Goroka;Papua New Guinea
MAG;AYMD;Madang;Papua New Guinea
HGU;AYMH;Mount Hagen;Papua New Guinea
LAE;AYNZ;Nadzab;Papua New Guinea
POM;AYPY;Port Moresby (Port Moresby Jacksons Intl);Papua New Guinea
WWK;AYWK;Wewak (Wewak Intl);Papua New Guinea
UAK;BGBW;Narssarssuaq (Narsarsuaq);Greenland
GOH;BGGH;Godthaab (Nuuk);Greenland
SFJ;BGSF;Sondrestrom (Sondre Stromfjord);Greenland
THU;BGTL;Thule (Thule Air Base);Greenland
AEY;BIAR;Akureyri;Iceland
EGS;BIEG;Egilsstadir;Iceland
HFN;BIHN;Hofn (Hornafjordur);Iceland
HZK;BIHU;Husavik;Iceland
IFJ;BIIS;Isafjordur;Iceland
KEF;BIKF;Keflavik (Keflavik International Airport);Iceland
PFJ;BIPA;Patreksfjordur;Iceland
RKV;BIRK;Reykjavik;Iceland
SIJ;BISI;Siglufjordur;Iceland
VEY;BIVM;Vestmannaeyjar;Iceland
YAM;CYAM;Sault Sainte Marie (Sault Ste Marie);Canada
YAV;CYAV;Winnipeg (Winnipeg St Andrews);Canada
YAW;CYAW;Halifax (Shearwater);Canada
YAY;CYAY;St. Anthony (St Anthony);Canada
YAZ;CYAZ;Tofino;Canada
YBB;CYBB;Pelly Bay (Kugaaruk);Canada
YBC;CYBC;Baie Comeau;Canada
YBG;CYBG;Bagotville;Canada
YBK;CYBK;Baker Lake;Canada
YBL;CYBL;Campbell River;Canada
YBR;CYBR;Brandon (Brandon Muni);Canada
YCB;CYCB;Cambridge Bay;Canada
YCD;CYCD;Nanaimo;Canada
YCG;CYCG;Castlegar;Canada
YCH;CYCH;Chatham (Miramichi);Canada
YCL;CYCL;Charlo;Canada
YCO;CYCO;Coppermine (Kugluktuk);Canada
YCT;CYCT;Coronation;Canada
YCW;CYCW;Chilliwack;Canada
YCY;CYCY;Clyde River;Canada
YZS;CYCZ;Coral Harbour (Fairmont Hot Springs);Canada
YDA;CYDA;Dawson (Dawson City);Canada
YDB;CYDB;Burwash;Canada
YDC;CYDC;Princeton;Canada
YDF;CYDF;Deer Lake;Canada
YDL;CYDL;Dease Lake;Canada
YDN;CYDN;Dauphin (Dauphin Barker);Canada
YDQ;CYDQ;Dawson Creek;Canada
YEG;CYEG;Edmonton (Edmonton Intl);Canada
YEK;CYEK;Eskimo Point (Arviat);Canada
YEN;CYEN;Estevan;Canada
YET;CYET;Edson;Canada
YEU;CYEU;Eureka;Canada
YEV;CYEV;Inuvik (Inuvik Mike Zubko);Canada
YFB;CYFB;Iqaluit;Canada
YFC;CYFC;Fredericton;Canada
YFO;CYFO;Flin Flon;Canada
YFR;CYFR;Fort Resolution;Canada
YFS;CYFS;Fort Simpson;Canada
YGK;CYGK;Kingston;Canada
YGL;CYGL;La Grande Riviere;Canada
YGP;CYGP;Gaspe;Canada
YGQ;CYGQ;Geraldton (Geraldton Greenstone Regional);Canada
YGR;CYGR;Iles De La Madeleine;Canada
YHB;CYHB;Hudson Bay;Canada
YHD;CYHD;Dryden (Dryden Rgnl);Canada
YHI;CYHI;Holman Island (Ulukhaktok Holman);Canada
YHK;CYHK;Gjoa Haven;Canada
YHM;CYHM;Hamilton;Canada
YHU;CYHU;Montreal (St Hubert);Canada
YHY;CYHY;Hay River;Canada
YHZ;CYHZ;Halifax (Halifax Intl);Canada
YIB;CYIB;Atikokan (Atikokan Muni);Canada
YIO;CYIO;Pond Inlet;Canada
YJN;CYJN;St. Jean (St Jean);Canada
YJT;CYJT;Stephenville;Canada
YKA;CYKA;Kamloops;Canada
YKF;CYKF;Waterloo;Canada
YKL;CYKL;Schefferville;Canada
YKY;CYKY;Kindersley;Canada
YKZ;CYKZ;Toronto (Buttonville Muni);Canada
YLD;CYLD;Chapleau;Canada
YLJ;CYLJ;Meadow Lake;Canada
YLL;CYLL;Lloydminster;Canada
YLT;CYLT;Alert;Canada
YLW;CYLW;Kelowna;Canada
YMA;CYMA;Mayo;Canada
YMJ;CYMJ;Moose Jaw (Moose Jaw Air Vice Marshal C M Mcewen);Canada
YMM;CYMM;Fort Mcmurray;Canada
YMO;CYMO;Moosonee;Canada
YMW;CYMW;Maniwaki;Canada
YMX;CYMX;Montreal (Montreal Intl Mirabel);Canada
YNA;CYNA;Natashquan;Canada
YND;CYND;Gatineau;Canada
YNM;CYNM;Matagami;Canada
YOC;CYOC;Old Crow;Canada
YOD;CYOD;Cold Lake;Canada
YOJ;CYOJ;High Level;Canada
YOW;CYOW;Ottawa (Ottawa Macdonald Cartier Intl);Canada
YPA;CYPA;Prince Albert (Prince Albert Glass Field);Canada
YPE;CYPE;Peace River;Canada
YPG;CYPG;Portage-la-prairie (Southport);Canada
YPL;CYPL;Pickle Lake;Canada
YPN;CYPN;Port Menier;Canada
YPQ;CYPQ;Peterborough;Canada
YPR;CYPR;Prince Pupert (Prince Rupert);Canada
YPY;CYPY;Fort Chipewyan;Canada
YQA;CYQA;Muskoka;Canada
YQB;CYQB;Quebec (Quebec Jean Lesage Intl);Canada
YQF;CYQF;Red Deer Industrial (Red Deer Regional);Canada
YQG;CYQG;Windsor;Canada
YQH;CYQH;Watson Lake;Canada
YQK;CYQK;Kenora;Canada
YQL;CYQL;Lethbridge;Canada
YQM;CYQM;Moncton (Greater Moncton Intl);Canada
YQQ;CYQQ;Comox;Canada
YQR;CYQR;Regina (Regina Intl);Canada
YQT;CYQT;Thunder Bay;Canada
YQU;CYQU;Grande Prairie;Canada
YQV;CYQV;Yorkton (Yorkton Muni);Canada
YQW;CYQW;North Battleford;Canada
YQX;CYQX;Gander (Gander Intl);Canada
YQY;CYQY;Sydney;Canada
YQZ;CYQZ;Quesnel;Canada
YRB;CYRB;Resolute (Resolute Bay);Canada
YRI;CYRI;Riviere Du Loup;Canada
YRJ;CYRJ;Roberval;Canada
YRM;CYRM;Rocky Mountain House;Canada
YRT;CYRT;Rankin Inlet;Canada
YSB;CYSB;Sudbury;Canada
YSC;CYSC;Sherbrooke;Canada
YSJ;CYSJ;St. John (Saint John);Canada
YSM;CYSM;Fort Smith;Canada
YSR;CYSR;Nanisivik;Canada
YSU;CYSU;Summerside;Canada
YSY;CYSY;Sachs Harbour;Canada
YTE;CYTE;Cape Dorset;Canada
YTH;CYTH;Thompson;Canada
YTR;CYTR;Trenton;Canada
YTS;CYTS;Timmins;Canada
YTZ;CYTZ;Toronto (City Centre);Canada
YUB;CYUB;Tuktoyaktuk;Canada
YUL;CYUL;Montreal (Pierre Elliott Trudeau Intl);Canada
YUT;CYUT;Repulse Bay;Canada
YUX;CYUX;Hall Beach;Canada
YUY;CYUY;Rouyn (Rouyn Noranda);Canada
YVC;CYVC;La Ronge;Canada
YVG;CYVG;Vermillion (Vermilion);Canada
YVM;CYVM;Broughton Island (Qikiqtarjuaq);Canada
YVO;CYVO;Val D'or (Val D Or);Canada
YVP;CYVP;Quujjuaq (Kuujjuaq);Canada
YVQ;CYVQ;Norman Wells;Canada
YVR;CYVR;Vancouver (Vancouver Intl);Canada
YVT;CYVT;Buffalo Narrows;Canada
YVV;CYVV;Wiarton;Canada
YWA;CYWA;Petawawa;Canada
YWG;CYWG;Winnipeg (Winnipeg Intl);Canada
YWK;CYWK;Wabush;Canada
YWL;CYWL;Williams Lake;Canada
YWY;CYWY;Wrigley;Canada
YXC;CYXC;Cranbrook (Canadian Rockies Intl);Canada
YXD;CYXD;Edmonton (Edmonton City Centre);Canada
YXE;CYXE;Saskatoon (Saskatoon J G Diefenbaker Intl);Canada
YXH;CYXH;Medicine Hat;Canada
YXJ;CYXJ;Fort Saint John (Fort St John);Canada
YXL;CYXL;Sioux Lookout;Canada
YXP;CYXP;Pangnirtung;Canada
YXR;CYXR;Earlton (Timiskaming Rgnl);Canada
YXS;CYXS;Prince George;Canada
YXT;CYXT;Terrace;Canada
YXU;CYXU;London;Canada
YXX;CYXX;Abbotsford;Canada
YXY;CYXY;Whitehorse (Whitehorse Intl);Canada
YYB;CYYB;North Bay;Canada
YYC;CYYC;Calgary (Calgary Intl);Canada
YYD;CYYD;Smithers;Canada
YYE;CYYE;Fort Nelson;Canada
YYF;CYYF;Penticton;Canada
YYG;CYYG;Charlottetown;Canada
YYH;CYYH;Spence Bay (Taloyoak);Canada
YYJ;CYYJ;Victoria (Victoria Intl);Canada
YYL;CYYL;Lynn Lake;Canada
YYN;CYYN;Swift Current;Canada
YYQ;CYYQ;Churchill;Canada
YYR;CYYR;Goose Bay;Canada
YYT;CYYT;St. John's (St Johns Intl);Canada
YYU;CYYU;Kapuskasing;Canada
YYW;CYYW;Armstrong;Canada
YYY;CYYY;Mont Joli;Canada
YYZ;CYYZ;Toronto (Lester B Pearson Intl);Canada
YZD;CYZD;Toronto (Downsview);Canada
YZE;CYZE;Gore Bay (Gore Bay Manitoulin);Canada
YZF;CYZF;Yellowknife;Canada
YZH;CYZH;Slave Lake;Canada
YZP;CYZP;Sandspit;Canada
YZR;CYZR;Sarnia (Chris Hadfield);Canada
YZT;CYZT;Port Hardy;Canada
YZU;CYZU;Whitecourt;Canada
YZV;CYZV;Sept-iles (Sept Iles);Canada
YZW;CYZW;Teslin;Canada
YZX;CYZX;Greenwood;Canada
ZFA;CZFA;Faro;Canada
ZFM;CZFM;Fort Mcpherson;Canada
BJA;DAAE;Bejaja (Soummam);Algeria
ALG;DAAG;Algier (Houari Boumediene);Algeria
DJG;DAAJ;Djanet (Tiska);Algeria
QFD;DAAK;Boufarik;Algeria
VVZ;DAAP;Illizi (Illizi Takhamalt);Algeria
TMR;DAAT;Tamanrasset;Algeria
GJL;DAAV;Jijel;Algeria
AAE;DABB;Annaba;Algeria
CZL;DABC;Constantine (Mohamed Boudiaf Intl);Algeria
TEE;DABS;Tebessa (Cheikh Larbi Tebessi);Algeria
HRM;DAFH;Tilrempt (Hassi R Mel);Algeria
TID;DAOB;Tiaret (Bou Chekif);Algeria
TIN;DAOF;Tindouf;Algeria
QAS;DAOI;Ech-cheliff (Ech Cheliff);Algeria
TAF;DAOL;Oran (Tafaraoui);Algeria
TLM;DAON;Tlemcen (Zenata);Algeria
ORN;DAOO;Oran (Es Senia);Algeria
MUW;DAOV;Ghriss;Algeria
AZR;DAUA;Adrar (Touat Cheikh Sidi Mohamed Belkebir);Algeria
BSK;DAUB;Biskra;Algeria
ELG;DAUE;El Golea;Algeria
GHA;DAUG;Ghardaia (Noumerat);Algeria
HME;DAUH;Hassi Messaoud (Oued Irara);Algeria
INZ;DAUI;In Salah;Algeria
TGR;DAUK;Touggourt (Sidi Mahdi);Algeria
LOO;DAUL;Laghouat;Algeria
TMX;DAUT;Timimoun;Algeria
OGX;DAUU;Ouargla;Algeria
IAM;DAUZ;Zarzaitine (In Amenas);Algeria
COO;DBBB;Cotonou (Cadjehoun);Benin
OUA;DFFD;Ouagadougou;Burkina Faso
BOY;DFOO;Bobo-dioulasso (Bobo Dioulasso);Burkina Faso
ACC;DGAA;Accra (Kotoka Intl);Ghana
TML;DGLE;Tamale;Ghana
NYI;DGSN;Sunyani;Ghana
TKD;DGTK;Takoradi;Ghana
ABJ;DIAP;Abidjan (Abidjan Felix Houphouet Boigny Intl);Cote d'Ivoire
BYK;DIBK;Bouake;Cote d'Ivoire
DJO;DIDL;Daloa;Cote d'Ivoire
HGO;DIKO;Korhogo;Cote d'Ivoire
MJC;DIMN;Man;Cote d'Ivoire
SPY;DISP;San Pedro;Cote d'Ivoire
ASK;DIYO;Yamoussoukro;Cote d'Ivoire
ABV;DNAA;Abuja (Nnamdi Azikiwe Intl);Nigeria
AKR;DNAK;Akure;Nigeria
BNI;DNBE;Benin;Nigeria
CBQ;DNCA;Calabar;Nigeria
ENU;DNEN;Enugu;Nigeria
QUS;DNGU;Gusau;Nigeria
IBA;DNIB;Ibadan;Nigeria
ILR;DNIL;Ilorin;Nigeria
JOS;DNJO;Jos (Yakubu Gowon);Nigeria
KAD;DNKA;Kaduna;Nigeria
KAN;DNKN;Kano (Mallam Aminu Intl);Nigeria
MIU;DNMA;Maiduguri;Nigeria
MDI;DNMK;Makurdi;Nigeria
LOS;DNMM;Lagos (Murtala Muhammed);Nigeria
MXJ;DNMN;Minna (Minna New);Nigeria
PHC;DNPO;Port Hartcourt (Port Harcourt Intl);Nigeria
SKO;DNSO;Sokoto (Sadiq Abubakar Iii Intl);Nigeria
YOL;DNYO;Yola;Nigeria
ZAR;DNZA;Zaria;Nigeria
MFQ;DRRM;Maradi;Niger
NIM;DRRN;Niamey (Diori Hamani);Niger
THZ;DRRT;Tahoua;Niger
AJY;DRZA;Agadez (Manu Dayak);Niger
ZND;DRZR;Zinder;Niger
MIR;DTMB;Monastir (Habib Bourguiba Intl);Tunisia
TUN;DTTA;Tunis (Carthage);Tunisia
GAF;DTTF;Gafsa;Tunisia
GAE;DTTG;Gabes;Tunisia
DJE;DTTJ;Djerba (Zarzis);Tunisia
EBM;DTTR;El Borma;Tunisia
SFA;DTTX;Sfax (Thyna);Tunisia
TOE;DTTZ;Tozeur (Nefta);Tunisia
LRL;DXNG;Niatougou (Niamtougou International);Togo
LFW;DXXX;Lome (Gnassingbe Eyadema Intl);Togo
ANR;EBAW;Antwerp (Deurne);Belgium
BRU;EBBR;Brussels (Brussels Natl);Belgium
CRL;EBCI;Charleroi (Brussels South);Belgium
QKT;EBKT;Kortrijk-vevelgem (Wevelgem);Belgium
LGG;EBLG;Liege;Belgium
OST;EBOS;Ostend (Oostende);Belgium
BBJ;EDAB;Bautzen;Germany
AOC;EDAC;Altenburg (Altenburg Nobitz);Germany
BBH;EDBH;Barth;Germany
SXF;EDDB;Berlin (Schonefeld);Germany
DRS;EDDC;Dresden;Germany
ERF;EDDE;Erfurt;Germany
FRA;EDDF;Frankfurt (Frankfurt Main);Germany
FMO;EDDG;Munster (Munster Osnabruck);Germany
HAM;EDDH;Hamburg;Germany
THF;EDDI;Berlin (Tempelhof);Germany
CGN;EDDK;Cologne (Koln Bonn);Germany
DUS;EDDL;Duesseldorf (Dusseldorf);Germany
MUC;EDDM;Munich (Franz Josef Strauss);Germany
NUE;EDDN;Nuernberg (Nurnberg);Germany
LEJ;EDDP;Leipzig (Leipzig Halle);Germany
SCN;EDDR;Saarbruecken (Saarbrucken);Germany
STR;EDDS;Stuttgart;Germany
TXL;EDDT;Berlin (Tegel);Germany
HAJ;EDDV;Hannover;Germany
BRE;EDDW;Bremen (Neuenland);Germany
HHN;EDFH;Hahn (Frankfurt Hahn);Germany
MHG;EDFM;Mannheim (Mannheim City);Germany
XFW;EDHI;Hamburg (Hamburg Finkenwerder);Germany
KEL;EDHK;Kiel (Kiel Holtenau);Germany
LBC;EDHL;Luebeck (Lubeck Blankensee);Germany
ZCA;EDLA;Arnsberg (Arnsberg Menden);Germany
ESS;EDLE;Essen (Essen Mulheim);Germany
MGL;EDLN;Moenchengladbach (Monchengladbach);Germany
PAD;EDLP;Paderborn (Paderborn Lippstadt);Germany
DTM;EDLW;Dortmund;Germany
AGB;EDMA;Augsburg;Germany
OBF;EDMO;Oberpfaffenhofen;Germany
FDH;EDNY;Friedrichshafen;Germany
SZW;EDOP;Parchim (Schwerin Parchim);Germany
ZSN;EDOV;Stendal (Stendal Borstel);Germany
BYU;EDQD;Bayreuth;Germany
HOQ;EDQM;Hof (Hof Plauen);Germany
ZNV;EDRK;Koblenz (Koblenz Winningen);Germany
ZQF;EDRT;Trier (Trier Fohren);Germany
ZQC;EDRY;Speyer;Germany
ZQL;EDTD;Donaueschingen (Donaueschingen Villingen);Germany
BWE;EDVE;Braunschweig (Braunschweig Wolfsburg);Germany
KSF;EDVK;Kassel (Kassel Calden);Germany
BRV;EDWB;Bremerhaven;Germany
EME;EDWE;Emden;Germany
WVN;EDWI;Wilhelmshaven (Wilhelmshaven Mariensiel);Germany
BMK;EDWR;Borkum;Germany
NRD;EDWY;Norderney;Germany
FLF;EDXF;Flensburg (Flensburg Schaferhaus);Germany
GWT;EDXW;Westerland (Westerland Sylt);Germany
KDL;EEKA;Kardla;Estonia
URE;EEKE;Kuressaare;Estonia
EPU;EEPU;Parnu;Estonia
TLL;EETN;Tallinn-ulemiste International (Tallinn);Estonia
TAY;EETU;Tartu;Estonia
ENF;EFET;Enontekio;Finland
KEV;EFHA;Halli;Finland
HEM;EFHF;Helsinki (Helsinki Malmi);Finland
HEL;EFHK;Helsinki (Helsinki Vantaa);Finland
HYV;EFHV;Hyvinkaa;Finland
IVL;EFIV;Ivalo;Finland
JOE;EFJO;Joensuu;Finland
JYV;EFJY;Jyvaskyla;Finland
KAU;EFKA;Kauhava;Finland
KEM;EFKE;Kemi (Kemi Tornio);Finland
KAJ;EFKI;Kajaani;Finland
KOK;EFKK;Kruunupyy;Finland
KAO;EFKS;Kuusamo;Finland
KTT;EFKT;Kittila;Finland
KUO;EFKU;Kuopio;Finland
LPP;EFLP;Lappeenranta;Finland
MHQ;EFMA;Mariehamn;Finland
MIK;EFMI;Mikkeli;Finland
OUL;EFOU;Oulu;Finland
POR;EFPO;Pori;Finland
RVN;EFRO;Rovaniemi;Finland
SVL;EFSA;Savonlinna;Finland
SOT;EFSO;Sodankyla;Finland
TMP;EFTP;Tampere (Tampere Pirkkala);Finland
TKU;EFTU;Turku;Finland
QVY;EFUT;Utti;Finland
VAA;EFVA;Vaasa;Finland
VRK;EFVR;Varkaus;Finland
BFS;EGAA;Belfast (Belfast Intl);United Kingdom
ENK;EGAB;Enniskillen (St Angelo);United Kingdom
BHD;EGAC;Belfast (Belfast City);United Kingdom
LDY;EGAE;Londonderry (City of Derry);United Kingdom
BHX;EGBB;Birmingham;United Kingdom
CVT;EGBE;Coventry;United Kingdom
GLO;EGBJ;Golouchestershire (Gloucestershire);United Kingdom
MAN;EGCC;Manchester;United Kingdom
NQY;EGDG;Newquai (St Mawgan);United Kingdom
LYE;EGDL;Lyneham;United Kingdom
YEO;EGDY;Yeovilton;United Kingdom
CWL;EGFF;Cardiff;United Kingdom
SWS;EGFH;Swansea;United Kingdom
BRS;EGGD;Bristol;United Kingdom
LPL;EGGP;Liverpool;United Kingdom
LTN;EGGW;London (Luton);United Kingdom
PLH;EGHD;Plymouth;United Kingdom
BOH;EGHH;Bournemouth;United Kingdom
SOU;EGHI;Southampton;United Kingdom
QLA;EGHL;Lasham;United Kingdom
ACI;EGJA;Alderney;Guernsey
GCI;EGJB;Guernsey;Guernsey
JER;EGJJ;Jersey;Jersey
ESH;EGKA;Shoreham By Sea (Shoreham);United Kingdom
BQH;EGKB;Biggin Hill;United Kingdom
LGW;EGKK;London (Gatwick);United Kingdom
LCY;EGLC;London (City);United Kingdom
FAB;EGLF;Farnborough;United Kingdom
BBS;EGLK;Blackbushe;United Kingdom
LHR;EGLL;London (Heathrow);United Kingdom
SEN;EGMC;Southend;United Kingdom
LYX;EGMD;Lydd;United Kingdom
MSE;EGMH;Manston;United Kingdom
CAX;EGNC;Carlisle;United Kingdom
BLK;EGNH;Blackpool;United Kingdom
HUY;EGNJ;Humberside;United Kingdom
BWF;EGNL;Barrow Island (Walney Island);United Kingdom
LBA;EGNM;Leeds (Leeds Bradford);United Kingdom
CEG;EGNR;Hawarden;United Kingdom
IOM;EGNS;Isle Of Man;Isle of Man
NCL;EGNT;Newcastle;United Kingdom
MME;EGNV;Teesside (Durham Tees Valley Airport);United Kingdom
EMA;EGNX;East Midlands (Nottingham East Midlands);United Kingdom
KOI;EGPA;Kirkwall;United Kingdom
LSI;EGPB;Sumburgh;United Kingdom
WIC;EGPC;Wick;United Kingdom
ABZ;EGPD;Aberdeen (Dyce);United Kingdom
INV;EGPE;Inverness;United Kingdom
GLA;EGPF;Glasgow;United Kingdom
EDI;EGPH;Edinburgh;United Kingdom
ILY;EGPI;Islay;United Kingdom
PIK;EGPK;Prestwick;United Kingdom
BEB;EGPL;Benbecula;United Kingdom
SDZ;EGPM;Scatsta;United Kingdom
DND;EGPN;Dundee;United Kingdom
SYY;EGPO;Stornoway;United Kingdom
TRE;EGPU;Tiree;United Kingdom
ADX;EGQL;Leuchars;United Kingdom
LMO;EGQS;Lossiemouth;United Kingdom
CBG;EGSC;Cambridge;United Kingdom
NWI;EGSH;Norwich;United Kingdom
STN;EGSS;London (Stansted);United Kingdom
EXT;EGTE;Exeter;United Kingdom
FZO;EGTG;Bristol (Bristol Filton);United Kingdom
OXF;EGTK;Oxford (Kidlington);United Kingdom
MHZ;EGUN;Mildenhall;United Kingdom
FFD;EGVA;Fairford;United Kingdom
BZZ;EGVN;Brize Norton;United Kingdom
ODH;EGVO;Odiham;United Kingdom
NHT;EGWU;Northolt;United Kingdom
QCY;EGXC;Coningsby;United Kingdom
BEQ;EGXH;Honington;United Kingdom
WTN;EGXW;Waddington;United Kingdom
KNF;EGYM;Marham;United Kingdom
MPN;EGYP;Mount Pleasant;Falkland Islands
AMS;EHAM;Amsterdam (Schiphol);Netherlands
MST;EHBK;Maastricht;Netherlands
EIN;EHEH;Eindhoven;Netherlands
GRQ;EHGG;Groningen (Eelde);Netherlands
DHR;EHKD;De Kooy;Netherlands
LWR;EHLW;Leeuwarden;Netherlands
RTM;EHRD;Rotterdam;Netherlands
UTC;EHSB;Soesterberg;Netherlands
ENS;EHTW;Enschede (Twenthe);Netherlands
LID;EHVB;Valkenburg;Netherlands
WOE;EHWO;Woensdrecht;Netherlands
ORK;EICK;Cork;Ireland
GWY;EICM;Galway;Ireland
DUB;EIDW;Dublin;Ireland
NOC;EIKN;Connaught (Ireland West Knock);Ireland
KIR;EIKY;Kerry;Ireland
SNN;EINN;Shannon;Ireland
SXL;EISG;Sligo;Ireland
WAT;EIWF;Waterford;Ireland
AAR;EKAH;Aarhus;Denmark
BLL;EKBI;Billund;Denmark
CPH;EKCH;Copenhagen (Kastrup);Denmark
EBJ;EKEB;Esbjerg;Denmark
KRP;EKKA;Karup;Denmark
ODE;EKOD;Odense;Denmark
RKE;EKRK;Copenhagen (Roskilde);Denmark
RNN;EKRN;Ronne (Bornholm Ronne);Denmark
SGD;EKSB;Soenderborg (Sonderborg);Denmark
SKS;EKSP;Skrydstrup;Denmark
TED;EKTS;Thisted;Denmark
FAE;EKVG;Vagar;Faroe Islands
STA;EKVJ;Stauning;Denmark
AAL;EKYT;Aalborg;Denmark
LUX;ELLX;Luxemburg (Luxembourg);Luxembourg
AES;ENAL;Alesund (Vigra);Norway
ANX;ENAN;Andoya (Andenes);Norway
ALF;ENAT;Alta;Norway
BNN;ENBN;Bronnoysund (Bronnoy);Norway
BOO;ENBO;Bodo;Norway
BGO;ENBR;Bergen (Flesland);Norway
BJF;ENBS;Batsfjord;Norway
KRS;ENCN;Kristiansand (Kjevik);Norway
BDU;ENDU;Bardufoss;Norway
EVE;ENEV;Harstad/Narvik (Evenes);Norway
VDB;ENFG;Fagernes (Leirin);Norway
FRO;ENFL;Floro;Norway
OSL;ENGM;Oslo (Gardermoen);Norway
HAU;ENHD;Haugesund (Karmoy);Norway
HAA;ENHK;Hasvik;Norway
KSU;ENKB;Kristiansund (Kvernberget);Norway
KKN;ENKR;Kirkenes (Hoybuktmoen);Norway
FAN;ENLI;Farsund (Lista);Norway
MOL;ENML;Molde (Aro);Norway
MJF;ENMS;Mosjoen (Kjaerstad);Norway
LKL;ENNA;Lakselv (Banak);Norway
NTB;ENNO;Notodden;Norway
OLA;ENOL;Orland;Norway
RRS;ENRO;Roros;Norway
RYG;ENRY;Rygge (Moss);Norway
LYR;ENSB;Svalbard (Longyear);Norway
SKE;ENSN;Skien (Geiteryggen);Norway
SRP;ENSO;Stord (Sorstokken);Norway
SSJ;ENST;Sandnessjoen (Stokka);Norway
TOS;ENTC;Tromso (Langnes);Norway
TRF;ENTO;Sandefjord (Torp);Norway
TRD;ENVA;Trondheim (Vaernes);Norway
SVG;ENZV;Stavanger (Sola);Norway
GDN;EPGD;Gdansk (Lech Walesa);Poland
KRK;EPKK;Krakow (Balice);Poland
KTW;EPKT;Katowice (Pyrzowice);Poland
POZ;EPPO;Poznan (Lawica);Poland
RZE;EPRZ;Rzeszow (Jasionka);Poland
SZZ;EPSC;Szczecin (Goleniow);Poland
OSP;EPSK;Slupsk (Redzikowo);Poland
WAW;EPWA;Warsaw (Okecie);Poland
WRO;EPWR;Wroclaw (Strachowice);Poland
IEG;EPZG;Zielona Gora (Babimost);Poland
RNB;ESDF;Ronneby;Sweden
GOT;ESGG;Gothenborg (Landvetter);Sweden
JKG;ESGJ;Joenkoeping (Jonkoping);Sweden
LDK;ESGL;Lidkoping;Sweden
GSE;ESGP;Gothenborg (Save);Sweden
KVB;ESGR;Skovde;Sweden
THN;ESGT;Trollhattan (Trollhattan Vanersborg);Sweden
KSK;ESKK;Karlskoga;Sweden
MXX;ESKM;Mora;Sweden
NYO;ESKN;Stockholm (Skavsta);Sweden
KID;ESMK;Kristianstad;Sweden
JLD;ESML;Landskrona;Sweden
OSK;ESMO;Oskarshamn;Sweden
KLR;ESMQ;Kalkmar (Kalmar);Sweden
MMX;ESMS;Malmoe (Sturup);Sweden
HAD;ESMT;Halmstad;Sweden
VXO;ESMX;Vaxjo (Kronoberg);Sweden
EVG;ESND;Sveg;Sweden
GEV;ESNG;Gallivare;Sweden
HUV;ESNH;Hudiksvall;Sweden
KRF;ESNK;Kramfors (Kramfors Solleftea);Sweden
LYC;ESNL;Lycksele;Sweden
SDL;ESNN;Sundsvall (Sundsvall Harnosand);Sweden
OER;ESNO;Ornskoldsvik;Sweden
KRN;ESNQ;Kiruna;Sweden
SFT;ESNS;Skelleftea;Sweden
UME;ESNU;Umea;Sweden
VHM;ESNV;Vilhelmina;Sweden
AJR;ESNX;Arvidsjaur;Sweden
ORB;ESOE;Orebro;Sweden
VST;ESOW;Vasteras;Sweden
LLA;ESPA;Lulea (Kallax);Sweden
ARN;ESSA;Stockholm (Arlanda);Sweden
BMA;ESSB;Stockholm (Bromma);Sweden
BLE;ESSD;Borlange;Sweden
HLF;ESSF;Hultsfred;Sweden
GVX;ESSK;Gavle;Sweden
LPI;ESSL;Linkoeping (Saab);Sweden
NRK;ESSP;Norrkoeping (Kungsangen);Sweden
VBY;ESSV;Visby;Sweden
SPM;ETAD;Spangdahlem (Spangdahlem Ab);Germany
RMS;ETAR;Ramstein (Ramstein Ab);Germany
GHF;ETEU;Giebelstadt (Giebelstadt Aaf);Germany
ZCN;ETHC;Celle;Germany
ZNF;ETID;Hanau (Hanau Aaf);Germany
GKE;ETNG;Geilenkirchen;Germany
RLG;ETNL;Laage;Germany
FEL;ETSF;Fuerstenfeldbruck (Furstenfeldbruck);Germany
GUT;ETUO;Guetersloh (Gutersloh);Germany
ALJ;FAAB;Alexander Bay;South Africa
AGZ;FAAG;Aggeneys;South Africa
BIY;FABE;Bisho (Bhisho);South Africa
BFN;FABL;Bloemfontein (Bloemfontein Intl);South Africa
CPT;FACT;Cape Town (Cape Town Intl);South Africa
DUR;FADN;Durban (Durban Intl);South Africa
ELS;FAEL;East London;South Africa
GCJ;FAGC;Johannesburg (Grand Central);South Africa
GRJ;FAGG;George;South Africa
HDS;FAHS;Hoedspruit (Hoedspruit Afb);South Africa
JNB;FAJS;Johannesburg (Johannesburg Intl);South Africa
KIM;FAKM;Kimberley;South Africa
KLZ;FAKZ;Kleinsee;South Africa
HLA;FALA;Johannesburg (Lanseria);South Africa
LAY;FALY;Ladysmith;South Africa
MGH;FAMG;Margate;South Africa
MEZ;FAMS;Musina (Messina);South Africa
NCS;FANC;Newcastle;South Africa
DUH;FAOH;Oudtshoorn;South Africa
PLZ;FAPE;Port Elizabeth (Port Elizabeth Intl);South Africa
PHW;FAPH;Phalaborwa;South Africa
PTG;FAPI;Polokwane (Polokwane International);South Africa
PZB;FAPM;Pietermaritzburg;South Africa
NTY;FAPN;Pilanesberg (Pilanesberg Intl);South Africa
UTW;FAQT;Queenstown;South Africa
RCB;FARB;Richard's Bay (Richards Bay);South Africa
SBU;FASB;Springbok;South Africa
SIS;FASS;Sishen;South Africa
SZK;FASZ;Skukuza;South Africa
LTA;FATZ;Tzaneen;South Africa
ULD;FAUL;Ulundi (Prince Mangosuthu Buthelezi);South Africa
UTN;FAUP;Upington;South Africa
UTT;FAUT;Umtata (Mthatha);South Africa
VRU;FAVB;Vryburg;South Africa
VIR;FAVG;Durban (Virginia);South Africa
PRY;FAWB;Pretoria (Wonderboom);South Africa
WEL;FAWM;Welkom;South Africa
FRW;FBFT;Francistown;Botswana
JWA;FBJW;Jwaneng;Botswana
BBK;FBKE;Kasane;Botswana
MUB;FBMN;Maun;Botswana
GBE;FBSK;Gaberone (Sir Seretse Khama Intl);Botswana
PKW;FBSP;Selebi-phikwe (Selebi Phikwe);Botswana
BZV;FCBB;Brazzaville (Maya Maya);Congo (Brazzaville)
FTX;FCOO;Owando;Congo (Kinshasa)
OUE;FCOU;Ouesso;Congo (Kinshasa)
PNR;FCPP;Pointe-noire (Pointe Noire);Congo (Brazzaville)
MTS;FDMS;Manzini (Matsapha);Swaziland
BGF;FEFF;Bangui (Bangui M Poko);Central African Republic
BBT;FEFT;Berberati;Central African Republic
BSG;FGBT;Bata;Equatorial Guinea
SSG;FGSL;Malabo;Equatorial Guinea
MRU;FIMP;Plaisance (Sir Seewoosagur Ramgoolam Intl);Mauritius
RRG;FIMR;Rodriguez Island (Plaine Corail);Mauritius
TKC;FKKC;Tiko;Cameroon
DLA;FKKD;Douala;Cameroon
MVR;FKKL;Maroua (Salak);Cameroon
FOM;FKKM;Foumban (Foumban Nkounja);Cameroon
NGE;FKKN;N'gaoundere (Ngaoundere);Cameroon
GOU;FKKR;Garoua;Cameroon
BFX;FKKU;Bafoussam;Cameroon
BPC;FKKV;Bamenda;Cameroon
YAO;FKKY;Yaounde (Yaounde Ville);Cameroon
LVI;FLLI;Livingstone;Zambia
LUN;FLLS;Lusaka (Lusaka Intl);Zambia
MFU;FLMF;Mfuwe;Zambia
NLA;FLND;Ndola;Zambia
KIW;FLSO;Southdowns;Zambia
HAH;FMCH;Moroni (Prince Said Ibrahim);Comoros
NWA;FMCI;Moheli (Bandaressalam);Comoros
AJN;FMCV;Anjouan (Ouani);Comoros
DZA;FMCZ;Dzaoudzi (Dzaoudzi Pamandzi);Mayotte
RUN;FMEE;St.-denis (St Denis Gillot);Reunion
ZSE;FMEP;St.-pierre (St Pierre Pierrefonds);Reunion
TNR;FMMI;Antananarivo (Ivato);Madagascar
ZVA;FMMN;Miandrivazo;Madagascar
SMS;FMMS;Sainte Marie;Madagascar
TMM;FMMT;Toamasina;Madagascar
MOQ;FMMV;Morondava;Madagascar
DIE;FMNA;Antsiranana (Arrachart);Madagascar
WMR;FMNC;Mananara (Avaratra);Madagascar
ZWA;FMND;Andapa;Madagascar
AMB;FMNE;Ambilobe;Madagascar
ANM;FMNH;Antalaha (Antsirabato);Madagascar
HVA;FMNL;Analalava;Madagascar
MJN;FMNM;Mahajanga (Philibert Tsiranana);Madagascar
NOS;FMNN;Nosy-be (Fascene);Madagascar
BPY;FMNQ;Besalampy;Madagascar
WMN;FMNR;Maroantsetra;Madagascar
SVB;FMNS;Sambava;Madagascar
VOH;FMNV;Vohemar (Vohimarina);Madagascar
WAI;FMNW;Antsohihy (Ambalabe);Madagascar
FTU;FMSD;Tolagnaro;Madagascar
WFI;FMSF;Fianarantsoa;Madagascar
RVA;FMSG;Farafangana;Madagascar
WVK;FMSK;Manakara;Madagascar
MNJ;FMSM;Mananjary;Madagascar
MXM;FMSR;Morombe;Madagascar
TLE;FMST;Toliara;Madagascar
SSY;FNBC;M'banza-congo (Mbanza Congo);Angola
BUG;FNBG;Benguela;Angola
CAB;FNCA;Cabinda;Angola
CPX;TJCP;Culebra Island (Culebra Airport);Puerto Rico
NOV;FNHU;Huambo;Angola
SVP;FNKU;Kuito;Angola
LAD;FNLU;Luanda (Luanda 4 De Fevereiro);Angola
MEG;FNMA;Malanje;Angola
SPP;FNME;Menongue;Angola
GXG;FNNG;Negage;Angola
PBN;FNPA;Porto Amboim;Angola
VHC;FNSA;Saurimo;Angola
SZA;FNSO;Soyo;Angola
SDD;FNUB;Lubango;Angola
LUO;FNUE;Luena;Angola
UGO;FNUG;Uige;Angola
XGN;FNXA;Xangongo;Angola
OYE;FOGO;Oyem;Gabon
OKN;FOGQ;Okondja;Gabon
LBQ;FOGR;Lambarene;Gabon
BMM;FOOB;Bitam;Gabon
POG;FOOG;Port Gentil;Gabon
OMB;FOOH;Omboue Hospial (Omboue Hopital);Gabon
MKU;FOOK;Makokou;Gabon
LBV;FOOL;Libreville (Leon M Ba);Gabon
MVB;FOON;Franceville (Mvengue);Gabon
PCP;FPPR;Principe;Sao Tome and Principe
TMS;FPST;Sao Tome (Sao Tome Intl);Sao Tome and Principe
BEW;FQBR;Beira;Mozambique
INH;FQIN;Inhambane;Mozambique
VXC;FQLC;Lichinga;Mozambique
MPM;FQMA;Maputo;Mozambique
MZB;FQMP;Mocimboa Da Praia;Mozambique
MNC;FQNC;Nacala;Mozambique
APL;FQNP;Nampula;Mozambique
POL;FQPB;Pemba;Mozambique
UEL;FQQL;Quelimane;Mozambique
TET;FQTT;Tete (Tete Chingodzi);Mozambique
VNX;FQVL;Vilankulu (Vilankulo);Mozambique
DES;FSDR;Desroches;Seychelles
SEZ;FSIA;Mahe (Seychelles Intl);Seychelles
PRI;FSPP;Praslin;Seychelles
AEH;FTTC;Abeche;Chad
MQQ;FTTD;Moundou;Chad
NDJ;FTTJ;N'djamena (Ndjamena Hassan Djamous);Chad
FYT;FTTY;Faya-largeau (Faya Largeau);Chad
BUQ;FVBU;Bulawayo (J M Nkomo Intl);Zimbabwe
BFO;FVCZ;Chiredzi (Buffalo Range);Zimbabwe
VFA;FVFA;Victoria Falls (Victoria Falls Intl);Zimbabwe
HRE;FVHA;Harare (Harare Intl);Zimbabwe
KAB;FVKB;Kariba (Kariba Intl);Zimbabwe
MVZ;FVMV;Masvingo (Masvingo Intl);Zimbabwe
GWE;FVTL;Gwert (Gweru Thornhill);Zimbabwe
WKM;FVWN;Hwange National Park;Zimbabwe
BLZ;FWCL;Blantyre (Chileka Intl);Malawi
KGJ;FWKA;Karonga;Malawi
LLW;FWKI;Lilongwe (Kamuzu Intl);Malawi
ZZU;FWUU;Mzuzu;Malawi
MSU;FXMM;Maseru (Moshoeshoe I Intl);Lesotho
FIH;FZAA;Kinshasa (Ndjili Intl);Congo (Kinshasa)
NLO;FZAB;Kinshasa (Ndolo);Congo (Kinshasa)
MNB;FZAG;Muanda;Congo (Kinshasa)
FDU;FZBO;Bandoundu (Bandundu);Congo (Kinshasa)
KKW;FZCA;Kikwit;Congo (Kinshasa)
MDK;FZEA;Mbandaka;Congo (Kinshasa)
BDT;FZFD;Gbadolite;Congo (Kinshasa)
GMA;FZFK;Gemena;Congo (Kinshasa)
LIQ;FZGA;Lisala;Congo (Kinshasa)
FKI;FZIA;Kisangani (Kisangani Simisini);Congo (Kinshasa)
IRP;FZJH;Isiro (Matari);Congo (Kinshasa)
BUX;FZKA;Bunia;Congo (Kinshasa)
BKY;FZMA;Bukavu/kavumu (Bukavu Kavumu);Congo (Kinshasa)
GOM;FZNA;Goma;Congo (Kinshasa)
KND;FZOA;Kindu;Congo (Kinshasa)
FBM;FZQA;Lubumashi (Lubumbashi Intl);Congo (Kinshasa)
KWZ;FZQM;Kolwezi;Congo (Kinshasa)
FMI;FZRF;Kalemie;Congo (Kinshasa)
KMN;FZSA;Kamina Base;Congo (Kinshasa)
KGA;FZUA;Kananga;Congo (Kinshasa)
MJM;FZWA;Mbuji-mayi (Mbuji Mayi);Congo (Kinshasa)
BKO;GABS;Bamako (Senou);Mali
GAQ;GAGO;Gao;Mali
KYS;GAKY;Kayes (Kayes Dag Dag);Mali
MZI;GAMB;Mopti (Ambodedjo);Mali
TOM;GATB;Tombouctou;Mali
BJL;GBYD;Banjul (Banjul Intl);Gambia
FUE;GCFV;Fuerteventura;Spain
VDE;GCHI;Hierro;Spain
SPC;GCLA;Santa Cruz De La Palma (La Palma);Spain
LPA;GCLP;Gran Canaria;Spain
ACE;GCRR;Las Palmas (Lanzarote);Spain
TFS;GCTS;Tenerife (Tenerife Sur);Spain
TFN;GCXO;Tenerife (Tenerife Norte);Spain
MLN;GEML;Melilla;Spain
FNA;GFLL;Freetown (Freetown Lungi);Sierra Leone
MLW;GLMR;Monrovia (Monrovia Spriggs Payne);Liberia
ROB;GLRB;Monrovia (Monrovia Roberts Intl);Liberia
AGA;GMAA;Agadir (Inezgane);Morocco
TTA;GMAT;Tan Tan (Plage Blanche);Morocco
FEZ;GMFF;Fes (Saiss);Morocco
ERH;GMFK;Er-rachidia (Moulay Ali Cherif);Morocco
MEK;GMFM;Meknes (Bassatine);Morocco
OUD;GMFO;Oujda (Angads);Morocco
RBA;GMME;Rabat (Sale);Morocco
CMN;GMMN;Casablanca (Mohammed V Intl);Morocco
RAK;GMMX;Marrakech (Menara);Morocco
NNA;GMMY;Kentira (Kenitra);Morocco
OZZ;GMMZ;Ouarzazate;Morocco
AHU;GMTA;Al Hociema (Cherif El Idrissi);Morocco
TTU;GMTN;Tetouan (Saniat Rmel);Morocco
TNG;GMTT;Tanger (Ibn Batouta);Morocco
ZIG;GOGG;Ziguinchor;Senegal
CSK;GOGS;Cap Skiring;Senegal
KLC;GOOK;Kaolack;Senegal
DKR;GOOY;Dakar (Leopold Sedar Senghor Intl);Senegal
XLS;GOSS;St. Louis (Saint Louis);Senegal
BXE;GOTB;Bakel;Senegal
KGG;GOTK;Kedougou;Senegal
TUD;GOTT;Tambacounda;Senegal
IEO;GQNA;Aioun El Atrouss;Mauritania
TIY;GQND;Tidjikja;Mauritania
KFA;GQNF;Kiffa;Mauritania
EMN;GQNI;Nema;Mauritania
KED;GQNK;Kaedi;Mauritania
NKC;GQNN;Nouakschott (Nouakchott);Mauritania
SEY;GQNS;Selibabi (Selibady);Mauritania
ATR;GQPA;Atar;Mauritania
NDB;GQPP;Nouadhibou;Mauritania
FIG;GUFA;Fira (Fria);Guinea
FAA;GUFH;Faranah;Guinea
LEK;GULB;Labe;Guinea
SID;GVAC;Amilcar Cabral (Amilcar Cabral Intl);Cape Verde
BVC;GVBA;Boa Vista (Rabil);Cape Verde
MMO;GVMA;Maio;Cape Verde
SNE;GVSN;Sao Nocolau Island (Preguica);Cape Verde
VXE;GVSV;Sao Vicente Island (Sao Pedro);Cape Verde
ADD;HAAB;Addis Ababa (Bole Intl);Ethiopia
AMH;HAAM;Arba Minch;Ethiopia
AXU;HAAX;Axum;Ethiopia
BJR;HABD;Bahar Dar (Bahir Dar);Ethiopia
DIR;HADR;Dire Dawa (Dire Dawa Intl);Ethiopia
GMB;HAGM;Gambella;Ethiopia
GDQ;HAGN;Gondar;Ethiopia
JIM;HAJM;Jimma;Ethiopia
LLI;HALL;Lalibella;Ethiopia
MQX;HAMK;Makale;Ethiopia
ASO;HASO;Asosa;Ethiopia
BJM;HBBA;Bujumbura (Bujumbura Intl);Burundi
HGA;HCMH;Hargeisa (Egal Intl);Somalia
BBO;HCMI;Berbera;Somalia
KMU;HCMK;Kismayu (Kisimayu);Somalia
ALY;HEAX;Alexandria (Alexandria Intl);Egypt
ABS;HEBL;Abu Simbel;Egypt
CAI;HECA;Cairo (Cairo Intl);Egypt
HRG;HEGN;Hurghada (Hurghada Intl);Egypt
EGR;HEGR;El Gorah (El Gora);Egypt
LXR;HELX;Luxor (Luxor Intl);Egypt
MUH;HEMM;Mersa-matruh (Mersa Matruh);Egypt
PSD;HEPS;Port Said;Egypt
SKV;HESC;St. Catherine (St Catherine Intl);Egypt
ASW;HESN;Aswan (Aswan Intl);Egypt
ELT;HETR;El-tor (El Tor);Egypt
EDL;HKEL;Eldoret (Eldoret Intl);Kenya
KIS;HKKI;Kisumu;Kenya
KTL;HKKT;Kitale;Kenya
LOK;HKLO;Lodwar;Kenya
LAU;HKLU;Lamu (Lamu Manda);Kenya
MBA;HKMO;Mombasa (Mombasa Moi Intl);Kenya
WIL;HKNW;Nairobi (Nairobi Wilson);Kenya
WJR;HKWJ;Wajir;Kenya
GHT;HLGT;Ghat;Libya
AKF;HLKF;Kufra;Libya
BEN;HLLB;Benghazi (Benina);Libya
SEB;HLLS;Sebha;Libya
TIP;HLLT;Tripoli (Tripoli Intl);Libya
LTD;HLTD;Ghadames (Ghadames East);Libya
GYI;HRYG;Gisenyi;Rwanda
KGL;HRYR;Kigali (Kigali Intl);Rwanda
KME;HRZA;Kamembe;Rwanda
DOG;HSDN;Dongola;Sudan
ELF;HSFS;El Fasher (El Fashir);Sudan
KSL;HSKA;Kassala;Sudan
EBD;HSOB;El Obeid;Sudan
JUB;HSSJ;Juba;South Sudan
MAK;HSSM;Malakal;Sudan
KRT;HSSS;Khartoum;Sudan
ARK;HTAR;Arusha;Tanzania
DAR;HTDA;Dar Es Salaam (Mwalimu Julius K Nyerere Intl);Tanzania
DOD;HTDO;Dodoma;Tanzania
IRI;HTIR;Iringa;Tanzania
JRO;HTKJ;Kilimanjaro (Kilimanjaro Intl);Tanzania
LKY;HTLM;Lake Manyara;Tanzania
MYW;HTMT;Mtwara;Tanzania
MWZ;HTMW;Mwanza;Tanzania
PMA;HTPE;Pemba;Tanzania
TGT;HTTG;Tanga;Tanzania
ZNZ;HTZA;Zanzibar;Tanzania
EBB;HUEN;Entebbe (Entebbe Intl);Uganda
SRT;HUSO;Soroti;Uganda
TIA;LATI;Tirana (Tirana Rinas);Albania
BOJ;LBBG;Bourgas (Burgas);Bulgaria
GOZ;LBGO;Gorna Orechovica (Gorna Oryahovitsa);Bulgaria
PDV;LBPD;Plovdiv;Bulgaria
SOF;LBSF;Sofia;Bulgaria
VAR;LBWN;Varna;Bulgaria
LCA;LCLK;Larnaca;Cyprus
PFO;LCPH;Paphos (Pafos Intl);Cyprus
AKT;LCRA;Akrotiri;Cyprus
DBV;LDDU;Dubrovnik;Croatia
OSI;LDOS;Osijek;Croatia
PUY;LDPL;Pula;Croatia
RJK;LDRI;Rijeka;Croatia
SPU;LDSP;Split;Croatia
ZAG;LDZA;Zagreb;Croatia
ZAD;LDZD;Zadar;Croatia
ALC;LEAL;Alicante;Spain
LEI;LEAM;Almeria;Spain
OVD;LEAS;Aviles (Asturias);Spain
ODB;LEBA;Cordoba;Spain
BIO;LEBB;Bilbao;Spain
BCN;LEBL;Barcelona;Spain
BJZ;LEBZ;Badajoz (Talavera La Real);Spain
LCG;LECO;La Coruna (A Coruna);Spain
GRO;LEGE;Gerona (Girona);Spain
GRX;LEGR;Granada;Spain
IBZ;LEIB;Ibiza;Spain
XRY;LEJR;Jerez;Spain
MJV;LELC;Murcia (Murcia San Javier);Spain
PKH;LGHL;Porto Heli (Alexion);Greece
MAD;LEMD;Madrid (Barajas);Spain
AGP;LEMG;Malaga;Spain
MAH;LEMH;Menorca;Spain
OZP;LEMO;Sevilla (Moron Ab);Spain
PNA;LEPP;Pamplona;Spain
REU;LERS;Reus;Spain
SLM;LESA;Salamanca;Spain
EAS;LESO;San Sebastian;Spain
SCQ;LEST;Santiago;Spain
LEU;LESU;Seo De Urgel;Spain
TOJ;LETO;Madrid (Torrejon);Spain
VLC;LEVC;Valencia;Spain
VLL;LEVD;Valladolid;Spain
VIT;LEVT;Vitoria;Spain
VGO;LEVX;Vigo;Spain
SDR;LEXJ;Santander;Spain
ZAZ;LEZG;Zaragoza (Zaragoza Ab);Spain
SVQ;LEZL;Sevilla;Spain
CQF;LFAC;Calais (Calais Dunkerque);France
LTQ;LFAT;Le Tourquet (Le Touquet Paris Plage);France
AGF;LFBA;Agen (La Garenne);France
BOD;LFBD;Bordeaux (Merignac);France
EGC;LFBE;Bergerac (Roumaniere);France
CNG;LFBG;Cognac (Chateaubernard);France
PIS;LFBI;Poitiers (Biard);France
LIG;LFBL;Limoges (Bellegarde);France
NIT;LFBN;Niort (Souche);France
TLS;LFBO;Toulouse (Blagnac);France
PUF;LFBP;Pau (Pau Pyrenees);France
LDE;LFBT;Tarbes (Lourdes);France
ANG;LFBU;Angouleme (Brie Champniers);France
BVE;LFSL;Brive (La Roche);France
PGX;LFBX;Perigueux (Bassillac);France
BIQ;LFBZ;Biarritz-bayonne (Anglet);France
XAC;LFCH;Arcachon (La Teste De Buch);France
LBI;LFCI;Albi (Le Sequestre);France
DCM;LFCK;Castres (Mazamet);France
RDZ;LFCR;Rodez (Marcillac);France
RYN;LFCY;Royan (Medis);France
RCO;LFDN;Rochefort (St Agnant);France
CMR;LFGA;Colmar (Houssen);France
DLE;LFGJ;Dole (Tavaux);France
OBS;LFHO;Aubenas-vals-lanas (Ardeche Meridionale);France
LPY;LFHP;Le Puy (Loudes);France
XBK;LFHS;Bourg (Ceyzeriat);France
XVF;LFHV;Vilefrance (Tarare);France
XMU;LFHY;Moulins (Montbeugny);France
ETZ;LFJL;Metz (Metz Nancy Lorraine);France
BIA;LFKB;Bastia (Poretta);France
CLY;LFKC;Calvi (Saint Catherine);France
FSC;LFKF;Figari (Sud Corse);France
AJA;LFKJ;Ajaccio (Campo Dell Oro);France
SOZ;LFKS;Solenzara;France
AUF;LFLA;Auxerre (Branches);France
CMF;LFLB;Chambery (Aix Les Bains);France
CFE;LFLC;Clermont-Ferrand (Auvergne);France
BOU;LFLD;Bourges;France
XCD;LFLH;Chalon (Champforgeuil);France
QNJ;LFLI;Annemasse;France
LYS;LFLL;Lyon (Saint Exupery);France
QNX;LFLM;Macon (Charnay);France
RNE;LFLO;Roanne (Renaison);France
NCY;LFLP;Annecy (Meythet);France
GNB;LFLS;Grenoble (Saint Geoirs);France
MCU;LFLT;Montlucon (Domerat);France
VAF;LFLU;Valence (Chabeuil);France
VHY;LFLV;Vichy (Charmeil);France
AUR;LFLW;Aurillac;France
CHR;LFLX;Chateauroux (Deols);France
LYN;LFLY;Lyon (Bron);France
QXB;LFMA;Aix-les-milles (Aix Les Milles);France
CEQ;LFMD;Cannes (Mandelieu);France
EBU;LFMH;St-Etienne (Boutheon);France
CCF;LFMK;Carcassonne (Salvaza);France
MRS;LFML;Marseille (Provence);France
NCE;LFMN;Nice (Cote D'Azur);France
PGF;LFMP;Perpignan (Rivesaltes);France
CTT;LFMQ;Le Castellet;France
MPL;LFMT;Montpellier (Mediterranee);France
BZR;LFMU;Beziers (Vias);France
AVN;LFMV;Avignon (Caumont);France
MEN;LFNB;Mende (Brenoux);France
BVA;LFOB;Beauvais (Tille);France
LEH;LFOH;Le Havre (Octeville);France
ORE;LFOJ;Orleans (Bricy);France
XCR;LFOK;Chalons (Vatry);France
URO;LFOP;Rouen (Vallee De Seine);France
TUF;LFOT;Tours (Val De Loire);France
CET;LFOU;Cholet (Le Pontreau);France
LVA;LFOV;Laval (Entrammes);France
LBG;LFPB;Paris (Le Bourget);France
CSF;LFPC;Creil;France
CDG;LFPG;Paris (Charles De Gaulle);France
TNF;LFPN;Toussous-le-noble (Toussus Le Noble);France
ORY;LFPO;Paris (Orly);France
POX;LFPT;Pontoise (Cormeilles En Vexin);France
QYR;LFQB;Troyes (Barberey);France
NVS;LFQG;Nevers (Fourchambault);France
LIL;LFQQ;Lille (Lesquin);France
BES;LFRB;Brest (Guipavas);France
CER;LFRC;Cherbourg (Maupertus);France
DNR;LFRD;Dinard (Pleurtuit);France
GFR;LFRF;Granville;France
DOL;LFRG;Deauville (St Gatien);France
LRT;LFRH;Lorient (Lann Bihoue);France
EDM;LFRI;La Roche-sur-yon (Les Ajoncs);France
CFR;LFRK;Caen (Carpiquet);France
LME;LFRM;Le Mans (Arnage);France
RNS;LFRN;Rennes (St Jacques);France
LAI;LFRO;Lannion;France
UIP;LFRQ;Quimper (Pluguffan);France
NTE;LFRS;Nantes (Nantes Atlantique);France
SBK;LFRT;St.-brieuc Armor (Armor);France
MXN;LFRU;Morlaix (Ploujean);France
VNE;LFRV;Vannes (Meucon);France
SNR;LFRZ;St.-nazaire (Montoir);France
MLH;LFSB;Mulhouse (Bale Mulhouse);France
DIJ;LFSD;Dijon (Longvic);France
MZM;LFSF;Metz (Frescaty);France
EPL;LFSG;Epinal (Mirecourt);France
ENC;LFSN;Nancy (Essey);France
RHE;LFSR;Reims (Champagne);France
SXB;LFST;Strasbourg (Entzheim);France
TLN;LFTH;Hyeres (Le Palyvestre);France
FNI;LFTW;Nimes (Garons);France
MQC;LFVM;Miquelon;Saint Pierre and Miquelon
FSP;LFVP;St.-pierre (St Pierre);Saint Pierre and Miquelon
PYR;LGAD;Andravida;Greece
AGQ;LGAG;Agrinion;Greece
AXD;LGAL;Alexandroupolis (Dimokritos);Greece
VOL;LGBL;Nea Anghialos (Nea Anchialos);Greece
JKH;LGHI;Chios;Greece
IOA;LGIO;Ioannina;Greece
HER;LGIR;Heraklion (Nikos Kazantzakis);Greece
KSO;LGKA;Kastoria (Aristotelis);Greece
KIT;LGKC;Kithira;Greece
EFL;LGKF;Keffallinia (Kefallinia);Greece
KLX;LGKL;Kalamata;Greece
KGS;LGKO;Kos;Greece
AOK;LGKP;Karpathos;Greece
CFU;LGKR;Kerkyra/corfu (Ioannis Kapodistrias Intl);Greece
KSJ;LGKS;Kasos;Greece
KVA;LGKV;Kavala (Megas Alexandros Intl);Greece
KZI;LGKZ;Kozani (Filippos);Greece
LRS;LGLE;Leros;Greece
LXS;LGLM;Limnos;Greece
LRA;LGLR;Larissa (Larisa);Greece
JMK;LGMK;Mykonos (Mikonos);Greece
MJT;LGMT;Mytilini (Mitilini);Greece
PVK;LGPZ;Preveza (Aktio);Greece
RHO;LGRP;Rhodos (Rhodes Diagoras);Greece
GPA;LGRX;Patras (Araxos);Greece
CHQ;LGSA;Chania (Souda);Greece
JSI;LGSK;Skiathos (Alexandros Papadiamantis);Greece
SMI;LGSM;Samos;Greece
LYM;EGMK;Lympne (Ashford);United Kingdom
JTR;LGSR;Thira (Santorini);Greece
JSH;LGST;Sitia;Greece
SKU;LGSY;Skiros;Greece
SKG;LGTS;Thessaloniki (Makedonia);Greece
ZTH;LGZA;Zakynthos (Dionysios Solomos);Greece
BUD;LHBP;Budapest (Ferihegy);Hungary
DEB;LHDC;Debrecen;Hungary
CRV;LIBC;Crotone;Italy
BRI;LIBD;Bari;Italy
FOG;LIBF;Foggia (Gino Lisa);Italy
TAR;LIBG;Grottaglie;Italy
LCC;LIBN;Lecce;Italy
PSR;LIBP;Pescara;Italy
BDS;LIBR;Brindisi (Casale);Italy
SUF;LICA;Lamezia (Lamezia Terme);Italy
CTA;LICC;Catania (Catania Fontanarossa);Italy
LMP;LICD;Lampedusa;Italy
PNL;LICG;Pantelleria;Italy
PMO;LICJ;Palermo;Italy
REG;LICR;Reggio Calabria;Italy
TPS;LICT;Trapani (Trapani Birgi);Italy
NSY;LICZ;Sigonella;Italy
AHO;LIEA;Alghero;Italy
DCI;LIED;Decimomannu;Italy
CAG;LIEE;Cagliari (Elmas);Italy
OLB;LIEO;Olbia (Olbia Costa Smeralda);Italy
TTB;LIET;Tortoli;Italy
MXP;LIMC;Milano (Malpensa);Italy
BGY;LIME;Bergamo (Bergamo Orio Al Serio);Italy
TRN;LIMF;Torino;Italy
ALL;LIMG;Albenga;Italy
GOA;LIMJ;Genoa (Genova Sestri);Italy
LIN;LIML;Milan (Linate);Italy
PMF;LIMP;Parma;Italy
QPZ;LIMS;Piacenza;Italy
CUF;LIMZ;Cuneo (Levaldigi);Italy
AVB;LIPA;Aviano (Aviano Ab);Italy
BZO;LIPB;Bolzano;Italy
BLQ;LIPE;Bologna;Italy
TSF;LIPH;Treviso;Italy
FRL;LIPK;Forli;Italy
VBS;LIPO;Brescia (Montichiari);Italy
TRS;LIPQ;Ronchi De Legionari (Ronchi Dei Legionari);Italy
RMI;LIPR;Rimini;Italy
VIC;LIPT;Vicenza;Italy
QPA;LIPU;Padova;Italy
VRN;LIPX;Villafranca;Italy
VCE;LIPZ;Venice (Venezia Tessera);Italy
SAY;LIQS;Siena (Ampugnano);Italy
CIA;LIRA;Rome (Ciampino);Italy
FCO;LIRF;Rome (Fiumicino);Italy
EBA;LIRJ;Marina Di Campo;Italy
QLT;LIRL;Latina;Italy
NAP;LIRN;Naples (Capodichino);Italy
PSA;LIRP;Pisa;Italy
FLR;LIRQ;Florence (Firenze);Italy
GRS;LIRS;Grosseto;Italy
PEG;LIRZ;Perugia;Italy
LJU;LJLJ;Ljubljana;Slovenia
MBX;LJMB;Maribor;Slovenia
POW;LJPZ;Portoroz;Slovenia
KLV;LKKV;Karlovy Vary;Czech Republic
OSR;LKMT;Ostrava (Mosnov);Czech Republic
PED;LKPD;Pardubice;Czech Republic
PRV;LKPO;Prerov;Czech Republic
PRG;LKPR;Prague (Ruzyne);Czech Republic
BRQ;LKTB;Brno (Turany);Czech Republic
TLV;LLBG;Tel-aviv (Ben Gurion);Israel
BEV;LLBS;Beer-sheba (Teyman);Israel
ETH;LLET;Elat (Eilat);Israel
HFA;LLHA;Haifa;Israel
RPN;LLIB;Rosh Pina (Mahanaim I Ben Yaakov);Israel
VDA;LLOV;Ovda;Israel
SDV;LLSD;Tel-aviv (Sde Dov);Israel
MLA;LMML;Malta (Luqa);Malta
GRZ;LOWG;Graz;Austria
INN;LOWI;Innsbruck;Austria
LNZ;LOWL;Linz;Austria
SZG;LOWS;Salzburg;Austria
VIE;LOWW;Vienna (Schwechat);Austria
SMA;LPAZ;Santa Maria (island) (Santa Maria);Portugal
BGC;LPBG;Braganca;Portugal
FLW;LPFL;Flores;Portugal
FAO;LPFR;Faro;Portugal
GRW;LPGR;Graciosa Island (Graciosa);Portugal
HOR;LPHR;Horta;Portugal
TER;LPLA;Lajes (terceira Island) (Lajes);Portugal
PDL;LPPD;Ponta Delgada;Portugal
PIX;LPPI;Pico;Portugal
OPO;LPPR;Porto;Portugal
PXO;LPPS;Porto Santo;Portugal
LIS;LPPT;Lisbon (Lisboa);Portugal
SJZ;LPSJ;Sao Jorge Island (Sao Jorge);Portugal
VRL;LPVR;Vila Real;Portugal
PDT;KPDT;Pendleton (Eastern Oregon Regional Airport);United States
OMO;LQMO;Mostar;Bosnia and Herzegovina
SJJ;LQSA;Sarajevo;Bosnia and Herzegovina
ARW;LRAR;Arad;Romania
BCM;LRBC;Bacau;Romania
BAY;LRBM;Baia Mare (Tautii Magheraus);Romania
BBU;LRBS;Bucharest (Aurel Vlaicu);Romania
CND;LRCK;Constanta (Mihail Kogalniceanu);Romania
CLJ;LRCL;Cluj-napoca (Cluj Napoca);Romania
CSB;LRCS;Caransebes;Romania
CRA;LRCV;Craiova;Romania
IAS;LRIA;Iasi;Romania
OMR;LROD;Oradea;Romania
OTP;LROP;Bucharest (Henri Coanda);Romania
SBZ;LRSB;Sibiu;Romania
SUJ;LRSM;Satu Mare;Romania
SCV;LRSV;Suceava (Stefan Cel Mare);Romania
TCE;LRTC;Tulcea (Cataloi);Romania
TGM;LRTM;Tirgu Mures (Transilvania Targu Mures);Romania
TSR;LRTR;Timisoara (Traian Vuia);Romania
GVA;LSGG;Geneva (Geneve Cointrin);Switzerland
SIR;LSGS;Sion;Switzerland
LUG;LSZA;Lugano;Switzerland
BRN;LSZB;Bern (Bern Belp);Switzerland
ZRH;LSZH;Zurich;Switzerland
ACH;LSZR;Altenrhein (St Gallen Altenrhein);Switzerland
SMV;LSZS;Samedan;Switzerland
ESB;LTAC;Ankara (Esenboga);Turkey
ANK;LTAD;Ankara (Etimesgut);Turkey
ADA;LTAF;Adana;Turkey
AFY;LTAH;Afyon;Turkey
AYT;LTAI;Antalya;Turkey
GZT;LTAJ;Gaziantep (Oguzeli);Turkey
KYA;LTAN;Konya;Turkey
MZH;LTAP;Merzifon;Turkey
VAS;LTAR;Sivas;Turkey
MLX;LTAT;Malatya (Erhac);Turkey
ASR;LTAU;Kayseri (Erkilet);Turkey
DNZ;LTAY;Denizli (Cardak);Turkey
IST;LTBA;Istanbul (Ataturk);Turkey
BZI;LTBF;Balikesir;Turkey
BDM;LTBG;Bandirma;Turkey
ESK;LTBI;Eskisehir;Turkey
ADB;LTBJ;Izmir (Adnan Menderes);Turkey
IGL;LTBL;Izmir (Cigli);Turkey
DLM;LTBS;Dalaman;Turkey
RIW;KRIW;Riverton WY (Riverton Regional);United States
BXN;LTBV;Bodrum (Imsik);Turkey
EZS;LTCA;Elazig;Turkey
DIY;LTCC;Diyabakir (Diyarbakir);Turkey
ERC;LTCD;Erzincan;Turkey
ERZ;LTCE;Erzurum;Turkey
TZX;LTCG;Trabzon;Turkey
MTJ;KMTJ;Montrose CO (Montrose Regional Airport);United States
VAN;LTCI;Van;Turkey
BAL;LTCJ;Batman;Turkey
KIV;LUKK;Chisinau (Chisinau Intl);Moldova
OHD;LWOH;Ohrid;Macedonia
SKP;LWSK;Skopje;Macedonia
GIB;LXGB;Gibraltar;Gibraltar
BEG;LYBE;Belgrade (Beograd);Serbia
INI;LYNI;Nis;Serbia
TGD;LYPG;Podgorica;Montenegro
PRN;LYPR;Pristina;Serbia
TIV;LYTV;Tivat;Montenegro
BTS;LZIB;Bratislava (M R Stefanik);Slovakia
KSC;LZKZ;Kosice;Slovakia
PZY;LZPP;Piestany;Slovakia
SLD;LZSL;Sliac;Slovakia
TAT;LZTT;Poprad (Tatry);Slovakia
NCA;MBNC;North Caicos;Turks and Caicos Islands
PLS;MBPV;Providenciales;Turks and Caicos Islands
XSC;MBSC;South Caicos;Turks and Caicos Islands
EPS;MDAB;Samana (Arroyo Barril Intl);Dominican Republic
BRX;MDBH;Barahona (Maria Montez Intl);Dominican Republic
LRM;MDLR;La Romana (Casa De Campo Intl);Dominican Republic
PUJ;MDPC;Punta Cana (Punta Cana Intl);Dominican Republic
POP;MDPP;Puerto Plata (Gregorio Luperon Intl);Dominican Republic
SDQ;MDSD;Santo Domingo (Las Americas Intl);Dominican Republic
STI;MDST;Santiago (Cibao Intl);Dominican Republic
CBV;MGCB;Coban;Guatemala
GUA;MGGT;Guatemala City (La Aurora);Guatemala
LCE;MHLC;La Ceiba (Goloson Intl);Honduras
SAP;MHLM;San Pedro Sula (La Mesa Intl);Honduras
GJA;MHNJ;Guanaja;Honduras
RTB;MHRO;Roatan (Juan Manuel Galvez Intl);Honduras
TEA;MHTE;Tela;Honduras
TGU;MHTG;Tegucigalpa (Toncontin Intl);Honduras
OCJ;MKBS;Ocho Rios (Boscobel);Jamaica
KIN;MKJP;Kingston (Norman Manley Intl);Jamaica
MBJ;MKJS;Montego Bay (Sangster Intl);Jamaica
POT;MKKJ;Port Antonio (Ken Jones);Jamaica
KTP;MKTP;Kingston (Tinson Pen);Jamaica
ACA;MMAA;Acapulco (General Juan N Alvarez Intl);Mexico
NTR;MMAN;Monterrey (Del Norte Intl);Mexico
AGU;MMAS;Aguascalientes (Jesus Teran Intl);Mexico
HUX;MMBT;Huatulco (Bahias De Huatulco Intl);Mexico
CVJ;MMCB;Cuernavaca (General Mariano Matamoros);Mexico
CME;MMCE;Ciudad Del Carmen (Ciudad Del Carmen Intl);Mexico
CUL;MMCL;Culiacan (Culiacan Intl);Mexico
CTM;MMCM;Chetumal (Chetumal Intl);Mexico
CEN;MMCN;Ciudad Obregon (Ciudad Obregon Intl);Mexico
CPE;MMCP;Campeche (Ingeniero Alberto Acuna Ongay Intl);Mexico
CJS;MMCS;Ciudad Juarez (Abraham Gonzalez Intl);Mexico
CUU;MMCU;Chihuahua (General R Fierro Villalobos Intl);Mexico
CVM;MMCV;Ciudad Victoria (General Pedro Jose Mendez Intl);Mexico
CZM;MMCZ;Cozumel (Cozumel Intl);Mexico
DGO;MMDO;Durango (Durango Intl);Mexico
TPQ;MMEP;Tepic;Mexico
ESE;MMES;Ensenada;Mexico
GDL;MMGL;Guadalajara (Don Miguel Hidalgo Y Costilla Intl);Mexico
GYM;MMGM;Guaymas (General Jose Maria Yanez Intl);Mexico
TCN;MMHC;Tehuacan;Mexico
HMO;MMHO;Hermosillo (General Ignacio P Garcia Intl);Mexico
CLQ;MMIA;Colima;Mexico
ISJ;MMIM;Isla Mujeres;Mexico
SLW;MMIO;Saltillo (Plan De Guadalupe Intl);Mexico
LZC;MMLC;Lazard Cardenas (Lazaro Cardenas);Mexico
LMM;MMLM;Los Mochis (Valle Del Fuerte Intl);Mexico
BJX;MMLO;Del Bajio (Guanajuato Intl);Mexico
LAP;MMLP;La Paz (General Manuel Marquez De Leon Intl);Mexico
LTO;MMLT;Loreto (Loreto Intl);Mexico
MAM;MMMA;Matamoros (General Servando Canales Intl);Mexico
MID;MMMD;Merida (Licenciado Manuel Crescencio Rejon Int);Mexico
MXL;MMML;Mexicali (General Rodolfo Sanchez Taboada Intl);Mexico
MLM;MMMM;Morelia (General Francisco J Mujica Intl);Mexico
MTT;MMMT;Minatitlan;Mexico
LOV;MMMV;Monclova (Monclova Intl);Mexico
MEX;MMMX;Mexico City (Licenciado Benito Juarez Intl);Mexico
MTY;MMMY;Monterrey (General Mariano Escobedo Intl);Mexico
MZT;MMMZ;Mazatlan (General Rafael Buelna Intl);Mexico
NOG;MMNG;Nogales (Nogales Intl);Mexico
NLD;MMNL;Nuevo Laredo (Quetzalcoatl Intl);Mexico
OAX;MMOX;Oaxaca (Xoxocotlan Intl);Mexico
PAZ;MMPA;Poza Rico (Tajin);Mexico
PBC;MMPB;Puebla (Hermanos Serdan Intl);Mexico
PCA;MMPC;Pachuca (Ingeniero Juan Guillermo Villasana);Mexico
PPE;MMPE;Punta Penasco (Puerto Penasco);Mexico
PDS;MMPG;Piedras Negras (Piedras Negras Intl);Mexico
UPN;MMPN;Uruapan (Licenciado Y Gen Ignacio Lopez Rayon);Mexico
PVR;MMPR;Puerto Vallarta (Licenciado Gustavo Diaz Ordaz Intl);Mexico
PXM;MMPS;Puerto Escondido (Puerto Escondido Intl);Mexico
QRO;MMQT;Queretaro (Queretaro Intercontinental);Mexico
REX;MMRX;Reynosa (General Lucio Blanco Intl);Mexico
SJD;MMSD;San Jose Del Cabo (Los Cabos Intl);Mexico
SLP;MMSP;San Luis Potosi (Ponciano Arriaga Intl);Mexico
TXA;MMTA;Tlaxcala;Mexico
TRC;MMTC;Torreon (Torreon Intl);Mexico
TGZ;MMTG;Tuxtla Gutierrez (Angel Albino Corzo);Mexico
TIJ;MMTJ;Tijuana (General Abelardo L Rodriguez Intl);Mexico
TAM;MMTM;Tampico (General Francisco Javier Mina Intl);Mexico
TSL;MMTN;Tamuin;Mexico
TLC;MMTO;Toluca (Licenciado Adolfo Lopez Mateos Intl);Mexico
TAP;MMTP;Tapachula (Tapachula Intl);Mexico
CUN;MMUN;Cancun (Cancun Intl);Mexico
VSA;MMVA;Villahermosa (C P A Carlos Rovirosa Intl);Mexico
VER;MMVR;Vera Cruz (General Heriberto Jara Intl);Mexico
ZCL;MMZC;Zacatecas (General Leobardo C Ruiz Intl);Mexico
ZIH;MMZH;Zihuatanejo (Ixtapa Zihuatanejo Intl);Mexico
ZMM;MMZM;Zamora;Mexico
ZLO;MMZO;Manzanillo (Playa De Oro Intl);Mexico
BEF;MNBL;Bluefields;Nicaragua
MGA;MNMG;Managua (Managua Intl);Nicaragua
PUZ;MNPC;Puerto Cabezas;Nicaragua
BOC;MPBO;Bocas Del Toro (Bocas Del Toro Intl);Panama
CHX;MPCH;Changuinola (Cap Manuel Nino Intl);Panama
DAV;MPDA;David (Enrique Malek Intl);Panama
HOW;MPHO;Howard;Panama
PAC;MPMG;Panama (Marcos A Gelabert Intl);Panama
PTY;MPTO;Panama City (Tocumen Intl);Panama
OTR;MRCC;Coto 47;Costa Rica
GLF;MRGF;Golfito;Costa Rica
LIR;MRLB;Liberia (Daniel Oduber Quiros Intl);Costa Rica
LIO;MRLM;Limon (Limon Intl);Costa Rica
NOB;MRNS;Nosara Beach (Nosara);Costa Rica
SJO;MROC;San Jose (Juan Santamaria Intl);Costa Rica
PMZ;MRPM;Palmar Sur;Costa Rica
XQP;MRQP;Quepos (La Managua);Costa Rica
SAL;MSLP;San Salvador (El Salvador Intl);El Salvador
CAP;MTCH;Cap Haitien (Cap Haitien Intl);Haiti
PAP;MTPP;Port-au-prince (Toussaint Louverture Intl);Haiti
BCA;MUBA;Baracoa Playa (Gustavo Rizo);Cuba
BYM;MUBY;Bayamo (Carlos Manuel De Cespedes);Cuba
AVI;MUCA;Ciego De Avila (Maximo Gomez);Cuba
CFG;MUCF;Cienfuegos (Jaime Gonzalez);Cuba
CYO;MUCL;Cayo Largo del Sur (Vilo Acuna Intl);Cuba
CMW;MUCM;Camaguey (Ignacio Agramonte Intl);Cuba
SCU;MUCU;Santiago De Cuba (Antonio Maceo Intl);Cuba
GAO;MUGT;Guantanamo (Mariana Grajales);Cuba
HAV;MUHA;Havana (Jose Marti Intl);Cuba
HOG;MUHG;Holguin (Frank Pais Intl);Cuba
LCL;MULM;La Coloma;Cuba
MOA;MUMO;Moa (Orestes Acosta);Cuba
MZO;MUMZ;Manzanillo (Sierra Maestra);Cuba
GER;MUNG;Nueva Gerona (Rafael Cabrera);Cuba
SNU;MUSC;Santa Clara (Abel Santamaria);Cuba
VRA;MUVR;Varadero (Juan Gualberto Gomez Intl);Cuba
VTU;MUVT;Las Tunas (Hermanos Ameijeiras);Cuba
CYB;MWCB;Cayman Barac (Gerrard Smith Intl);Cayman Islands
GCM;MWCR;Georgetown (Owen Roberts Intl);Cayman Islands
ASD;MYAF;Andros Town (Fresh Creek);Bahamas
MHH;MYAM;Marsh Harbor (Marsh Harbour);Bahamas
SAQ;MYAN;San Andros;Bahamas
AXP;MYAP;Spring Point;Bahamas
TCB;MYAT;Treasure Cay;Bahamas
CCZ;MYBC;Chub Cay;Bahamas
BIM;MYBS;Alice Town (South Bimini);Bahamas
GGT;MYEF;Great Exuma (Exuma Intl);Bahamas
ELH;MYEH;North Eleuthera;Bahamas
GHB;MYEM;Governor's Harbor (Governors Harbour);Bahamas
RSD;MYER;Rock Sound;Bahamas
FPO;MYGF;Freeport (Grand Bahama Intl);Bahamas
IGA;MYIG;Matthew Town;Bahamas
LGI;MYLD;Dead Man's Cay (Deadmans Cay);Bahamas
SML;MYLS;Stella Maris;Bahamas
MYG;MYMM;Mayaguana;Bahamas
NAS;MYNN;Nassau (Lynden Pindling Intl);Bahamas
ZSA;MYSM;Cockburn Town (San Salvador);Bahamas
BZE;MZBZ;Belize City (Philip S W Goldson Intl);Belize
AIT;NCAI;Aitutaki;Cook Islands
RAR;NCRG;Avarua (Rarotonga Intl);Cook Islands
NAN;NFFN;Nandi (Nadi Intl);Fiji
SUV;NFNA;Nausori (Nausori Intl);Fiji
TBU;NFTF;Tongatapu (Fua Amotu Intl);Tonga
VAV;NFTV;Vava'u (Vavau Intl);Tonga
TRW;NGTA;Tarawa (Bonriki Intl);Kiribati
TBF;NGTE;Tabiteuea North;Kiribati
MTL;YMND;Maitland (Maitland Airport);Australia
WLS;NLWW;Wallis;Wallis and Futuna
APW;NSFA;Faleolo (Faleolo Intl);Samoa
PPG;NSTU;Pago Pago (Pago Pago Intl);American Samoa
RUR;NTAR;Rurutu;French Polynesia
TUB;NTAT;Tubuai;French Polynesia
AAA;NTGA;Anaa;French Polynesia
TIH;NTGC;Tikehau;French Polynesia
REA;NTGE;Reao;French Polynesia
FAV;NTGF;Fakarava;French Polynesia
XMH;NTGI;Manihi;French Polynesia
GMR;NTGJ;Totegegie;French Polynesia
KKR;NTGK;Kaukura Atoll (Kaukura);French Polynesia
MKP;NTGM;Makemo;French Polynesia
PKP;NTGP;Puka Puka;French Polynesia
TKP;NTGT;Takapoto;French Polynesia
AXR;NTGU;Arutua;French Polynesia
MVT;NTGV;Mataiva;French Polynesia
TKX;NTKR;Takaroa;French Polynesia
NHV;NTMD;Nuku Hiva;French Polynesia
BOB;NTTB;Bora Bora;French Polynesia
RGI;NTTG;Rangiroa;French Polynesia
HUH;NTTH;Huahine Island (Huahine);French Polynesia
MOZ;NTTM;Moorea;French Polynesia
HOI;NTTO;Hao Island (Hao);French Polynesia
MAU;NTTP;Maupiti;French Polynesia
RFP;NTTR;Raiatea Island (Raiatea);French Polynesia
VLI;NVVV;Port-vila (Port Vila Bauerfield);Vanuatu
KNQ;NWWD;Kone;New Caledonia
KOC;NWWK;Koumac;New Caledonia
LIF;NWWL;Lifou;New Caledonia
GEA;NWWM;Noumea (Magenta);New Caledonia
MEE;NWWR;Mare;New Caledonia
TOU;NWWU;Touho;New Caledonia
UVE;NWWV;Ouvea;New Caledonia
NOU;NWWW;Noumea (La Tontouta);New Caledonia
AKL;NZAA;Auckland (Auckland Intl);New Zealand
TUO;NZAP;Taupo;New Zealand
AMZ;NZAR;Ardmore;New Zealand
CHC;NZCH;Christchurch (Christchurch Intl);New Zealand
CHT;NZCI;Chatham Island (Chatham Islands);New Zealand
DUD;NZDN;Dunedin;New Zealand
GIS;NZGS;Gisborne;New Zealand
MON;NZGT;Glentanner;New Zealand
HKK;NZHK;Hokitika;New Zealand
HLZ;NZHN;Hamilton;New Zealand
KKE;NZKK;Kerikeri;New Zealand
KAT;NZKT;Kaitaia;New Zealand
ALR;NZLX;Alexandra;New Zealand
GTN;NZMC;Mount Cook;New Zealand
TEU;NZMO;Manapouri;New Zealand
MRO;NZMS;Masterton;New Zealand
NPL;NZNP;New Plymouth;New Zealand
NSN;NZNS;Nelson;New Zealand
IVC;NZNV;Invercargill;New Zealand
OAM;NZOU;Oamaru;New Zealand
PMR;NZPM;Palmerston North;New Zealand
PPQ;NZPP;Paraparaumu;New Zealand
ZQN;NZQN;Queenstown International (Queenstown);New Zealand
ROT;NZRO;Rotorua;New Zealand
TRG;NZTG;Tauranga;New Zealand
TIU;NZTU;Timaru;New Zealand
BHE;NZWB;Woodbourne;New Zealand
WKA;NZWF;Wanaka;New Zealand
WHK;NZWK;Whakatane;New Zealand
WLG;NZWN;Wellington (Wellington Intl);New Zealand
WRE;NZWR;Whangarei;New Zealand
WSZ;NZWS;Westport;New Zealand
WAG;NZWU;Wanganui;New Zealand
HEA;OAHR;Herat;Afghanistan
JAA;OAJL;Jalalabad;Afghanistan
KBL;OAKB;Kabul (Kabul Intl);Afghanistan
KDH;OAKN;Kandahar;Afghanistan
MMZ;OAMN;Maimama (Maimana);Afghanistan
MZR;OAMS;Mazar-i-sharif (Mazar I Sharif);Afghanistan
UND;OAUZ;Kunduz (Konduz);Afghanistan
BAH;OBBI;Bahrain (Bahrain Intl);Bahrain
AHB;OEAB;Abha;Saudi Arabia
HOF;OEAH;Al-ahsa (Al Ahsa);Saudi Arabia
ABT;OEBA;El-baha (Al Baha);Saudi Arabia
BHH;OEBH;Bisha;Saudi Arabia
DMM;OEDF;Dammam (King Fahd Intl);Saudi Arabia
DHA;OEDR;Dhahran (King Abdulaziz Ab);Saudi Arabia
GIZ;OEGN;Gizan (King Abdullah Bin Abdulaziz);Saudi Arabia
ELQ;OEGS;Gassim;Saudi Arabia
URY;OEGT;Guriat;Saudi Arabia
HAS;OEHL;Hail;Saudi Arabia
JED;OEJN;Jeddah (King Abdulaziz Intl);Saudi Arabia
HBT;OEKK;King Khalid Mil.city (King Khaled Military City);Saudi Arabia
MED;OEMA;Madinah (Prince Mohammad Bin Abdulaziz);Saudi Arabia
EAM;OENG;Nejran;Saudi Arabia
AQI;OEPA;Hafr Al-batin (Qaisumah);Saudi Arabia
RAH;OERF;Rafha;Saudi Arabia
RUH;OERK;Riyadh (King Khaled Intl);Saudi Arabia
RAE;OERR;Arar;Saudi Arabia
SHW;OESH;Sharurah;Saudi Arabia
DGE;YMDG;Mudgee (Mudgee Airport);Australia
SLF;OESL;Sulayel;Saudi Arabia
TUU;OETB;Tabuk;Saudi Arabia
TIF;OETF;Taif;Saudi Arabia
TUI;OETR;Turaif;Saudi Arabia
EJH;OEWJ;Wejh;Saudi Arabia
YNB;OEYN;Yenbo;Saudi Arabia
ABD;OIAA;Abadan;Iran
QMJ;OIAI;Masjed Soleiman (Shahid Asyaee);Iran
MRX;OIAM;Bandar Mahshahr (Mahshahr);Iran
AWZ;OIAW;Ahwaz;Iran
BUZ;OIBB;Bushehr;Iran
KIH;OIBK;Kish Island;Iran
BDH;OIBL;Bandar Lengeh;Iran
KSH;OICC;Bakhtaran (Shahid Ashrafi Esfahani);Iran
SDG;OICS;Sanandaj;Iran
RAS;OIGG;Rasht;Iran
THR;OIII;Teheran (Mehrabad Intl);Iran
BND;OIKB;Bandar Abbas (Bandar Abbass Intl);Iran
KER;OIKK;Kerman;Iran
XBJ;OIMB;Birjand;Iran
WSY;YWHI;Airlie Beach (Whitsunday Airstrip);Australia
RZR;OINR;Ramsar;Iran
SYZ;OISS;Shiraz (Shiraz Shahid Dastghaib Intl);Iran
QSA;LELL;Sabadell (Sabadell Airport);Spain
TBZ;OITT;Tabriz (Tabriz Intl);Iran
AZD;OIYY;Yazd (Yazd Shahid Sadooghi);Iran
ZBR;OIZC;Chah Bahar;Iran
ZAH;OIZH;Zahedan (Zahedan Intl);Iran
AMM;OJAI;Amman (Queen Alia Intl);Jordan
ADJ;OJAM;Amman (Marka Intl);Jordan
AQJ;OJAQ;Aqaba (Aqaba King Hussein Intl);Jordan
OMF;OJMF;Mafraq (King Hussein);Jordan
KWI;OKBK;Kuwait (Kuwait Intl);Kuwait
BEY;OLBA;Beirut (Rafic Hariri Intl);Lebanon
AUH;OMAA;Abu Dhabi (Abu Dhabi Intl);United Arab Emirates
AZI;OMAD;Abu Dhabi (Bateen);United Arab Emirates
MVA;BIRL;Myvatn (Reykjahlid Airport);Iceland
DXB;OMDB;Dubai (Dubai Intl);United Arab Emirates
FJR;OMFJ;Fujeirah (Fujairah Intl);United Arab Emirates
RKT;OMRK;Ras Al Khaimah (Ras Al Khaimah Intl);United Arab Emirates
SHJ;OMSJ;Sharjah (Sharjah Intl);United Arab Emirates
KHS;OOKB;Khasab;Oman
MSH;OOMA;Masirah;Oman
MCT;OOMS;Muscat (Seeb Intl);Oman
SLL;OOSA;Salalah;Oman
TTH;OOTH;Thumrait;Oman
LYP;OPFA;Faisalabad (Faisalabad Intl);Pakistan
GWD;OPGD;Gwadar;Pakistan
GIL;OPGT;Gilgit;Pakistan
KHI;OPKC;Karachi (Jinnah Intl);Pakistan
LHE;OPLA;Lahore (Allama Iqbal Intl);Pakistan
MFG;OPMF;Muzaffarabad;Pakistan
MJD;OPMJ;Moenjodaro;Pakistan
MUX;OPMT;Multan (Multan Intl);Pakistan
WNS;OPNH;Nawabshah;Pakistan
PJG;OPPG;Panjgur;Pakistan
PSI;OPPI;Pasni;Pakistan
PEW;OPPS;Peshawar (Peshawar Intl);Pakistan
UET;OPQT;Quetta;Pakistan
RYK;OPRK;Rahim Yar Khan (Sheikh Zayed);Pakistan
ISB;OPRN;Islamabad (Chaklala);Pakistan
RAZ;OPRT;Rawala Kot (Rawalakot);Pakistan
1RL;K1RL;Point Roberts (Point Roberts Airpark);United States
SKZ;OPSK;Sukkur;Pakistan
SDT;OPSS;Saidu Sharif;Pakistan
SUL;OPSU;Sui;Pakistan
BDN;OPTH;Talhar;Pakistan
PZH;OPZB;Zhob;Pakistan
BSR;ORMM;Basrah (Basrah Intl);Iraq
ALP;OSAP;Aleppo (Aleppo Intl);Syria
DAM;OSDI;Damascus (Damascus Intl);Syria
DEZ;OSDZ;Deire Zor (Deir Zzor);Syria
LTK;OSLK;Latakia (Bassel Al Assad Intl);Syria
PMS;OSPR;Palmyra;Syria
DOH;OTBD;Doha (Doha Intl);Qatar
CIS;PCIS;Canton Island (Canton);Kiribati
ROP;PGRO;Rota (Rota Intl);Northern Mariana Islands
SPN;PGSN;Saipan (Francisco C Ada Saipan Intl);Northern Mariana Islands
UAM;PGUA;Andersen (Andersen Afb);Guam
GUM;PGUM;Agana (Guam Intl);Guam
TIQ;PGWT;West Tinian (Tinian Intl);Northern Mariana Islands
NSO;YSCO;Scone (Scone Airport);Australia
MAJ;PKMJ;Majuro (Marshall Islands Intl);Marshall Islands
KWA;PKWA;Kwajalein (Bucholz Aaf);Marshall Islands
CXI;PLCH;Kiritimati (Cassidy Intl);Kiribati
MDY;PMDY;Midway (Midway Atoll);Midway Islands
TKK;PTKK;Chuuk (Chuuk Intl);Micronesia
PNI;PTPN;Pohnpei (Pohnpei Intl);Micronesia
ROR;PTRO;Babelthuap;Palau
KSA;PTSA;Kosrae;Micronesia
YAP;PTYA;Yap (Yap Intl);Micronesia
KNH;RCBS;Kinmen (Shang Yi);Taiwan
PIF;RCDC;Pingtung (Pingtung South);Taiwan
TTT;RCFN;Fengnin;Taiwan
GNI;RCGI;Green Island (Lyudao);Taiwan
KHH;RCKH;Kaohsiung (Kaohsiung Intl);Taiwan
CYI;RCKU;Chiayi;Taiwan
KYD;RCLY;Lanyu;Taiwan
RMQ;RCMQ;Taichung (Ching Chuang Kang);Taiwan
CES;YCNK;Cessnock (Cessnock Airport);Australia
TNN;RCNN;Tainan;Taiwan
MZG;RCQC;Makung (Magong);Taiwan
TSA;RCSS;Taipei (Sungshan);Taiwan
TPE;RCTP;Taipei (Taoyuan Intl);Taiwan
WOT;RCWA;Wang An;Taiwan
HUN;RCYU;Hualien;Taiwan
NRT;RJAA;Tokyo (Narita Intl);Japan
MMJ;RJAF;Matsumoto;Japan
IBR;RJAH;Ibaraki (Hyakuri);Japan
IWO;RJAW;Iwojima (Iwo Jima);Japan
SHM;RJBD;Nanki-shirahama (Nanki Shirahama);Japan
OBO;RJCB;Obihiro;Japan
CTS;RJCC;Sapporo (New Chitose);Japan
HKD;RJCH;Hakodate;Japan
SPK;RJCJ;Chitose;Japan
MMB;RJCM;Memanbetsu;Japan
SHB;RJCN;Nakashibetsu;Japan
WKJ;RJCW;Wakkanai;Japan
IKI;RJDB;Iki;Japan
UBJ;RJDC;Yamaguchi (Yamaguchi Ube);Japan
TSJ;RJDT;Tsushima;Japan
MBE;RJEB;Monbetsu;Japan
AKJ;RJEC;Asahikawa;Japan
OIR;RJEO;Okushiri;Japan
RIS;RJER;Rishiri Island (Rishiri);Japan
KUM;RJFC;Yakushima;Japan
FUJ;RJFE;Fukue;Japan
FUK;RJFF;Fukuoka;Japan
TNE;RJFG;Tanegashima (New Tanegashima);Japan
KOJ;RJFK;Kagoshima;Japan
KMI;RJFM;Miyazaki;Japan
OIT;RJFO;Oita;Japan
KKJ;RJFR;Kitakyushu (New Kitakyushu);Japan
KMJ;RJFT;Kumamoto;Japan
NGS;RJFU;Nagasaki;Japan
ASJ;RJKA;Amami;Japan
TKN;RJKN;Tokunoshima;Japan
KMQ;RJNK;Kanazawa (Komatsu);Japan
OKI;RJNO;Oki Island (Oki);Japan
TOY;RJNT;Toyama;Japan
HIJ;RJOA;Hiroshima;Japan
OKJ;RJOB;Okayama;Japan
IZO;RJOC;Izumo;Japan
YGJ;RJOH;Miho (Yonago Kitaro);Japan
KCZ;RJOK;Kochi;Japan
MYJ;RJOM;Matsuyama;Japan
ITM;RJOO;Osaka (Osaka Intl);Japan
TTJ;RJOR;Tottori;Japan
TKS;RJOS;Tokushima;Japan
TAK;RJOT;Takamatsu;Japan
AOJ;RJSA;Aomori;Japan
GAJ;RJSC;Yamagata;Japan
HNA;RJSI;Hanamaki;Japan
AXT;RJSK;Akita;Japan
MSJ;RJSM;Misawa (Misawa Ab);Japan
SDJ;RJSS;Sendai;Japan
LKS;LKSZ;Sazena;Czech Republic
HAC;RJTH;Hachijojima;Japan
OIM;RJTO;Oshima;Japan
HND;RJTT;Tokyo (Tokyo Intl);Japan
OKO;RJTY;Yokota (Yokota Ab);Japan
KWJ;RKJJ;Kwangju (Gwangju);South Korea
RSU;RKJY;Yeosu;South Korea
SHO;RKND;Sokch'o (Sokcho);South Korea
KAG;RKNN;Kangnung (Gangneung);South Korea
CJU;RKPC;Cheju (Jeju Intl);South Korea
PUS;RKPK;Busan (Gimhae Intl);South Korea
USN;RKPU;Ulsan;South Korea
SSN;RKSM;Seoul East (Seoul Ab);South Korea
OSN;RKSO;Osan (Osan Ab);South Korea
GMP;RKSS;Seoul (Gimpo);South Korea
KPO;RKTH;Pohang;South Korea
TAE;RKTN;Taegu (Daegu Ab);South Korea
YEC;RKTY;Yechon (Yecheon);South Korea
OKA;ROAH;Okinawa (Naha);Japan
DNA;RODN;Kadena (Kadena Ab);Japan
ISG;ROIG;Ishigaki;Japan
UEO;ROKJ;Kumejima;Japan
MMD;ROMD;Minami Daito;Japan
MMY;ROMY;Miyako;Japan
KTD;RORK;Kitadaito;Japan
SHI;RORS;Shimojishima;Japan
RNJ;RORY;Yoron;Japan
OGN;ROYN;Yonaguni Jima (Yonaguni);Japan
MNL;RPLL;Manila (Ninoy Aquino Intl);Philippines
MYP;UTAM;Mary (Mary Airport);Turkmenistan
CBO;RPMC;Cotabato;Philippines
CGY;RPML;Ladag (Cagayan De Oro);Philippines
PAG;RPMP;Pagadian;Philippines
ZAM;RPMZ;Zamboanga (Zamboanga Intl);Philippines
BAG;RPUB;Baguio;Philippines
TAC;RPVA;Tacloban (Daniel Z Romualdez);Philippines
BCD;RPVB;Bacolod;Philippines
DGT;RPVD;Dumaguete;Philippines
ILO;RPVI;Iloilo;Philippines
KLO;RPVK;Kalibo;Philippines
PPS;RPVP;Puerto Princesa;Philippines
SJI;RPVS;San Jose (Antique);Philippines
COC;SAAC;Concordia (Comodoro Pierrestegui);Argentina
GHU;SAAG;Gualeguaychu;Argentina
PRA;SAAP;Parana (General Urquiza);Argentina
ROS;SAAR;Rosario;Argentina
SFN;SAAV;Santa Fe (Sauce Viejo);Argentina
AEP;SABE;Buenos Aires (Aeroparque Jorge Newbery);Argentina
COR;SACO;Cordoba (Ambrosio L V Taravella);Argentina
LPG;SADL;La Plata;Argentina
MDZ;SAME;Mendoza (El Plumerillo);Argentina
LGS;SAMM;Malargue;Argentina
AFA;SAMR;San Rafael;Argentina
CTC;SANC;Catamarca;Argentina
SDE;SANE;Santiago Del Estero;Argentina
IRJ;SANL;La Rioja;Argentina
TUC;SANT;Tucuman (Teniente Benjamin Matienzo);Argentina
UAQ;SANU;San Julian (San Juan);Argentina
RCU;SAOC;Rio Cuarto (Rio Cuarto Area De Material);Argentina
VDR;SAOD;Villa Dolores;Argentina
LUQ;SAOU;San Luis;Argentina
CNQ;SARC;Corrientes;Argentina
RES;SARE;Resistencia;Argentina
FMA;SARF;Formosa;Argentina
IGR;SARI;Iguazu Falls (Cataratas Del Iguazu);Argentina
AOL;SARL;Paso De Los Libres;Argentina
PSS;SARP;Posadas;Argentina
SLA;SASA;Salta;Argentina
JUJ;SASJ;Jujuy;Argentina
ORA;SASO;Oran;Argentina
EHL;SAVB;El Bolson;Argentina
CRD;SAVC;Comodoro Rivadavia;Argentina
EQS;SAVE;Esquel;Argentina
REL;SAVT;Trelew (Almirante Zar);Argentina
VDM;SAVV;Viedma (Gobernador Castello);Argentina
PMY;SAVY;Puerto Madryn (El Tehuelche);Argentina
PUD;SAWD;Puerto Deseado;Argentina
RGA;SAWE;Rio Grande;Argentina
RGL;SAWG;Rio Gallegos;Argentina
USH;SAWH;Ushuaia (Ushuaia Malvinas Argentinas);Argentina
ULA;SAWJ;San Julian;Argentina
PMQ;SAWP;Perito Moreno;Argentina
RZA;SAWU;Santa Cruz;Argentina
BHI;SAZB;Bahia Blanca (Comandante Espora);Argentina
MDQ;SAZM;Mar Del Plata;Argentina
NQN;SAZN;Neuquen (Presidente Peron);Argentina
RSA;SAZR;Santa Rosa;Argentina
BRC;SAZS;San Carlos De Bariloche;Argentina
TDL;SAZT;Tandil;Argentina
VLG;SAZV;Villa Gesell;Argentina
CPC;SAZY;San Martin Des Andes (Aviador C Campos);Argentina
CDJ;SBAA;Conceicao Do Araguaia;Brazil
AQA;SBAQ;Araracuara (Araraquara);Brazil
AJU;SBAR;Aracaju (Santa Maria);Brazil
AFL;SBAT;Alta Floresta;Brazil
ARU;SBAU;Aracatuba;Brazil
BEL;SBBE;Belem (Val De Cans Intl);Brazil
BGX;SBBG;Bage (Comandante Gustavo Kraemer);Brazil
PLU;SBBH;Belo Horizonte (Pampulha Carlos Drummond De Andrade);Brazil
BFH;SBBI;Curitiba (Bacacheri);Brazil
BSB;SBBR;Brasilia (Presidente Juscelino Kubitschek);Brazil
BAU;SBBU;Bauru;Brazil
BVB;SBBV;Boa Vista;Brazil
CAC;SBCA;Cascavel;Brazil
CNF;SBCF;Belo Horizonte (Tancredo Neves Intl);Brazil
CGR;SBCG;Campo Grande;Brazil
XAP;SBCH;Chapeco;Brazil
CLN;SBCI;Carolina;Brazil
CCM;SBCM;Criciuma (Forquilhinha);Brazil
CAW;SBCP;Campos (Bartolomeu Lisandro);Brazil
CMG;SBCR;Corumba (Corumba Intl);Brazil
CWB;SBCT;Curitiba (Afonso Pena);Brazil
CRQ;SBCV;Caravelas;Brazil
CXJ;SBCX;Caxias Do Sul (Campo Dos Bugres);Brazil
CGB;SBCY;Cuiaba (Marechal Rondon);Brazil
CZS;SBCZ;Cruzeiro do Sul;Brazil
PPB;SBDN;President Prudente (Presidente Prudente);Brazil
MAO;SBEG;Manaus (Eduardo Gomes Intl);Brazil
IGU;SBFI;Foz Do Iguacu (Cataratas Intl);Brazil
FLN;SBFL;Florianopolis (Hercilio Luz);Brazil
FEN;SBFN;Fernando Do Noronha (Fernando De Noronha);Brazil
FOR;SBFZ;Fortaleza (Pinto Martins Intl);Brazil
GIG;SBGL;Rio De Janeiro (Galeao Antonio Carlos Jobim);Brazil
GYN;SBGO;Goiania (Santa Genoveva);Brazil
GRU;SBGR;Sao Paulo (Guarulhos Gov Andre Franco Montouro);Brazil
ATM;SBHT;Altamira;Brazil
IOS;SBIL;Ilheus;Brazil
IPN;SBIP;Ipatinga (Usiminas);Brazil
IMP;SBIZ;Imperatriz (Prefeito Renato Moreira);Brazil
JDF;SBJF;Juiz De Fora (Francisco De Assis);Brazil
JPA;SBJP;Joao Pessoa (Presidente Castro Pinto);Brazil
JOI;SBJV;Joinville (Lauro Carneiro De Loyola);Brazil
CPV;SBKG;Campina Grande (Presidente Joao Suassuna);Brazil
VCP;SBKP;Campinas (Viracopos);Brazil
LIP;SBLN;Lins;Brazil
LDB;SBLO;Londrina;Brazil
LAZ;SBLP;Bom Jesus Da Lapa;Brazil
MAB;SBMA;Maraba;Brazil
MGF;SBMG;Maringa (Regional De Maringa Silvio Name Junior);Brazil
MOC;SBMK;Montes Claros (Mario Ribeiro);Brazil
GRM;KCKC;Grand Marais (Grand Marais Cook County Airport);United States
MCZ;SBMO;Maceio (Zumbi Dos Palmares);Brazil
MCP;SBMQ;Macapa;Brazil
MNX;SBMY;Manicore;Brazil
NVT;SBNF;Navegantes (Ministro Victor Konder Intl);Brazil
GEL;SBNM;Santo Angelo;Brazil
NAT;SBNT;Natal (Augusto Severo);Brazil
POA;SBPA;Porto Alegre (Salgado Filho);Brazil
POO;SBPC;Pocos De Caldas;Brazil
PFB;SBPF;Passo Fundo (Lauro Kurtz);Brazil
PET;SBPK;Pelotas;Brazil
PNZ;SBPL;Petrolina (Senador Nilo Coelho);Brazil
PNB;SBPN;Porto Nacional;Brazil
PMG;SBPP;Ponta Pora;Brazil
PVH;SBPV;Porto Velho (Governador Jorge Teixeira De Oliveira);Brazil
RBR;SBRB;Rio Branco (PlÃ¡cido de Castro);Brazil
REC;SBRF;Recife (Guararapes Gilberto Freyre Intl);Brazil
RIG;SBRG;Rio Grande;Brazil
SDU;SBRJ;Rio De Janeiro (Santos Dumont);Brazil
RAO;SBRP;Ribeirao Preto (Leite Lopes);Brazil
STU;SBSC;Rio De Janeiro (Santa Cruz);Brazil
SJK;SBSJ;Sao Jose Dos Campos (Professor Urbano Ernesto Stumpf);Brazil
SLZ;SBSL;Sao Luis (Marechal Cunha Machado Intl);Brazil
CGH;SBSP;Sao Paulo (Congonhas);Brazil
SJP;SBSR;Sao Jose Do Rio Preto;Brazil
SSZ;SBST;Santos (Base Aerea De Santos);Brazil
SSA;SBSV;Salvador (Deputado Luis Eduardo Magalhaes);Brazil
TMT;SBTB;Oriximina (Trombetas);Brazil
THE;SBTE;Teresina (Senador Petronio Portella);Brazil
TFF;SBTF;Tefe;Brazil
TBT;SBTT;Tabatinga;Brazil
TUR;SBTU;Tucurui;Brazil
PAV;SBUF;Paulo Alfonso (Paulo Afonso);Brazil
URG;SBUG;Uruguaiana (Rubem Berta);Brazil
UDI;SBUL;Uberlandia (Ten Cel Av Cesar Bombonato);Brazil
UBA;SBUR;Uberaba;Brazil
VAG;SBVG;Varginha (Major Brigadeiro Trompowsky);Brazil
BVH;SBVH;Vilhena;Brazil
VIX;SBVT;Vitoria (Goiabeiras);Brazil
QPS;SBYS;Piracununga (Campo Fontenelle);Brazil
ARI;SCAR;Arica (Chacalluta);Chile
BBA;SCBA;Balmaceda;Chile
CCH;SCCC;Chile Chico;Chile
CJC;SCCF;Calama (El Loa);Chile
PUQ;SCCI;Punta Arenas (Carlos Ibanez Del Campo Intl);Chile
GXQ;SCCY;Coyhaique (Teniente Vidal);Chile
IQQ;SCDA;Iquique (Diego Aracena Intl);Chile
SCL;SCEL;Santiago (Arturo Merino Benitez Intl);Chile
ANF;SCFA;Antofagasta (Cerro Moreno Intl);Chile
LSQ;SCGE;Los Angeles (Maria Dolores);Chile
CCP;SCIE;Concepcion (Carriel Sur Intl);Chile
IPC;SCIP;Easter Island (Mataveri Intl);Chile
ZOS;SCJO;Osorno (Canal Bajo Carlos Hott Siebert);Chile
LSC;SCSE;La Serena (La Florida);Chile
ZCO;SCTC;Temuco (Maquehue);Chile
PMC;SCTE;Puerto Montt (El Tepual Intl);Chile
WCH;SCTN;Chaiten;Chile
ZAL;SCVD;Valdivia (Pichoy);Chile
ATF;SEAM;Ambato (Chachoan);Ecuador
OCC;SECO;Coca (Francisco De Orellana);Ecuador
CUE;SECU;Cuenca (Mariscal Lamar);Ecuador
GPS;SEGS;Galapagos (Seymour);Ecuador
GYE;SEGU;Guayaquil (Jose Joaquin De Olmedo Intl);Ecuador
XMS;SEMC;Macas (Coronel E Carvajal);Ecuador
MCH;SEMH;Machala (General Manuel Serrano);Ecuador
MEC;SEMT;Manta (Eloy Alfaro Intl);Ecuador
PVO;SEPV;Portoviejo (Reales Tamarindos);Ecuador
UIO;SEQU;Quito (Mariscal Sucre Intl);Ecuador
SNC;SESA;Salinas (General Ulpiano Paez);Ecuador
TPC;SETR;Tarapoa;Ecuador
TUA;SETU;Tulcan (Teniente Coronel Luis A Mantilla);Ecuador
ASU;SGAS;Asuncion (Silvio Pettirossi Intl);Paraguay
AXM;SKAR;Armenia (El Eden);Colombia
PUU;SKAS;Puerto Asis (Tres De Mayo);Colombia
BGA;SKBG;Bucaramanga (Palonegro);Colombia
BOG;SKBO;Bogota (Eldorado Intl);Colombia
BAQ;SKBQ;Barranquilla (Ernesto Cortissoz);Colombia
BSC;SKBS;Bahia Solano (Jose Celestino Mutis);Colombia
BUN;SKBU;Buenaventura (Gerardo Tobar Lopez);Colombia
CUC;SKCC;Cucuta (Camilo Daza);Colombia
CTG;SKCG;Cartagena (Rafael Nunez);Colombia
CLO;SKCL;Cali (Alfonso Bonilla Aragon Intl);Colombia
TCO;SKCO;Tumaco (La Florida);Colombia
CZU;SKCZ;Corozal (Las Brujas);Colombia
EJA;SKEJ;Barrancabermeja (Yariguies);Colombia
FLA;SKFL;Florencia (Gustavo Artunduaga Paredes);Colombia
GPI;SKGP;Guapi (Juan Casiano);Colombia
IBE;SKIB;Ibague (Perales);Colombia
IPI;SKIP;Ipiales (San Luis);Colombia
LET;SKLT;Leticia (Alfredo Vasquez Cobo);Colombia
EOH;SKMD;Medellin (Olaya Herrera);Colombia
MGN;SKMG;Magangue (Baracoa);Colombia
MTR;SKMR;Monteria (Los Garzones);Colombia
MVP;SKMU;Mitu (Fabio Alberto Leon Bentley);Colombia
MZL;SKMZ;Manizales (La Nubia);Colombia
NVA;SKNV;Neiva (Benito Salas);Colombia
OCV;SKOC;Ocana (Aguas Claras);Colombia
OTU;SKOT;Otu;Colombia
PCR;SKPC;Puerto Carreno;Colombia
PEI;SKPE;Pereira (Matecana);Colombia
PPN;SKPP;Popayan (Guillermo Leon Valencia);Colombia
PSO;SKPS;Pasto (Antonio Narino);Colombia
PVA;SKPV;Providencia (El Embrujo);Colombia
MDE;SKRG;Rio Negro (Jose Maria Cordova);Colombia
RCH;SKRH;Rio Hacha (Almirante Padilla);Colombia
SJE;SKSJ;San Jose Del Guaviare (Jorge E Gonzalez Torres);Colombia
SMR;SKSM;Santa Marta (Simon Bolivar);Colombia
ADZ;SKSP;San Andres Island (Gustavo Rojas Pinilla);Colombia
SVI;SKSV;San Vincente De Caguan (Eduardo Falla Solano);Colombia
TME;SKTM;Tame;Colombia
AUC;SKUC;Arauca (Santiago Perez);Colombia
UIB;SKUI;Quibdo (El Carano);Colombia
ULQ;SKUL;Tulua (Farfan);Colombia
VUP;SKVP;Valledupar (Alfonso Lopez Pumarejo);Colombia
VVC;SKVV;Villavicencio (Vanguardia);Colombia
BJO;SLBJ;Bermejo;Bolivia
CBB;SLCB;Cochabamba (Jorge Wilsterman);Bolivia
CIJ;SLCO;Cobija (Heroes Del Acre);Bolivia
LPB;SLLP;La Paz (El Alto Intl);Bolivia
POI;SLPO;Potosi (Capitan Nicolas Rojas);Bolivia
PSZ;SLPS;Puerto Suarez (Tte De Av Salvador Ogaya G);Bolivia
SRE;SLSU;Sucre (Juana Azurduy De Padilla);Bolivia
TJA;SLTJ;Tarija (Capitan Oriel Lea Plaza);Bolivia
TDD;SLTR;Trinidad (Tte Av Jorge Henrich Arauz);Bolivia
VVI;SLVR;Santa Cruz (Viru Viru Intl);Bolivia
BYC;SLYA;Yacuiba;Bolivia
PBM;SMJP;Zandery (Johan A Pengel Intl);Suriname
CAY;SOCA;Cayenne (Rochambeau);French Guiana
PCL;SPCL;Pucallpa (Cap Fap David Abenzur Rengifo Intl);Peru
CHM;SPEO;Chimbote (Teniente Jaime A De Montreuil Morales);Peru
CIX;SPHI;Chiclayo (Capt Jose A Quinones Gonzales Intl);Peru
AYP;SPHO;Ayacucho (Coronel Fap Alfredo Mendivil Duarte);Peru
ANS;SPHY;Andahuaylas;Peru
ATA;SPHZ;Anta (Comandante Fap German Arias Graziani);Peru
LIM;SPIM;Lima (Jorge Chavez Intl);Peru
JJI;SPJI;Juanjui;Peru
JUL;SPJL;Juliaca;Peru
MGC;KMGC;Michigan City (Michigan City Municipal Airport);United States
TBP;SPME;Tumbes (Pedro Canga);Peru
YMS;SPMS;Yurimaguas (Moises Benzaquen Rengifo);Peru
CHH;SPPY;Chachapoyas;Peru
IQT;SPQT;Iquitos (Coronel Francisco Secada Vignetta Intl);Peru
AQP;SPQU;Arequipa (Rodriguez Ballon);Peru
TRU;SPRU;Trujillo (Capitan Carlos Martinez De Pinillos);Peru
PIO;SPSO;Pisco (Pisco Intl);Peru
TPP;SPST;Tarapoto (Cadete Guillermo Del Castillo Paredes);Peru
TCQ;SPTN;Tacna (Coronel Carlos Ciriani Santa Rosa Intl);Peru
PEM;SPTU;Puerto Maldonado (Padre Aldamiz);Peru
PIU;SPUR;Piura (Capitan Fap Guillermo Concha Iberico);Peru
TYL;SPYL;Talara (Capitan Montes);Peru
CUZ;SPZO;Cuzco (Teniente Alejandro Velasco Astete Intl);Peru
MVD;SUMU;Montevideo (Carrasco Intl);Uruguay
STY;SUSO;Salto (Nueva Hesperides Intl);Uruguay
AGV;SVAC;Acarigua (Oswaldo Guevara Mujica);Venezuela
AAO;SVAN;Anaco;Venezuela
BLA;SVBC;Barcelona (General Jose Antonio Anzoategui Intl);Venezuela
BNS;SVBI;Barinas;Venezuela
BRM;SVBM;Barquisimeto (Barquisimeto Intl);Venezuela
CBL;SVCB;Ciudad Bolivar;Venezuela
CAJ;SVCN;Canaima;Venezuela
CUP;SVCP;Carupano (General Jose Francisco Bermudez);Venezuela
CZE;SVCR;Coro (Jose Leonardo Chirinos);Venezuela
CUM;SVCU;Cumana (Antonio Jose De Sucre);Venezuela
GUI;SVGI;Guiria;Venezuela
GUQ;SVGU;Guanare;Venezuela
LSP;SVJC;Paraguana (Josefa Camejo);Venezuela
LFR;SVLF;La Fria;Venezuela
MAR;SVMC;Maracaibo (La Chinita Intl);Venezuela
MRD;SVMD;Merida (Alberto Carnevalli);Venezuela
PMV;SVMG;Porlamar (Del Caribe Intl Gen Santiago Marino);Venezuela
CCS;SVMI;Caracas (Simon Bolivar Intl);Venezuela
MUN;SVMT;Maturin;Venezuela
PYH;SVPA;Puerto Ayacucho (Casique Aramare);Venezuela
PBL;SVPC;Puerto Cabello (General Bartolome Salom Intl);Venezuela
PZO;SVPR;Guayana (General Manuel Carlos Piar);Venezuela
SVZ;SVSA;San Antonio (San Antonio Del Tachira);Venezuela
STD;SVSO;Santo Domingo (Mayor Buenaventura Vivas);Venezuela
SFH;SVSP;San Felipe (Sub Teniente Nestor Arias);Venezuela
SFD;SVSR;San Fernando De Apure;Venezuela
SOM;SVST;San Tome;Venezuela
STB;SVSZ;Santa Barbara (Santa Barbara Del Zulia);Venezuela
TUV;SVTC;Tucupita;Venezuela
VLN;SVVA;Valencia (Arturo Michelena Intl);Venezuela
VLV;SVVL;Valera (Dr Antonio Nicolas Briceno);Venezuela
VDP;SVVP;Valle De La Pascua;Venezuela
LTM;SYLT;Lethem;Guyana
ANU;TAPA;Antigua (V C Bird Intl);Antigua and Barbuda
BGI;TBPB;Bridgetown (Grantley Adams Intl);Barbados
DCF;TDCF;Canefield;Dominica
DOM;TDPD;Dominica (Melville Hall);Dominica
FDF;TFFF;Fort-de-france (Le Lamentin);Martinique
SFG;TFFG;St. Martin (Grand Case);Guadeloupe
PTP;TFFR;Pointe-a-pitre (Le Raizet);Guadeloupe
GND;TGPY;Point Salines (Point Salines Intl);Grenada
STT;TIST;St. Thomas (Cyril E King);Virgin Islands
STX;TISX;St. Croix Island (Henry E Rohlsen);Virgin Islands
BQN;TJBQ;Aguadilla (Rafael Hernandez);Puerto Rico
FAJ;TJFA;Fajardo (Diego Jimenez Torres);Puerto Rico
SIG;TJIG;San Juan (Fernando Luis Ribas Dominicci);Puerto Rico
MAZ;TJMZ;Mayaguez (Eugenio Maria De Hostos);Puerto Rico
PSE;TJPS;Ponce (Mercedita);Puerto Rico
SJU;TJSJ;San Juan (Luis Munoz Marin Intl);Puerto Rico
SKB;TKPK;Basse Terre (Robert L Bradshaw);Saint Kitts and Nevis
SLU;TLPC;Castries (George F L Charles);Saint Lucia
UVF;TLPL;Hewandorra (Hewanorra Intl);Saint Lucia
AUA;TNCA;Oranjestad (Reina Beatrix Intl);Aruba
BON;TNCB;Kralendijk (Flamingo);Netherlands Antilles
CUR;TNCC;Willemstad (Hato);Netherlands Antilles
EUX;TNCE;Oranjestad (F D Roosevelt);Netherlands Antilles
SXM;TNCM;Philipsburg (Princess Juliana Intl);Netherlands Antilles
AXA;TQPF;The Valley (Wallblake);Anguilla
TAB;TTCP;Scarborough (Crown Point);Trinidad and Tobago
POS;TTPP;Port-of-spain (Piarco);Trinidad and Tobago
EIS;TUPJ;Tortola (Terrance B Lettsome Intl);British Virgin Islands
AET;PFAL;Allakaket (Allakaket Airport);United States
CIW;TVSC;Canouan Island (Canouan);Saint Vincent and the Grenadines
MQS;TVSM;Mustique;Saint Vincent and the Grenadines
SVD;TVSV;Kingstown (E T Joshua);Saint Vincent and the Grenadines
ALA;UAAA;Alma-ata (Almaty);Kazakhstan
TSE;UACC;Tselinograd (Astana Intl);Kazakhstan
DMB;UADD;Dzhambul (Taraz);Kazakhstan
FRU;UAFM;Bishkek (Manas);Kyrgyzstan
OSS;UAFO;Osh;Kyrgyzstan
CIT;UAII;Chimkent (Shymkent);Kazakhstan
YAK;PAYA;Yakutat;United States
URA;UARR;Uralsk;Kazakhstan
PWQ;UASP;Pavlodar;Kazakhstan
PLX;UASS;Semiplatinsk (Semipalatinsk);Kazakhstan
AKX;UATT;Aktyubinsk;Kazakhstan
GYD;UBBB;Baku (Heydar Aliyev);Azerbaijan
YKS;UEEE;Yakutsk;Russia
MJZ;UERR;Mirnyj (Mirny);Russia
BQS;UHBB;Blagoveschensk (Ignatyevo);Russia
KHV;UHHH;Khabarovsk (Novy);Russia
MQT;KMQT;Marquette (Sawyer International Airport);United States
PVS;UHMD;Provideniya Bay;Russia
GDX;UHMM;Magadan (Sokol);Russia
PWE;UHMP;Pevek;Russia
PKC;UHPP;Petropavlovsk (Yelizovo);Russia
UUS;UHSS;Yuzhno-sakhalinsk (Khomutovo);Russia
VVO;UHWW;Vladivostok (Knevichi);Russia
HTA;UIAA;Chita (Kadala);Russia
BTK;UIBB;Bratsk;Russia
IKT;UIII;Irkutsk;Russia
UUD;UIUU;Ulan-ude (Mukhino);Russia
KBP;UKBB;Kiev (Boryspil Intl);Ukraine
DOK;UKCC;Donetsk (Donetsk Intl);Ukraine
DNK;UKDD;Dnepropetrovsk (Dnipropetrovsk Intl);Ukraine
SIP;UKFF;Simferopol (Simferopol Intl);Ukraine
IEV;UKKK;Kiev (Zhuliany Intl);Ukraine
LWO;UKLL;Lvov (Lviv Intl);Ukraine
IMT;KIMT;Iron Mountain (Ford Airport);United States
ODS;UKOO;Odessa (Odesa Intl);Ukraine
LED;ULLI;St. Petersburg (Pulkovo);Russia
MMK;ULMM;Murmansk;Russia
GME;UMGG;Gomel;Belarus
VTB;UMII;Vitebsk;Belarus
KGD;UMKK;Kaliningrad (Khrabrovo);Russia
MHP;UMMM;Minsk (Minsk 1);Belarus
MSQ;UMMS;Minsk 2;Belarus
ABA;UNAA;Abakan;Russia
BAX;UNBB;Barnaul;Russia
KEJ;UNEE;Kemorovo (Kemerovo);Russia
OMS;UNOO;Omsk;Russia
KRR;URKK;Krasnodar (Pashkovskiy);Russia
MCX;URML;Makhachkala (Uytash);Russia
MRV;URMM;Mineralnye Vody (Mineralnyye Vody);Russia
STW;URMT;Stavropol (Shpakovskoye);Russia
ROV;URRR;Rostov (Rostov Na Donu);Russia
AER;URSS;Sochi;Russia
ASF;URWA;Astrakhan;Russia
VOG;URWW;Volgograd (Gumrak);Russia
CEK;USCC;Chelyabinsk (Balandino);Russia
MQF;USCM;Magnetiogorsk (Magnitogorsk);Russia
GBZ;NZGB;Claris (Great Barrier Island);New Zealand
NJC;USNN;Nizhnevartovsk;Russia
PEE;USPP;Perm (Bolshoye Savino);Russia
SGC;USRR;Surgut;Russia
SVX;USSS;Yekaterinburg (Koltsovo);Russia
ASB;UTAA;Ashkhabad (Ashgabat);Turkmenistan
KRW;UTAK;Krasnovodsk (Turkmenbashi);Turkmenistan
DYU;UTDD;Dushanbe;Tajikistan
BHK;UTSB;Bukhara;Uzbekistan
SKD;UTSS;Samarkand;Uzbekistan
IUD;OTBH;Doha (Al Udeid AB);Qatar
TAS;UTTT;Tashkent (Yuzhny);Uzbekistan
BZK;UUBP;Bryansk;Russia
SVO;UUEE;Moscow (Sheremetyevo);Russia
KLD;UUEM;Tver (Migalovo);Russia
VOZ;UUOO;Voronezh (Chertovitskoye);Russia
VKO;UUWW;Moscow (Vnukovo);Russia
SCW;UUYY;Syktyvkar;Russia
KZN;UWKD;Kazan;Russia
REN;UWOO;Orenburg;Russia
UFA;UWUU;Ufa;Russia
KBY;UWWW;Samara (Kurumoch);Russia
AMD;VAAH;Ahmedabad;India
AKD;VAAK;Akola;India
IXU;VAAU;Aurangabad;India
BOM;VABB;Mumbai (Chhatrapati Shivaji Intl);India
PAB;VABI;Bilaspur;India
BHJ;VABJ;Bhuj;India
IXG;VABM;Belgaum;India
BDQ;VABO;Baroda (Vadodara);India
BHO;VABP;Bhopal;India
BHU;VABV;Bhaunagar (Bhavnagar);India
NMB;VADN;Daman;India
GOI;VAGO;Goa;India
IDR;VAID;Indore (Devi Ahilyabai Holkar);India
JLR;VAJB;Jabalpur;India
JGA;VAJM;Jamnagar;India
IXY;VAKE;Kandla;India
HJR;VAKJ;Khajuraho;India
KLH;VAKP;Kolhapur;India
IXK;VAKS;Keshod;India
NAG;VANP;Nagpur (Dr Ambedkar Intl);India
ISK;VANR;Nasik Road;India
PNQ;VAPO;Pune;India
PBD;VAPR;Porbandar;India
RAJ;VARK;Rajkot;India
RPR;VARP;Raipur;India
SSE;VASL;Sholapur;India
STV;VASU;Surat;India
UDR;VAUD;Udaipur;India
CMB;VCBI;Colombo (Bandaranaike Intl Colombo);Sri Lanka
RML;VCCC;Colombo (Colombo Ratmalana);Sri Lanka
GOY;VCCG;Galoya (Amparai);Sri Lanka
JAF;VCCJ;Jaffna (Kankesanturai);Sri Lanka
TRR;VCCT;Trinciomalee (China Bay);Sri Lanka
KIK;ORKK;Kirkuk (Kirkuk AB);Iraq
PNH;VDPP;Phnom-penh (Phnom Penh Intl);Cambodia
REP;VDSR;Siem-reap (Siem Reap);Cambodia
IXA;VEAT;Agartala;India
AJL;VEAZ;Aizwal (Aizawl);India
IXB;VEBD;Baghdogra (Bagdogra);India
BBI;VEBS;Bhubaneswar (Bhubaneshwar);India
CCU;VECC;Kolkata (Netaji Subhash Chandra Bose Intl);India
COH;VECO;Cooch-behar (Cooch Behar);India
DBD;VEDB;Dhanbad;India
ESC;KESC;Escanaba (Delta County Airport);United States
GAY;VEGY;Gaya;India
IMF;VEIM;Imphal;India
IXW;VEJS;Jamshedpur;India
JRH;VEJT;Jorhat;India
IXH;VEKR;Kailashahar;India
IXS;VEKU;Silchar;India
IXI;VELR;Lilabari;India
MOH;VEMN;Mohanbari (Dibrugarh);India
PAT;VEPT;Patina (Patna);India
IXR;VERC;Ranchi (Birsa Munda);India
RRK;VERK;Rourkela;India
VTZ;VEVZ;Vishakhapatnam;India
CXB;VGCB;Cox's Bazar (Coxs Bazar);Bangladesh
CGP;VGEG;Chittagong (Shah Amanat Intl);Bangladesh
IRD;VGIS;Ishurdi;Bangladesh
JSR;VGJR;Jessore;Bangladesh
RJH;VGRJ;Rajshahi (Shah Mokhdum);Bangladesh
SPD;VGSD;Saidpur;Bangladesh
ZYL;VGSY;Sylhet Osmani (Osmany Intl);Bangladesh
DAC;VGZR;Dhaka (Zia Intl);Bangladesh
HKG;VHHH;Hong Kong (Hong Kong Intl);Hong Kong
AGR;VIAG;Agra;India
IXD;VIAL;Allahabad;India
ATQ;VIAR;Amritsar;India
VNS;VIBN;Varanasi;India
KUU;VIBR;Kulu (Kullu Manali);India
IXC;VICG;Chandigarh;India
DED;VIDN;Dehra Dun (Dehradun);India
DEL;VIDP;Delhi (Indira Gandhi Intl);India
GWL;VIGR;Gwalior;India
JDH;VIJO;Jodhpur;India
JAI;VIJP;Jaipur;India
JSA;VIJR;Jaisalmer;India
IXJ;VIJU;Jammu;India
KNU;VIKA;Kanpur;India
KTU;VIKO;Kota;India
LUH;VILD;Ludhiaha (Ludhiana);India
IXL;VILH;Leh;India
LKO;VILK;Lucknow;India
IXP;VIPK;Pathankot;India
PGH;VIPT;Nainital (Pantnagar);India
SXR;VISR;Srinagar;India
TNI;VIST;Satna;India
LPQ;VLLB;Luang Prabang (Luang Phabang Intl);Laos
PKZ;VLPS;Pakse;Laos
ZVK;VLSK;Savannakhet;Laos
VTE;VLVT;Vientiane (Wattay Intl);Laos
MFM;VMMC;Macau (Macau Intl);Macau
BWA;VNBW;Bhairawa (Bhairahawa);Nepal
KTM;VNKT;Kathmandu (Tribhuvan Intl);Nepal
PKR;VNPK;Pokhara;Nepal
SIF;VNSI;Simara;Nepal
BIR;VNVT;Biratnagar;Nepal
AGX;VOAT;Agatti Island (Agatti);India
BLR;VOBL;Bangalore;India
BEP;VOBI;Bellary;India
VGA;VOBZ;Vijayawada;India
CJB;VOCB;Coimbatore;India
COK;VOCI;Kochi (Cochin);India
CCJ;VOCL;Calicut;India
CDP;VOCP;Cuddapah;India
HYD;VOHY;Hyderabad;India
IXM;VOMD;Madurai;India
IXE;VOML;Mangalore;India
MAA;VOMM;Madras (Chennai Intl);India
IXZ;VOPB;Port Blair;India
PNY;VOPC;Pendicherry (Pondicherry);India
RJA;VORY;Rajahmundry;India
TIR;VOTP;Tirupeti (Tirupati);India
TRZ;VOTR;Tiruchirappalli (Trichy);India
TRV;VOTV;Trivandrum (Thiruvananthapuram Intl);India
PBH;VQPR;Thimphu (Paro);Bhutan
MLE;VRMM;Male (Male Intl);Maldives
DMK;VTBD;Bangkok (Don Muang Intl);Thailand
NAH;WAMH;Naha;Indonesia
UTP;VTBU;Pattaya (U Taphao Intl);Thailand
LPT;VTCL;Lampang;Thailand
PRH;VTCP;Phrae;Thailand
HHQ;VTPH;Prachuap Khiri Khan (Hua Hin);Thailand
PHS;VTPP;Phitsanulok;Thailand
NAW;VTSC;Narathiwat;Thailand
KBV;VTSG;Krabi;Thailand
PAN;VTSK;Pattani;Thailand
USM;VTSM;Ko Samui (Samui);Thailand
HKT;VTSP;Phuket (Phuket Intl);Thailand
HDY;VTSS;Hat Yai (Hat Yai Intl);Thailand
TST;VTST;Trang;Thailand
UTH;VTUD;Udon Thani;Thailand
SNO;VTUI;Sakon Nakhon;Thailand
LOE;VTUL;Loei;Thailand
OKB;KOKB;Fraser Island (Orchid Beach);Australia
DAD;VVDN;Danang (Danang Intl);Vietnam
HAN;VVNB;Hanoi (Noibai Intl);Vietnam
NHA;VVNT;Nhatrang;Vietnam
SGN;VVTS;Ho Chi Minh City (Tansonnhat Intl);Vietnam
BMQ;KBMQ;Bamburi;Kenya
HEH;VYHH;Heho;Burma
KET;VYKG;Kengtung;Burma
MWA;KMWA;Marion (Williamson Country Regional Airport);United States
KYP;VYKP;Kyaukpyu;Burma
LSH;VYLS;Lashio;Burma
MDL;VYMD;Mandalay (Mandalay Intl);Burma
MGZ;VYME;Myeik;Burma
MYT;VYMK;Myitkyina;Burma
MOG;VYMS;Mong Hsat;Burma
PBU;VYPT;Putao;Burma
AKY;VYSW;Sittwe;Burma
SNW;VYTD;Thandwe;Burma
THL;VYTL;Tachilek (Tachileik);Burma
RGN;VYYY;Yangon (Yangon Intl);Burma
UPG;WAAA;Ujung Pandang (Hasanuddin);Indonesia
BIK;WABB;Biak (Frans Kaisiepo);Indonesia
NBX;WABI;Nabire;Indonesia
TIM;WABP;Timika (Moses Kilangin);Indonesia
DJJ;WAJJ;Jayapura (Sentani);Indonesia
WMX;WAJW;Wamena;Indonesia
MKQ;WAKK;Merauke (Mopah);Indonesia
GTO;WAMG;Gorontalo (Jalaluddin);Indonesia
ICN;RKSI;Seoul (Incheon Intl);South Korea
PLW;WAML;Palu (Mutiara);Indonesia
MDC;WAMM;Manado (Sam Ratulangi);Indonesia
PSJ;WAMP;Poso (Kasiguncu);Indonesia
OTI;WAMR;Morotai Island (Pitu);Indonesia
TTE;WAMT;Ternate (Sultan Babullah);Indonesia
LUW;WAMW;Luwuk (Bubung);Indonesia
AMQ;WAPP;Ambon (Pattimura);Indonesia
FKQ;WASF;Fak Fak;Indonesia
KNG;WASK;Kaimana;Indonesia
BXB;WASO;Babo;Indonesia
MKW;WASR;Manokwari (Rendani);Indonesia
SOQ;WASS;Sorong (Jefman);Indonesia
BTU;WBGB;Bintulu;Malaysia
KCH;WBGG;Kuching (Kuching Intl);Malaysia
LMN;WBGJ;Limbang;Malaysia
MUR;WBGM;Marudi;Malaysia
MYY;WBGR;Miri;Malaysia
SBW;WBGS;Sibu;Malaysia
LDU;WBKD;Lahad Datu;Malaysia
BKI;WBKK;Kota Kinabalu (Kota Kinabalu Intl);Malaysia
LBU;WBKL;Labuan;Malaysia
TWU;WBKW;Tawau;Malaysia
BWN;WBSB;Bandar Seri Begawan (Brunei Intl);Brunei
PKU;WIBB;Pekanbaru (Sultan Syarif Kasim Ii);Indonesia
DUM;WIBD;Dumai (Pinang Kampai);Indonesia
CGK;WIII;Jakarta (Soekarno Hatta Intl);Indonesia
GNS;WIMB;Gunung Sitoli (Binaka);Indonesia
PDG;WIPT;Padang (Minangkabau);Indonesia
MES;WIMM;Medan (Polonia);Indonesia
KTG;WIOK;Ketapang (Rahadi Usman);Indonesia
PNK;WIOO;Pontianak (Supadio);Indonesia
DJB;WIPA;Jambi (Sultan Thaha);Indonesia
BKS;WIPL;Bengkulu (Fatmawati Soekarno);Indonesia
PLM;WIPP;Palembang (Sultan Mahmud Badaruddin Ii);Indonesia
RGT;WIPR;Rengat (Japura);Indonesia
BTJ;WITT;Banda Aceh (Sultan Iskandarmuda);Indonesia
AOR;WMKA;Alor Setar (Sultan Abdul Halim);Malaysia
KBR;WMKC;Kota Bahru (Sultan Ismail Petra);Malaysia
KUA;WMKD;Kuantan;Malaysia
KTE;WMKE;Kerteh;Malaysia
IPH;WMKI;Ipoh (Sultan Azlan Shah);Malaysia
JHB;WMKJ;Johor Bahru (Sultan Ismail);Malaysia
KUL;WMKK;Kuala Lumpur (Kuala Lumpur Intl);Malaysia
LGK;WMKL;Pulau (Langkawi Intl);Malaysia
MKZ;WMKM;Malacca;Malaysia
TGG;WMKN;Kuala Terengganu (Sultan Mahmud);Malaysia
PEN;WMKP;Penang (Penang Intl);Malaysia
DIL;WPDL;Dili (Presidente Nicolau Lobato Intl);East Timor
QPG;WSAP;Paya Lebar;Singapore
XSP;WSSL;Singapore (Seletar);Singapore
SIN;WSSS;Singapore (Changi Intl);Singapore
ABM;YBAM;Amberley (Bamaga Injinoo);Australia
ASP;YBAS;Alice Springs;Australia
BNE;YBBN;Brisbane (Brisbane Intl);Australia
OOL;YBCG;Coolangatta (Gold Coast);Australia
CNS;YBCS;Cairns (Cairns Intl);Australia
CTL;YBCV;Charlieville (Charleville);Australia
ISA;YBMA;Mount Isa;Australia
MCY;YBMC;Maroochydore (Sunshine Coast);Australia
MKY;YBMK;Mackay;Australia
PPP;YBPN;Prosserpine (Proserpine Whitsunday Coast);Australia
ROK;YBRK;Rockhampton;Australia
TSV;YBTL;Townsville;Australia
WEI;YBWP;Weipa;Australia
AVV;YMAV;Avalon;Australia
ABX;YMAY;Albury;Australia
MEB;YMEN;Melbourne (Melbourne Essendon);Australia
HBA;YMHB;Hobart;Australia
LST;YMLT;Launceston;Australia
MBW;YMMB;Melbourne (Melbourne Moorabbin);Australia
MEL;YMML;Melbourne (Melbourne Intl);Australia
ADL;YPAD;Adelaide (Adelaide Intl);Australia
JAD;YPJT;Perth (Perth Jandakot);Australia
KTA;YPKA;Karratha;Australia
KGI;YPKG;Kalgoorlie (Kalgoorlie Boulder);Australia
KNX;YPKU;Kununurra;Australia
LEA;YPLM;Learmonth;Australia
PHE;YPPD;Port Hedland (Port Hedland Intl);Australia
PER;YPPH;Perth (Perth Intl);Australia
UMR;YPWR;Woomera;Australia
XCH;YPXM;Christmas Island;Christmas Island
BWU;YSBK;Sydney (Sydney Bankstown);Australia
CBR;YSCB;Canberra;Australia
CFS;YSCH;Coff's Harbour (Coffs Harbour);Australia
CDU;YSCN;Camden;Australia
DBO;YSDU;Dubbo;Australia
NLK;YSNF;Norfolk Island (Norfolk Island Intl);Norfolk Island
RCM;YSRI;Richmond;Australia
SYD;YSSY;Sydney (Sydney Intl);Australia
TMW;YSTW;Tamworth;Australia
WGA;YSWG;Wagga Wagga;Australia
PEK;ZBAA;Beijing (Capital Intl);China
HLD;ZBLA;Hailar (Dongshan);China
TSN;ZBTJ;Tianjin (Binhai);China
TYN;ZBYN;Taiyuan (Wusu);China
CAN;ZGGG;Guangzhou (Baiyun Intl);China
CSX;ZGHA;Changcha (Huanghua);China
KWL;ZGKL;Guilin (Liangjiang);China
NNG;ZGNN;Nanning (Wuxu);China
SZX;ZGSZ;Shenzhen (Baoan Intl);China
CGO;ZHCC;Zhengzhou (Xinzheng);China
WUH;ZHHH;Wuhan (Tianhe);China
FNJ;ZKPY;Pyongyang (Pyongyang Intl);Korea
ZGC;ZLLL;Lanzhou (Zhongchuan);China
XIY;ZLXY;Xi'an (Xianyang);China
ULN;ZMUB;Ulan Bator (Chinggis Khaan Intl);Mongolia
KMG;ZPPP;Kunming (Wujiaba);China
XMN;ZSAM;Xiamen (Gaoqi);China
KHN;ZSCN;Nanchang (Changbei Intl);China
FOC;ZSFZ;Fuzhou (Changle);China
HGH;ZSHC;Hangzhou (Xiaoshan);China
NGB;ZSNB;Ninbo (Lishe);China
NKG;ZSNJ;Nanjing (Lukou);China
HFE;ZSOF;Hefei (Luogang);China
TAO;ZSQD;Qingdao (Liuting);China
SHA;ZSSS;Shanghai (Hongqiao Intl);China
YNT;ZSYT;Yantai (Laishan);China
CKG;ZUCK;Chongqing (Jiangbei);China
KWE;ZUGY;Guiyang (Longdongbao);China
CTU;ZUUU;Chengdu (Shuangliu);China
XIC;ZUXC;Xichang (Qingshan);China
KHG;ZWSH;Kashi;China
HTN;ZWTN;Hotan;China
URC;ZWWW;Urumqi (Diwopu);China
HRB;ZYHB;Harbin (Taiping);China
HOJ;LOIH;Hohenems;Austria
DLC;ZYTL;Dalian (Zhoushuizi);China
PVG;ZSPD;Shanghai (Pudong);China
TOD;WMBT;Tioman (Pulau Tioman);Malaysia
SZB;WMSA;Kuala Lumpur (Subang-Sultan Abdul Aziz Shah Intl);Malaysia
NTQ;RJNW;Wajima (Noto);Japan
HBE;HEBA;Alexandria (Borg El Arab Intl);Egypt
BTI;PABA;Barter Island (Barter Island Lrrs);United States
K03;PAWT;Fort Wainwright (Wainwright As);United States
LUR;PALU;Cape Lisburne (Cape Lisburne Lrrs);United States
PIZ;PPIZ;Point Lay (Point Lay Lrrs);United States
ITO;PHTO;Hilo (Hilo Intl);United States
ORL;KORL;Orlando (Executive);United States
BTT;PABT;Bettles;United States
Z84;PACL;Clear Mews (Clear);United States
UTO;PAIM;Indian Mountains (Indian Mountain Lrrs);United States
FYU;PFYU;Fort Yukon;United States
SVW;PASV;Sparrevohn (Sparrevohn Lrrs);United States
FRN;PAFR;Fort Richardson (Bryant Ahp);United States
TLJ;PATL;Tatalina (Tatalina Lrrs);United States
CZF;PACZ;Cape Romanzof (Cape Romanzof Lrrs);United States
BED;KBED;Bedford (Laurence G Hanscom Fld);United States
SNP;PASN;St. Paul Island (St Paul Island);United States
EHM;PAEH;Cape Newenham (Cape Newenham Lrrs);United States
PBV;PAPB;Point Barrow (St George);United States
ILI;PAIL;Iliamna;United States
PTU;PAPM;Port Moller (Platinum);United States
BMX;PABM;Big Mountain (Big Mountain Afs);United States
OSC;KOSC;Oscoda (Oscoda Wurtsmith);United States
OAR;KOAR;Fort Ord (Marina Muni);United States
MHR;KMHR;Sacramento (Sacramento Mather);United States
BYS;KBYS;Fort Irwin (Bicycle Lake Aaf);United States
NXP;KNXP;Twenty Nine Palms (Twentynine Palms Eaf);United States
FSM;KFSM;Fort Smith (Fort Smith Rgnl);United States
MRI;PAMR;Anchorage (Merrill Fld);United States
GNT;KGNT;Grants (Grants Milan Muni);United States
PNC;KPNC;Ponca City (Ponca City Rgnl);United States
SVN;KSVN;Hunter Aaf;United States
GFK;KGFK;Grand Forks (Grand Forks Intl);United States
PBF;KPBF;Pine Bluff (Grider Fld);United States
NSE;KNSE;Milton (Whiting Fld Nas North);United States
HNM;PHHN;Hana;United States
PRC;KPRC;Prescott (Ernest A Love Fld);United States
TTN;KTTN;Trenton (Trenton Mercer);United States
BOS;KBOS;Boston (General Edward Lawrence Logan Intl);United States
SUU;KSUU;Fairfield (Travis Afb);United States
RME;KRME;Rome (Griffiss Afld);United States
ENV;KENV;Wendover;United States
BFM;KBFM;Mobile (Mobile Downtown);United States
OAK;KOAK;Oakland (Metropolitan Oakland Intl);United States
OMA;KOMA;Omaha (Eppley Afld);United States
NOW;KNOW;Port Angeles (Port Angeles Cgas);United States
OGG;PHOG;Kahului;United States
ICT;KICT;Wichita (Wichita Mid Continent);United States
MCI;KMCI;Kansas City (Kansas City Intl);United States
MSN;KMSN;Madison (Dane Co Rgnl Truax Fld);United States
DLG;PADL;Dillingham;United States
HRO;KHRO;Harrison (Boone Co);United States
PHX;KPHX;Phoenix (Phoenix Sky Harbor Intl);United States
BGR;KBGR;Bangor (Bangor Intl);United States
FXE;KFXE;Fort Lauderdale (Fort Lauderdale Executive);United States
GGG;KGGG;Longview (East Texas Rgnl);United States
AND;KAND;Andersen (Anderson Rgnl);United States
GEG;KGEG;Spokane (Spokane Intl);United States
HWO;KHWO;Hollywood (North Perry);United States
SFO;KSFO;San Francisco (San Francisco Intl);United States
CTB;KCTB;Cutbank (Cut Bank Muni);United States
ARA;KARA;Louisiana (Acadiana Rgnl);United States
GNV;KGNV;Gainesville (Gainesville Rgnl);United States
MEM;KMEM;Memphis (Memphis Intl);United States
DUG;KDUG;Douglas (Bisbee Douglas Intl);United States
BIG;PABI;Delta Junction (Allen Aaf);United States
CNW;KCNW;Waco (Tstc Waco);United States
ANN;PANT;Annette Island;United States
CAR;KCAR;Caribou (Caribou Muni);United States
LRF;KLRF;Jacksonville (Little Rock Afb);United States
HUA;KHUA;Redstone (Redstone Aaf);United States
POB;KPOB;Fort Bragg (Pope Field);United States
DHT;KDHT;Dalhart (Dalhart Muni);United States
DLF;KDLF;Del Rio (Laughlin Afb);United States
LAX;KLAX;Los Angeles (Los Angeles Intl);United States
ANB;KANB;Anniston (Anniston Metro);United States
CLE;KCLE;Cleveland (Cleveland Hopkins Intl);United States
DOV;KDOV;Dover (Dover Afb);United States
CVG;KCVG;Cincinnati (Cincinnati Northern Kentucky Intl);United States
FME;KFME;Fort Meade (Tipton);United States
NID;KNID;China (China Lake Naws);United States
HON;KHON;Huron (Huron Rgnl);United States
JNU;PAJN;Juneau (Juneau Intl);United States
LFT;KLFT;Lafayette (Lafayette Rgnl);United States
EWR;KEWR;Newark (Newark Liberty Intl);United States
BOI;KBOI;Boise (Boise Air Terminal);United States
INS;KINS;Indian Springs (Creech Afb);United States
GCK;KGCK;Garden City (Garden City Rgnl);United States
MOT;KMOT;Minot (Minot Intl);United States
HHI;PHHI;Wahiawa (Wheeler Aaf);United States
MXF;KMXF;Montgomery (Maxwell Afb);United States
RBM;KRBM;Robinson (Robinson Aaf);United States
DAL;KDAL;Dallas (Dallas Love Fld);United States
FCS;KFCS;Fort Carson (Butts Aaf);United States
HLN;KHLN;Helena (Helena Rgnl);United States
NKX;KNKX;Miramar (Miramar Mcas);United States
LUF;KLUF;Phoenix (Luke Afb);United States
HRT;KHRT;Mary Esther (Hurlburt Fld);United States
HHR;KHHR;Hawthorne (Jack Northrop Fld Hawthorne Muni);United States
HUL;KHUL;Houlton (Houlton Intl);United States
END;KEND;Enid (Vance Afb);United States
NTD;KNTD;Point Mugu (Point Mugu Nas);United States
EDW;KEDW;Edwards Afb;United States
LCH;KLCH;Lake Charles (Lake Charles Rgnl);United States
KOA;PHKO;Kona (Kona Intl At Keahole);United States
MYR;KMYR;Myrtle Beach (Myrtle Beach Intl);United States
NLC;KNLC;Lemoore (Lemoore Nas);United States
ACK;KACK;Nantucket (Nantucket Mem);United States
FAF;KFAF;Fort Eustis (Felker Aaf);United States
HOP;KHOP;Hopkinsville (Campbell Aaf);United States
DCA;KDCA;Washington (Ronald Reagan Washington Natl);United States
NHK;KNHK;Patuxent River (Patuxent River Nas);United States
PSX;KPSX;Palacios (Palacios Muni);United States
BYH;KBYH;Blytheville (Arkansas Intl);United States
ACY;KACY;Atlantic City (Atlantic City Intl);United States
TIK;KTIK;Oklahoma City (Tinker Afb);United States
ECG;KECG;Elizabeth City (Elizabeth City Cgas Rgnl);United States
PUB;KPUB;Pueblo (Pueblo Memorial);United States
PQI;KPQI;Presque Isle (Northern Maine Rgnl At Presque Isle);United States
IKR;KIKR;Kirtland A.f.b. (Kirtland Air Force Base);United States
GRF;KGRF;Fort Lewis (Gray Aaf);United States
ADQ;PADQ;Kodiak;United States
UPP;PHUP;Opolu (Upolu);United States
FLL;KFLL;Fort Lauderdale (Fort Lauderdale Hollywood Intl);United States
MKO;KMKO;Muskogee (Davis Fld);United States
INL;KINL;International Falls (Falls Intl);United States
SLC;KSLC;Salt Lake City (Salt Lake City Intl);United States
CDS;KCDS;Childress (Childress Muni);United States
BIX;KBIX;Biloxi (Keesler Afb);United States
LSF;KLSF;Fort Benning (Lawson Aaf);United States
NQI;KNQI;Kingsville (Kingsville Nas);United States
FRI;KFRI;Fort Riley (Marshall Aaf);United States
MDT;KMDT;Harrisburg (Harrisburg Intl);United States
LNK;KLNK;Lincoln;United States
LAN;KLAN;Lansing (Capital City);United States
MUE;PHMU;Kamuela (Waimea Kohala);United States
MSS;KMSS;Massena (Massena Intl Richards Fld);United States
HKY;KHKY;Hickory (Hickory Rgnl);United States
SPG;KSPG;St. Petersburg (Albert Whitted);United States
FMY;KFMY;Fort Myers (Page Fld);United States
IAH;KIAH;Houston (George Bush Intercontinental);United States
MLT;KMLT;Millinocket (Millinocket Muni);United States
ADW;KADW;Camp Springs (Andrews Afb);United States
INT;KINT;Winston-salem (Smith Reynolds);United States
VCV;KVCV;Victorville (Southern California Logistics);United States
CEW;KCEW;Crestview (Bob Sikes);United States
GTB;KGTB;Fort Drum (Wheeler Sack Aaf);United States
PHN;KPHN;Port Huron (St Clair Co Intl);United States
BFL;KBFL;Bakersfield (Meadows Fld);United States
ELP;KELP;El Paso (El Paso Intl);United States
HRL;KHRL;Harlingen (Valley Intl);United States
CAE;KCAE;Columbia (Columbia Metropolitan);United States
DMA;KDMA;Tucson (Davis Monthan Afb);United States
NPA;KNPA;Pensacola (Pensacola Nas);United States
PNS;KPNS;Pensacola (Pensacola Rgnl);United States
RDR;KRDR;Red River (Grand Forks Afb);United States
HOU;KHOU;Houston (William P Hobby);United States
BKF;KBKF;Buckley (Buckley Afb);United States
ORT;PAOR;Northway;United States
PAQ;PAAQ;Palmer (Palmer Muni);United States
PIT;KPIT;Pittsburgh (Pittsburgh Intl);United States
BRW;PABR;Barrow (Wiley Post Will Rogers Mem);United States
EFD;KEFD;Houston (Ellington Fld);United States
NUW;KNUW;Whidbey Island (Whidbey Island Nas);United States
ALI;KALI;Alice (Alice Intl);United States
VAD;KVAD;Valdosta (Moody Afb);United States
MIA;KMIA;Miami (Miami Intl);United States
SEA;KSEA;Seattle (Seattle Tacoma Intl);United States
CHA;KCHA;Chattanooga (Lovell Fld);United States
BDR;KBDR;Stratford (Igor I Sikorsky Mem);United States
JAN;KJAN;Jackson (Jackson Evers Intl);United States
GLS;KGLS;Galveston (Scholes Intl At Galveston);United States
LGB;KLGB;Long Beach;United States
HDH;PHDH;Dillingham;United States
IPT;KIPT;Williamsport (Williamsport Rgnl);United States
IND;KIND;Indianapolis (Indianapolis Intl);United States
SZL;KSZL;Knobnoster (Whiteman Afb);United States
AKC;KAKR;Akron (Akron Fulton Intl);United States
GWO;KGWO;Greenwood (Greenwood Leflore);United States
HPN;KHPN;White Plains (Westchester Co);United States
FOK;KFOK;West Hampton Beach (Francis S Gabreski);United States
JBR;KJBR;Jonesboro (Jonesboro Muni);United States
TNX;KTNX;Tonopah (Tonopah Test Range);United States
LNA;KLNA;West Palm Beach (Palm Beach Co Park);United States
NZY;KNZY;San Diego (North Island Nas);United States
BIF;KBIF;El Paso (Biggs Aaf);United States
YUM;KYUM;Yuma (Yuma Mcas Yuma Intl);United States
CNM;KCNM;Carlsbad (Cavern City Air Terminal);United States
DLH;KDLH;Duluth (Duluth Intl);United States
BET;PABE;Bethel;United States
LOU;KLOU;Louisville (Bowman Fld);United States
FHU;KFHU;Fort Huachuca (Sierra Vista Muni Libby Aaf);United States
LIH;PHLI;Lihue;United States
HUF;KHUF;Terre Haute (Terre Haute Intl Hulman Fld);United States
HVR;KHVR;Havre (Havre City Co);United States
MWH;KMWH;Grant County Airport (Grant Co Intl);United States
MPV;KMPV;Montpelier (Edward F Knapp State);United States
RIC;KRIC;Richmond (Richmond Intl);United States
SHV;KSHV;Shreveport (Shreveport Rgnl);United States
CDV;PACV;Cordova (Merle K Mudhole Smith);United States
ORF;KORF;Norfolk (Norfolk Intl);United States
BPT;KBPT;Beaumont (Southeast Texas Rgnl);United States
SAV;KSAV;Savannah (Savannah Hilton Head Intl);United States
HIF;KHIF;Ogden (Hill Afb);United States
OME;PAOM;Nome;United States
SPB;KSPB;San Luis (Scappoose Industrial Airpark);United States
PIE;KPIE;St. Petersburg (St Petersburg Clearwater Intl);United States
MNM;KMNM;Macon (Menominee Marinette Twin Co);United States
CXO;KCXO;Conroe (Lone Star Executive);United States
SCC;PASC;Deadhorse;United States
SAT;KSAT;San Antonio (San Antonio Intl);United States
ROC;KROC;Rochester (Greater Rochester Intl);United States
COF;KCOF;Coco Beach (Patrick Afb);United States
TEB;KTEB;Teterboro;United States
RCA;KRCA;Rapid City (Ellsworth Afb);United States
RDU;KRDU;Raleigh-durham (Raleigh Durham Intl);United States
DAY;KDAY;Dayton (James M Cox Dayton Intl);United States
ENA;PAEN;Kenai (Kenai Muni);United States
MLC;KMLC;Mcalester (Mc Alester Rgnl);United States
IAG;KIAG;Niagara Falls (Niagara Falls Intl);United States
CFD;KCFD;Bryan (Coulter Fld);United States
PHF;KPHF;Newport News (Newport News Williamsburg Intl);United States
ESF;KESF;Alexandria (Esler Rgnl);United States
LTS;KLTS;Altus (Altus Afb);United States
TUS;KTUS;Tucson (Tucson Intl);United States
MIB;KMIB;Minot (Minot Afb);United States
BAB;KBAB;Marysville (Beale Afb);United States
IKK;KIKK;Kankakee (Greater Kankakee);United States
GSB;KGSB;Goldsboro (Seymour Johnson Afb);United States
PVD;KPVD;Providence (Theodore Francis Green State);United States
SBY;KSBY;Salisbury (Salisbury Ocean City Wicomico Rgnl);United States
RIU;KRIU;Rancho Murieta;United States
BUR;KBUR;Burbank (Bob Hope);United States
DTW;KDTW;Detroit (Detroit Metro Wayne Co);United States
TPA;KTPA;Tampa (Tampa Intl);United States
PMB;KPMB;Pembina (Pembina Muni);United States
POE;KPOE;Fort Polk (Polk Aaf);United States
EIL;PAEI;Fairbanks (Eielson Afb);United States
HIB;KHIB;Hibbing (Chisholm Hibbing);United States
LFK;KLFK;Lufkin (Angelina Co);United States
MAF;KMAF;Midland (Midland Intl);United States
GRB;KGRB;Green Bay (Austin Straubel Intl);United States
ADM;KADM;Ardmore (Ardmore Muni);United States
WRI;KWRI;Wrightstown (Mc Guire Afb);United States
NKT;KNKT;Cherry Point (Cherry Point Mcas);United States
SBO;KSBO;Santa Barbara (Emanuel Co);United States
AGS;KAGS;Bush Field (Augusta Rgnl At Bush Fld);United States
ISN;KISN;Williston (Sloulin Fld Intl);United States
LIT;KLIT;Little Rock (Adams Fld);United States
SWF;KSWF;Newburgh (Stewart Intl);United States
BDE;KBDE;Baudette (Baudette Intl);United States
SAC;KSAC;Sacramento (Sacramento Executive);United States
HOM;PAHO;Homer;United States
TBN;KTBN;Fort Leonardwood (Waynesville Rgnl Arpt At Forney Fld);United States
MGE;KMGE;Marietta (Dobbins Arb);United States
SKA;KSKA;Spokane (Fairchild Afb);United States
HTL;KHTL;Houghton Lake (Roscommon Co);United States
PAM;KPAM;Panama City (Tyndall Afb);United States
DFW;KDFW;Dallas-Fort Worth (Dallas Fort Worth Intl);United States
MLB;KMLB;Melbourne (Melbourne Intl);United States
TCM;KTCM;Tacoma (Mc Chord Afb);United States
AUS;KAUS;Austin (Austin Bergstrom Intl);United States
LCK;KLCK;Columbus (Rickenbacker Intl);United States
TYS;KTYS;Knoxville (Mc Ghee Tyson);United States
HLR;KHLR;Fort Hood (Hood Aaf);United States
STL;KSTL;St. Louis (Lambert St Louis Intl);United States
MIV;KMIV;Millville (Millville Muni);United States
SPS;KSPS;Wichita Falls (Sheppard Afb Wichita Falls Muni);United States
LUK;KLUK;Cincinnati (Cincinnati Muni Lunken Fld);United States
ATL;KATL;Atlanta (Hartsfield Jackson Atlanta Intl);United States
MER;KMER;Merced (Castle);United States
MCC;KMCC;Sacramento (Mc Clellan Afld);United States
GRR;KGRR;Grand Rapids (Gerald R Ford Intl);United States
INK;KINK;Wink (Winkler Co);United States
FAT;KFAT;Fresno (Fresno Yosemite Intl);United States
VRB;KVRB;Vero Beach (Vero Beach Muni);United States
IPL;KIPL;Imperial (Imperial Co);United States
BNA;KBNA;Nashville (Nashville Intl);United States
LRD;KLRD;Laredo (Laredo Intl);United States
EDF;PAED;Anchorage (Elmendorf Afb);United States
OTZ;PAOT;Kotzebue (Ralph Wien Mem);United States
AOO;KAOO;Altoona (Altoona Blair Co);United States
DYS;KDYS;Abilene (Dyess Afb);United States
ELD;KELD;El Dorado (South Arkansas Rgnl At Goodwin Fld);United States
LGA;KLGA;New York (La Guardia);United States
TLH;KTLH;Tallahassee (Tallahassee Rgnl);United States
DPA;KDPA;West Chicago (Dupage);United States
ACT;KACT;Waco (Waco Rgnl);United States
AUG;KAUG;Augusta (Augusta State);United States
INJ;KINJ;Hillsboro (Hillsboro Muni);United States
NIP;KNIP;Jacksonville (Jacksonville Nas);United States
MKL;KMKL;Jackson (Mc Kellar Sipes Rgnl);United States
MKK;PHMK;Molokai;United States
FTK;KFTK;Fort Knox (Godman Aaf);United States
SJT;KSJT;San Angelo (San Angelo Rgnl Mathis Fld);United States
CXL;KCXL;Calexico (Calexico Intl);United States
CIC;KCIC;Chico (Chico Muni);United States
BTV;KBTV;Burlington (Burlington Intl);United States
JAX;KJAX;Jacksonville (Jacksonville Intl);United States
DRO;KDRO;Durango (Durango La Plata Co);United States
IAD;KIAD;Washington (Washington Dulles Intl);United States
CLL;KCLL;College Station (Easterwood Fld);United States
SFF;KSFF;Spokane (Felts Fld);United States
MKE;KMKE;Milwaukee (General Mitchell Intl);United States
ABI;KABI;Abilene (Abilene Rgnl);United States
COU;KCOU;Columbia (Columbia Rgnl);United States
PDX;KPDX;Portland (Portland Intl);United States
TNT;KTNT;Miami (Dade Collier Training And Transition);United States
PBI;KPBI;West Palm Beach (Palm Beach Intl);United States
FTW;KFTW;Fort Worth (Fort Worth Meacham Intl);United States
OGS;KOGS;Ogdensburg (Ogdensburg Intl);United States
FMH;KFMH;Falmouth (Otis Angb);United States
BFI;KBFI;Seattle (Boeing Fld King Co Intl);United States
SKF;KSKF;San Antonio (Lackland Afb Kelly Fld Annex);United States
HNL;PHNL;Honolulu (Honolulu Intl);United States
DSM;KDSM;Des Moines (Des Moines Intl);United States
EWN;KEWN;New Bern (Craven Co Rgnl);United States
SAN;KSAN;San Diego (San Diego Intl);United States
MLU;KMLU;Monroe (Monroe Rgnl);United States
SSC;KSSC;Sumter (Shaw Afb);United States
ONT;KONT;Ontario (Ontario Intl);United States
GVT;KGVT;Greenvile (Majors);United States
ROW;KROW;Roswell (Roswell Intl Air Center);United States
DET;KDET;Detroit (Coleman A Young Muni);United States
BRO;KBRO;Brownsville (Brownsville South Padre Island Intl);United States
DHN;KDHN;Dothan (Dothan Rgnl);United States
WWD;KWWD;Wildwood (Cape May Co);United States
NFL;KNFL;Fallon (Fallon Nas);United States
MTC;KMTC;Mount Clemens (Selfridge Angb);United States
FMN;KFMN;Farmington (Four Corners Rgnl);United States
CRP;KCRP;Corpus Christi (Corpus Christi Intl);United States
SYR;KSYR;Syracuse (Syracuse Hancock Intl);United States
NQX;KNQX;Key West (Key West Nas);United States
MDW;KMDW;Chicago (Chicago Midway Intl);United States
SJC;KSJC;San Jose (Norman Y Mineta San Jose Intl);United States
HOB;KHOB;Hobbs (Lea Co Rgnl);United States
PNE;KPNE;Philadelphia (Northeast Philadelphia);United States
DEN;KDEN;Denver (Denver Intl);United States
PHL;KPHL;Philadelphia (Philadelphia Intl);United States
SUX;KSUX;Sioux City (Sioux Gateway Col Bud Day Fld);United States
MCN;KMCN;Macon (Middle Georgia Rgnl);United States
TCS;KTCS;Truth Or Consequences (Truth Or Consequences Muni);United States
PMD;KPMD;Palmdale (Palmdale Rgnl Usaf Plt 42);United States
RND;KRND;San Antonio (Randolph Afb);United States
NJK;KNJK;El Centro (El Centro Naf);United States
CMH;KCMH;Columbus (Port Columbus Intl);United States
FYV;KFYV;Fayetteville (Drake Fld);United States
FSI;KFSI;Fort Sill (Henry Post Aaf);United States
PNM;KPNM;Princeton (Princeton Muni);United States
FFO;KFFO;Dayton (Wright Patterson Afb);United States
GAL;PAGA;Galena (Edward G Pitka Sr);United States
MWL;KMWL;Mineral Wells;United States
IAB;KIAB;Wichita (Mc Connell Afb);United States
NBG;KNBG;New Orleans (New Orleans Nas Jrb);United States
TXK;KTXK;Texarkana (Texarkana Rgnl Webb Fld);United States
PBG;KPBG;Plattsburgh (Plattsburgh Intl);United States
APG;KAPG;Aberdeen (Phillips Aaf);United States
TCC;KTCC;Tucumcari (Tucumcari Muni);United States
ANC;PANC;Anchorage (Ted Stevens Anchorage Intl);United States
GRK;KGRK;Killeen (Robert Gray Aaf);United States
ZUN;KZUN;Zuni Pueblo (Black Rock);United States
BLI;KBLI;Bellingham (Bellingham Intl);United States
NQA;KNQA;Millington (Millington Rgnl Jetport);United States
EKN;KEKN;Elkins (Elkins Randolph Co Jennings Randolph);United States
HFD;KHFD;Hartford (Hartford Brainard);United States
SFZ;KSFZ;Smithfield (North Central State);United States
MOB;KMOB;Mobile (Mobile Rgnl);United States
NUQ;KNUQ;Mountain View (Moffett Federal Afld);United States
SAF;KSAF;Santa Fe (Santa Fe Muni);United States
BKH;PHBK;Barking Sands (Barking Sands Pmrf);United States
DRI;KDRI;Deridder (Beauregard Rgnl);United States
BSF;PHSF;Bradshaw Field (Bradshaw Aaf);United States
OLS;KOLS;Nogales (Nogales Intl);United States
MCF;KMCF;Tampa (Macdill Afb);United States
BLV;KBLV;Belleville (Scott Afb Midamerica);United States
OPF;KOPF;Miami (Opa Locka);United States
DRT;KDRT;Del Rio (Del Rio Intl);United States
RSW;KRSW;Fort Myers (Southwest Florida Intl);United States
AKN;PAKN;King Salmon;United States
MUI;KMUI;Muir (Muir Aaf);United States
JHM;PHJH;Lahania-kapalua (Kapalua);United States
JFK;KJFK;New York (John F Kennedy Intl);United States
HST;KHST;Homestead (Homestead Arb);United States
RAL;KRAL;Riverside (Riverside Muni);United States
FLV;KFLV;Fort Leavenworth (Sherman Aaf);United States
WAL;KWAL;Wallops Island (Wallops Flight Facility);United States
HMN;KHMN;Alamogordo (Holloman Afb);United States
NXX;KNXX;Willow Grove (Willow Grove Nas Jrb);United States
CYS;KCYS;Cheyenne (Cheyenne Rgnl Jerry Olson Fld);United States
SCK;KSCK;Stockton (Stockton Metropolitan);United States
CHS;KCHS;Charleston (Charleston Afb Intl);United States
RNO;KRNO;Reno (Reno Tahoe Intl);United States
KTN;PAKT;Ketchikan (Ketchikan Intl);United States
YIP;KYIP;Detroit (Willow Run);United States
VBG;KVBG;Lompoc (Vandenberg Afb);United States
BHM;KBHM;Birmingham (Birmingham Intl);United States
NEL;KNEL;Lakehurst (Lakehurst Naes);United States
SYA;PASY;Shemya (Eareckson As);United States
LSV;KLSV;Las Vegas (Nellis Afb);United States
RIV;KRIV;Riverside (March Arb);United States
MOD;KMOD;Modesto (Modesto City Co Harry Sham);United States
SMF;KSMF;Sacramento (Sacramento Intl);United States
UGN;KUGN;Chicago (Waukegan Rgnl);United States
COS;KCOS;Colorado Springs (City Of Colorado Springs Muni);United States
BUF;KBUF;Buffalo (Buffalo Niagara Intl);United States
SKY;KSKY;Sandusky (Griffing Sandusky);United States
PAE;KPAE;Everett (Snohomish Co);United States
MUO;KMUO;Mountain Home (Mountain Home Afb);United States
CDC;KCDC;Cedar City (Cedar City Rgnl);United States
BDL;KBDL;Windsor Locks (Bradley Intl);United States
MFE;KMFE;Mcallen (Mc Allen Miller Intl);United States
NGU;KNGU;Norfolk (Norfolk Ns);United States
CEF;KCEF;Chicopee Falls (Westover Arb Metropolitan);United States
LBB;KLBB;Lubbock (Lubbock Preston Smith Intl);United States
ORD;KORD;Chicago (Chicago Ohare Intl);United States
BCT;KBCT;Boca Raton;United States
FAI;PAFA;Fairbanks (Fairbanks Intl);United States
NYG;KNYG;Quantico (Quantico Mcaf);United States
CVS;KCVS;Clovis (Cannon Afb);United States
NGF;PHNG;Kaneohe Bay (Kaneohe Bay Mcaf);United States
OFF;KOFF;Omaha (Offutt Afb);United States
GKN;PAGK;Gulkana;United States
ART;KART;Watertown (Watertown Intl);United States
PSP;KPSP;Palm Springs (Palm Springs Intl);United States
AMA;KAMA;Amarillo (Rick Husband Amarillo Intl);United States
FOD;KFOD;Fort Dodge (Fort Dodge Rgnl);United States
BAD;KBAD;Shreveport (Barksdale Afb);United States
FOE;KFOE;Topeka (Forbes Fld);United States
COT;KCOT;Cotulla (Cotulla Lasalle Co);United States
ILM;KILM;Wilmington (Wilmington Intl);United States
BTR;KBTR;Baton Rouge (Baton Rouge Metro Ryan Fld);United States
NMM;KNMM;Meridian (Meridian Nas);United States
TYR;KTYR;Tyler (Tyler Pounds Rgnl);United States
BWI;KBWI;Baltimore (Baltimore Washington Intl);United States
HBR;KHBR;Hobart (Hobart Muni);United States
LNY;PHNY;Lanai;United States
AEX;KAEX;Alexandria (Alexandria Intl);United States
WSD;KWSD;White Sands (Condron Aaf);United States
CDB;PACD;Cold Bay;United States
TUL;KTUL;Tulsa (Tulsa Intl);United States
SIT;PASI;Sitka (Sitka Rocky Gutierrez);United States
ISP;KISP;Islip (Long Island Mac Arthur);United States
MSP;KMSP;Minneapolis (Minneapolis St Paul Intl);United States
ILG;KILG;Wilmington (New Castle);United States
DUT;PADU;Unalaska;United States
MSY;KMSY;New Orleans (Louis Armstrong New Orleans Intl);United States
PWM;KPWM;Portland (Portland Intl Jetport);United States
OKC;KOKC;Oklahoma City (Will Rogers World);United States
ALB;KALB;Albany (Albany Intl);United States
VDZ;PAVD;Valdez (Valdez Pioneer Fld);United States
LFI;KLFI;Hampton (Langley Afb);United States
SNA;KSNA;Santa Ana (John Wayne Arpt Orange Co);United States
CBM;KCBM;Colombus (Columbus Afb);United States
TMB;KTMB;Kendall-tamiami (Kendall Tamiami Executive);United States
NTU;KNTU;Oceana (Oceana Nas);United States
GUS;KGUS;Peru (Grissom Arb);United States
CPR;KCPR;Casper (Natrona Co Intl);United States
VPS;KVPS;Valparaiso (Eglin Afb);United States
SEM;KSEM;Selma (Craig Fld);United States
EYW;KEYW;Key West (Key West Intl);United States
CLT;KCLT;Charlotte (Charlotte Douglas Intl);United States
LAS;KLAS;Las Vegas (Mc Carran Intl);United States
MCO;KMCO;Orlando (Orlando Intl);United States
FLO;KFLO;Florence (Florence Rgnl);United States
GTF;KGTF;Great Falls (Great Falls Intl);United States
YNG;KYNG;Youngstown (Youngstown Warren Rgnl);United States
FBK;PAFB;Fort Wainwright (Ladd Aaf);United States
MMV;KMMV;Mackminnville (Mc Minnville Muni);United States
WRB;KWRB;Macon (Robins Afb);United States
BKK;VTBS;Bangkok (Suvarnabhumi Intl);Thailand
KDI;WAWW;Kendari (Wolter Monginsidi);Indonesia
SBG;WITB;Sabang (Maimun Saleh);Indonesia
MLG;WARA;Malang (Abdul Rachman Saleh);Indonesia
BDO;WICC;Bandung (Husein Sastranegara);Indonesia
CBN;WICD;Cirebon (Penggung);Indonesia
JOG;WARJ;Yogyakarta (Adi Sutjipto);Indonesia
CXP;WIHL;Cilacap (Tunggul Wulung);Indonesia
PCB;WIHP;Jakarta (Pondok Cabe);Indonesia
SRG;WARS;Semarang (Achmad Yani);Indonesia
BTH;WIDD;Batam (Hang Nadim);Indonesia
TJQ;WIOD;Tanjung Pandan (H As Hanandjoeddin);Indonesia
PGK;WIPK;Pangkal Pinang (Depati Amir);Indonesia
TNJ;WIDN;Tanjung Pinang (Kijang);Indonesia
SIQ;WIDS;Singkep (Dabo);Indonesia
BDJ;WAOO;Banjarmasin (Syamsudin Noor);Indonesia
PKN;WAOI;Pangkalan Bun (Iskandar);Indonesia
PKY;WAOP;Palangkaraya (Tjilik Riwut);Indonesia
MOF;WATC;Maumere (Wai Oti);Indonesia
ENE;WATE;Ende (H Hasan Aroeboesman);Indonesia
RTG;WATG;Ruteng (Satar Tacik);Indonesia
KOE;WATT;Kupang (El Tari);Indonesia
LBJ;WATO;Labuhan Bajo (Mutiara Ii);Indonesia
BPN;WALL;Balikpapan (Sepinggan);Indonesia
TRK;WALR;Taraken (Juwata);Indonesia
SRI;WALS;Samarinda (Temindung);Indonesia
AMI;WADA;Mataram (Selaparang);Indonesia
BMU;WADB;Bima (Muhammad Salahuddin);Indonesia
WGP;WADW;Waingapu (Mau Hau);Indonesia
SUB;WARR;Surabaya (Juanda);Indonesia
SOC;WARQ;Solo City (Adi Sumarmo Wiryokusumo);Indonesia
CNX;VTCC;Chiang Mai (Chiang Mai Intl);Thailand
CEI;VTCT;Chiang Rai (Chiang Rai Intl);Thailand
NST;VTSF;Nakhon Si Thammarat;Thailand
DPS;WADD;Denpasar (Bali Ngurah Rai);Indonesia
NAK;VTUQ;Nakhon Ratchasima;Thailand
KOP;VTUW;Nakhon Phanom;Thailand
UBP;VTUU;Ubon Ratchathani;Thailand
KKC;VTUK;Khon Kaen;Thailand
THS;VTPO;Sukhothai;Thailand
ATH;LGAV;Athens (Eleftherios Venizelos Intl);Greece
NGO;RJGG;Nagoya (Chubu Centrair Intl);Japan
UKB;RJBE;Kobe;Japan
PUW;KPUW;Pullman (Pullman-Moscow Rgnl);United States
LWS;KLWS;Lewiston (Lewiston Nez Perce Co);United States
ELM;KELM;Elmira (Elmira Corning Rgnl);United States
ITH;KITH;Ithaca (Ithaca Tompkins Rgnl);United States
MRY;KMRY;Monterey (Monterey Peninsula);United States
SBA;KSBA;Santa Barbara (Santa Barbara Muni);United States
DAB;KDAB;Daytona Beach (Daytona Beach Intl);United States
LPX;EVLA;Liepaja (Liepaja Intl);Latvia
RIX;EVRA;Riga (Riga Intl);Latvia
SQQ;EYSA;Siauliai (Siauliai Intl);Lithuania
HLJ;EYSB;Barysiai;Lithuania
KUN;EYKA;Kaunas (Kaunas Intl);Lithuania
PLQ;EYPA;Palanga (Palanga Intl);Lithuania
VNO;EYVI;Vilnius (Vilnius Intl);Lithuania
PNV;EYPP;Panevezys (Pajuostis);Lithuania
EVN;UDYZ;Yerevan (Zvartnots);Armenia
LWN;UDSG;Gyumri;Armenia
ASA;HHSB;Assab (Assab Intl);Eritrea
ASM;HHAS;Asmara (Asmara Intl);Eritrea
MSW;HHMS;Massawa (Massawa Intl);Eritrea
GZA;LVGZ;Gaza (Yasser Arafat Intl);Palestine
RIY;OYRN;Mukalla (Riyan);Yemen
BUS;UGSB;Batumi;Georgia
KUT;UGKO;Kutaisi (Kopitnari);Georgia
TBS;UGTB;Tbilisi;Georgia
TAI;OYTZ;Taiz (Taiz Intl);Yemen
HOD;OYHD;Hodeidah (Hodeidah Intl);Yemen
ADE;OYAA;Aden (Aden Intl);Yemen
AXK;OYAT;Ataq;Yemen
AAY;OYGD;Al Ghaidah Intl;Yemen
SAH;OYSN;Sanaa (Sanaa Intl);Yemen
FMM;EDJA;Memmingen (Allgau);Germany
BHN;OYBN;Beihan;Yemen
SCT;OYSQ;Socotra (Socotra Intl);Yemen
NAV;LTAZ;Nevsehir (Kapadokya);Turkey
EZE;SAEZ;Buenos Aires (Ministro Pistarini);Argentina
EBL;ORER;Erbil (Erbil Intl);Iraq
EMD;YEML;Emerald;Australia
KIX;RJBB;Osaka (Kansai);Japan
JRB;KJRB;New York (Wall Street Heliport);United States
TAG;RPVT;Tagbilaran;Philippines
JAV;BGJN;Ilulissat;Greenland
JCH;BGCH;Qasigiannguit;Greenland
JEG;BGEM;Aasiaat;Greenland
PMI;LEPA;Palma de Mallorca (Son Sant Joan);Spain
DRW;YPDN;Darwin (Darwin Intl);Australia
URT;VTSB;Surat Thani;Thailand
NYU;VYBR;Nyuang U (Bagan Intl);Burma
MPH;RPXE;Caticlan (Godofredo P);Philippines
TKA;PATK;Talkeetna;United States
GZM;LMMG;Gozo (Xewkija Heliport);Malta
HVN;KHVN;New Haven (Tweed-New Haven Airport);United States
AVL;KAVL;Asheville (Asheville Regional Airport);United States
GSO;KGSO;Greensboro (Piedmont Triad);United States
FSD;KFSD;Sioux Falls;United States
AYQ;YAYE;Uluru (Ayers Rock);Australia
MHT;KMHT;Manchester NH (Manchester Regional Airport);United States
APF;KAPF;Naples (Naples Muni);United States
RDN;WMPR;Redang;Malaysia
SDF;KSDF;Louisville (Louisville International Airport);United States
CHO;KCHO;Charlottesville VA (Charlottesville-Albemarle);United States
ROA;KROA;Roanoke VA (Roanoke Regional);United States
LEX;KLEX;Lexington KY (Blue Grass);United States
EVV;KEVV;Evansville (Evansville Regional);United States
ABQ;KABQ;Albuquerque (Albuquerque International Sunport);United States
BZN;KBZN;Bozeman (Gallatin Field);United States
BIL;KBIL;Billings (Billings Logan International Airport);United States
BTM;KBTM;Butte (Bert Mooney Airport);United States
TVC;KTVC;Traverse City (Cherry Capital Airport);United States
FRS;MGTK;Flores (Mundo Maya International);Guatemala
BHB;KBHB;Bar Harbor (Hancock County - Bar Harbor);United States
RKD;KRKD;Rockland (Knox County Regional Airport);United States
JAC;KJAC;Jacksn Hole (Jackson Hole Airport);United States
RFD;KRFD;Rockford (Chicago Rockford International Airport );United States
DME;UUDD;Moscow (Domododevo);Russia
SYX;ZJSY;Sanya (Phoenix International);China
MFN;NZMF;Milford Sound (Milford Sound Airport);New Zealand
TSS;NONE;New York (East 34th Street Heliport);United States
LJG;ZPLJ;Lijiang (Lijiang Airport);China
GSP;KGSP;Greenville (Greenville-Spartanburg International);United States
BMI;KBMI;Bloomington (Central Illinois Rgnl);United States
GPT;KGPT;Gulfport (Gulfport-Biloxi);United States
AZO;KAZO;Kalamazoo;United States
TOL;KTOL;Toledo;United States
FWA;KFWA;Fort Wayne;United States
DEC;KDEC;Decatur;United States
CID;KCID;Cedar Rapids;United States
LSE;KLSE;La Crosse (La Crosse Municipal);United States
CWA;KCWA;Wassau (Central Wisconsin);United States
PIA;KPIA;Peoria (Peoria Regional);United States
ATW;KATW;Appleton;United States
RST;KRST;Rochester;United States
CMI;KCMI;Champaign;United States
MHK;KMHK;Manhattan (Manhattan Reigonal);United States
KGC;YKSC;Kingscote (Kingscote Airport);Australia
HVB;YHBA;Hervey Bay (Hervey Bay Airport);Australia
DLU;ZPDL;Dali;China
SSH;HESH;Sharm El Sheikh (Sharm El Sheikh Intl);Egypt
FKL;KFKL;Franklin;United States
NBO;HKJK;Nairobi (Jomo Kenyatta International);Kenya
SEU;HTSN;Seronera;Tanzania
FTE;SAWC;El Calafate;Argentina
ARM;YARM;Armidale;Australia
GJT;KGJT;Grand Junction (Grand Junction Regional);United States
SGU;KSGU;Saint George (St George Muni);United States
DWH;KDWH;Houston (David Wayne Hooks Field);United States
S46;XS46;Port O'Connor (Port O'Connor Airfield);United States
SRQ;KSRQ;Sarasota (Sarasota Bradenton Intl);United States
VNY;KVNY;Van Nuys;United States
BDA;TXKF;Bermuda (Bermuda Intl);Bermuda
MLI;KMLI;Moline (Quad City Intl);United States
PFN;KPFN;Panama City (Panama City Bay Co Intl);United States
HIR;AGGH;Honiara (Honiara International);Solomon Islands
PPT;NTAA;Papeete (Faa'a International);French Polynesia
INU;ANYN;Nauru (Nauru Intl);Nauru
FUN;NGFU;Funafuti (Funafuti International);Tuvalu
OVB;UNNT;Novosibirsk (Tolmachevo);Russia
XKH;VLXK;Phon Savan (Xieng Khouang);Laos
BIS;KBIS;Bismarck (Bismarck Municipal Airport);United States
TEX;KTEX;Telluride;United States
INC;ZLIC;Yinchuan;China
HGN;VTCH;Mae Hong Son;Thailand
RAP;KRAP;Rapid City (Rapid City Regional Airport);United States
CLD;KCRQ;Carlsbad (McClellan-Palomar Airport);United States
FNT;KFNT;Flint (Bishop International);United States
DVO;RPMD;Davao (Francisco Bangoy International);Philippines
FNC;LPMA;Funchal (Madeira);Portugal
STM;SBSN;Santarem;Brazil
KOS;VDSV;Sihanoukville;Cambodia
YOA;CYOA;Ekati;Canada
NPE;NZNR;NAPIER (Napier);New Zealand
LEV;NFNB;Levuka (Levuka Airfield);Fiji
LXA;ZULS;Lhasa (Lhasa-Gonggar);China
RDD;KRDD;Redding (Redding Muni);United States
EUG;KEUG;Eugene (Mahlon Sweet Fld);United States
IDA;KIDA;Idaho Falls (Idaho Falls Rgnl);United States
MFR;KMFR;Medford (Rogue Valley Intl Medford);United States
KBZ;NZKI;Kaikoura;New Zealand
RDM;KRDM;Redmond-Bend (Roberts Fld);United States
PCN;NZPN;Picton (Koromiko);New Zealand
WDH;FYWV;Windhoek (Windhoek Hosea Kutako International Airport );Namibia
YWH;CYWH;Victoria (Victoria Inner Harbour Airport);Canada
CXH;CAQ3;Vancouver (Vancouver Coal Harbour);Canada
TNA;ZSJN;Jinan;China
CZX;ZSCG;Changzhou;China
YBP;ZUYB;Yibin;China
TJM;USTR;Tyumen (Roschino);Russia
CAK;KCAK;Akron (Akron Canton Regional Airport);United States
HSV;KHSV;Huntsville (Huntsville International Airport-Carl T Jones Field);United States
PKB;KPKB;PARKERSBURG (Mid-Ohio Valley Regional Airport);United States
MGM;KMGM;MONTGOMERY (Montgomery Regional Airport );United States
TRI;KTRI;BRISTOL (Tri-Cities Regional Airport);United States
PAH;KPAH;PADUCAH (Barkley Regional Airport);United States
JIB;HDAM;Djibouti (Ambouli International Airport);Djibouti
HAK;ZJHK;Haikou (Meilan);China
MFA;HTMA;Mafia Island (Mafia);Tanzania
FCA;KFCA;Kalispell (Glacier Park Intl);United States
PGA;KPGA;Page (Page Municipal Airport);United States
UII;MHUT;Utila (Utila Airport);Honduras
MBS;KMBS;Saginaw (Mbs Intl);United States
BGM;KBGM;Binghamton (Greater Binghamton Edwin A Link Fld);United States
BGW;ORBI;Baghdad (Baghdad International Airport);Iraq
NNT;VTCN;Nan;Thailand
ROI;VTUV;Roi Et;Thailand
BFV;VTUO;Buri Ram;Thailand
TDX;VTBO;Trat;Thailand
BLH;KBLH;Blythe (Blythe Airport);United States
CRK;RPLC;Angeles City (Diosdado Macapagal International);Philippines
SDK;WBKS;Sandakan;Malaysia
LXG;VLLN;Luang Namtha;Laos
ODY;VLOS;Muang Xay (Oudomxay);Laos
SHE;ZYTX;Shenyang (Shenyang Taoxian International Airport);China
DOY;ZSDY;Dongying (Dongying Airport);China
MNI;TRPG;Geralds (John A. Osborne Airport);Montserrat
PSG;PAPG;Petersburg (Petersburg James A. Johnson);United States
LYA;ZHLY;Luoyang (Luoyang Airport);China
XUZ;ZSXZ;Xuzhou (Xuzhou Guanyin Airport);China
MWQ;VYMW;Magwe;Burma
KHM;VYKI;Khamti;Burma
DLI;VVDL;Dalat;Vietnam
VKG;VVRG;Rach Gia;Vietnam
CAH;VVCM;Ca Mau;Vietnam
VCL;VVCA;Chu Lai;Vietnam
TBB;VVTH;Tuy Hoa (Dong Tac);Vietnam
PYY;VTCI;Pai;Thailand
BWK;LDSB;Brac;Croatia
NSI;FKYS;Yaounde (Yaounde Nsimalen);Cameroon
CKY;GUCY;Conakry;Guinea
AAH;EDKA;Aachen (Flugplatz Merzbrueck);Germany
FKB;EDSB;Karlsruhe/Baden-Baden (Baden Airpark);Germany
SFB;KSFB;Sanford (Orlando Sanford Intl);United States
JST;KJST;Johnstown (John Murtha Johnstown-Cambria County Airport);United States
LUA;VNLK;Lukla;Nepal
BHP;VNBJ;Bhojpur;Nepal
LDN;VNLD;Lamidanda;Nepal
JMO;VNJS;Jomsom;Nepal
NGX;VNMA;Manang;Nepal
PPL;VNPL;Phaplu;Nepal
RUM;VNRT;Rumjatar;Nepal
DNP;VNDG;Dang (Tulsipur);Nepal
RUK;VNRK;Rukumkot;Nepal
JUM;VNJL;Jumla;Nepal
HRJ;VNCJ;Chaurjhari;Nepal
TPJ;VNTJ;Taplejung;Nepal
TMI;VNTR;Tumling Tar;Nepal
SKH;VNSK;Surkhet;Nepal
IMK;VNST;Simikot;Nepal
DOP;VNDP;Dolpa;Nepal
BJH;VNBG;Bajhang;Nepal
DHI;VNDH;Dhangarhi;Nepal
MWX;RKJB;Muan;South Korea
JTY;LGPL;Astypalaia;Greece
JIK;LGIK;Ikaria;Greece
JKL;LGKY;Kalymnos (Kalymnos Island);Greece
MLO;LGML;Milos;Greece
JNX;LGNX;Cyclades Islands (Naxos);Greece
PAS;LGPA;Paros;Greece
KZS;LGKJ;Kastelorizo;Greece
RMF;HEMA;Marsa Alam (Marsa Alam Intl);Egypt
NRN;EDLV;Weeze (Niederrhein);Germany
USU;RPVV;Busuanga;Philippines
BXU;RPME;Butuan;Philippines
DPL;RPMG;Dipolog;Philippines
LAO;RPLI;Laoag (Laoag Intl);Philippines
LGP;RPLP;Legazpi;Philippines
OZC;RPMO;Ozamis;Philippines
CEB;RPVM;Cebu (Mactan Cebu Intl);Philippines
NOE;EDWS;Norden (Sonderlandeplatz Norden-Norddeich);Germany
JUI;EDWJ;Juist (Verkehrslandeplatz Juist);Germany
BPS;SBPS;Porto Seguro (Aeroporto de Porto Seguro);Brazil
PMW;SBPJ;Palmas;Brazil
CLV;SBCN;Caldas Novas;Brazil
MSO;KMSO;Missoula (Missoula Intl);United States
BKQ;YBCK;Blackall;Australia
BDB;YBUD;Bundaberg;Australia
GCN;KGCN;Grand Canyon (Grand Canyon National Park Airport);United States
SGR;KSGR;Sugar Land (Sugar Land Regional Airport);United States
HIS;YHYN;Hayman Island (Hayman Island Airport);Australia
APA;KAPA;Denver (Centennial);United States
CVN;KCVN;Clovis (Clovis Muni);United States
FST;KFST;Fort Stockton (Fort Stockton Pecos Co);United States
LVS;KLVS;Las Vegas (Las Vegas Muni);United States
IWS;KIWS;Houston (West Houston);United States
LHX;KLHX;La Junta (La Junta Muni);United States
LRU;KLRU;Las Cruces (Las Cruces Intl);United States
BKD;KBKD;Breckenridge (Stephens Co);United States
TPL;KTPL;Temple (Draughon Miller Central Texas Rgnl);United States
OZA;KOZA;Ozona (Ozona Muni);United States
WKL;HI07;Waikoloa Village (Waikoloa Heliport);United States
LAK;CYKD;Aklavik;Canada
YWJ;CYWJ;Deline;Canada
ZFN;CZFN;Tulita;Canada
YGH;CYGH;Fort Good Hope;Canada
YPC;CYPC;Paulatuk;Canada
SAB;TNCS;Saba (Juancho E. Yrausquin);Netherlands Antilles
EGE;KEGE;Vail (Eagle Co Rgnl);United States
SKN;ENSK;Stokmarknes (Skagen);Norway
CGF;KCGF;Richmond Heights (Cuyahoga County);United States
MFD;KMFD;Mansfield (Mansfield Lahm Regional);United States
CSG;KCSG;Columbus (Columbus Metropolitan Airport);United States
LAW;KLAW;Lawton (Lawton-Fort Sill Regional Airport);United States
FNL;KFNL;Fort Collins (Fort Collins Loveland Muni);United States
FLG;KFLG;Flagstaff (Flagstaff Pulliam Airport);United States
TVL;KTVL;South Lake Tahoe (Lake Tahoe Airport);United States
TWF;KTWF;Twin Falls (Magic Valley Regional Airport);United States
MVY;KMVY;Vineyard Haven MA (Martha's Vineyard);United States
GON;KGON;Groton CT (Groton New London);United States
STC;KSTC;Saint Cloud (Saint Cloud Regional Airport);United States
GTR;KGTR;Columbus Mississippi (Golden Triangle Regional Airport);United States
GOJ;UWGG;Nizhniy Novgorod (Nizhny Novgorod);Russia
ERI;KERI;Erie (Erie Intl Tom Ridge Fld);United States
HYA;KHYA;Barnstable (Barnstable Muni Boardman Polando Fld);United States
SPR;MZ10;San Pedro;Belize
SDX;KSEZ;Sedona;United States
MGW;KMGW;Morgantown (Morgantown Muni Walter L Bill Hart Fld);United States
CRW;KCRW;Charleston (Yeager);United States
AVP;KAVP;Scranton (Wilkes Barre Scranton Intl);United States
BJI;KBJI;Bemidji (Bemidji Regional Airport);United States
THG;YTNG;Biloela (Thangool);Australia
FGI;NSFI;Apia (Fagali'i);Samoa
BNK;YBNA;Ballina Byron Bay (Ballina Byron Gateway);Australia
FAR;KFAR;Fargo (Hector International Airport);United States
MKC;KMKC;Kansas City (Downtown);United States
AZA;KIWA;Mesa (Phoenix-Mesa Gateway);United States
GCC;KGCC;Gillette (Gillette-Campbell County Airport);United States
TOF;UNTT;Tomsk (Tomsk Bogashevo Airport);Russia
NZJ;KNZJ;Santa Ana (El Toro);United States
PHY;VTPB;Phetchabun;Thailand
CJM;VTSE;Chumphon;Thailand
JZH;ZUJZ;Jiuzhaigou (Jiuzhaigou Huanglong);China
SWA;ZGOW;Shantou (Wai Sha Airport);China
GEO;SYCJ;Georgetown (Cheddi Jagan Intl);Guyana
AGT;SGES;Ciudad del Este;Paraguay
OGL;SYGO;Georgetown (Ogle);Guyana
KAI;SYKA;Kaieteur;Guyana
DNH;ZLDH;Dunhuang (Dunhuang Airport);China
AOI;LIPY;Ancona (Falconara);Italy
ECA;K6D9;East Tawas (Iosco County);United States
CPO;SCHA;Copiapo;Chile
TCP;HETB;Taba (Taba Intl);Egypt
LYB;MWCL;Little Cayman (Edward Bodden Airfield);Cayman Islands
BJV;LTFE;Bodrum (Bodrum - Milas);Turkey
TBJ;DTKA;Tabarka (7 Novembre);Tunisia
SAW;LTFJ;Istanbul (Sabiha Gokcen);Turkey
SCE;KUNV;State College Pennsylvania (University Park Airport);United States
BME;YPBR;Broome;Australia
NTL;YWLM;Newcastle (Newcastle Airport);Australia
KLU;LOWK;Klagenfurt (Woerthersee International Airport);Austria
HFT;ENHF;Hammerfest (Hammerfest Airport);Norway
HVG;ENHV;Honningsvag (Valan);Norway
MEH;ENMR;Mehamn;Norway
VDS;ENVD;Vadso (Airport);Norway
IKA;OIIE;Tehran (Imam Khomeini);Iran
MHD;OIMM;Mashhad;Iran
UIK;UIBS;Ust Ilimsk (Ust-Ilimsk);Russia
MEI;KMEI;Meridian (Key Field);United States
SPI;KSPI;Springfield (Abraham Lincoln Capital);United States
HKV;LB14;Haskovo (Uzundzhovo);Bulgaria
CEZ;KCEZ;Cortez (Cortez Muni);United States
HDN;KHDN;Hayden (Yampa Valley);United States
GUP;KGUP;Gallup (Gallup Muni);United States
LBL;KLBL;Liberal (Liberal Muni);United States
LAA;KLAA;Lamar (Lamar Muni);United States
GLD;KGLD;Goodland (Renner Fld);United States
COD;KCOD;Cody (Yellowstone Rgnl);United States
HOV;ENOV;Orsta-Volda (Hovden);Norway
SGF;KSGF;Springfield (Springfield Branson Natl);United States
NVK;ENNK;Narvik (Framnes);Norway
BVG;ENBV;Berlevag;Norway
FBU;ENFB;Oslo (Fornebu);Norway
NSK;UOOO;Norilsk (Alykel);Russia
AAQ;URKA;Anapa (Vityazevo);Russia
JLN;KJLN;Joplin (Joplin Rgnl);United States
ABE;KABE;Allentown (Lehigh Valley Intl);United States
XNA;KXNA;Bentonville (NW Arkansas Regional);United States
GUW;UATG;Atyrau;Kazakhstan
KZO;UAOO;Kzyl-Orda;Kazakhstan
SBN;KSBN;South Bend (South Bend Rgnl);United States
BKA;UUBB;Moscow (Bykovo);Russia
ARH;ULAA;Arkhangelsk (Talagi);Russia
RTW;UWSS;Saratov (Central);Russia
NUX;USMU;Novy Urengoy (Novyi Urengoy);Russia
NOJ;USRO;Noyabrsk;Russia
SCO;UATE;Aktau;Kazakhstan
NNM;ULAM;Naryan-Mar;Russia
PKV;ULOO;Pskov (Kresty);Russia
KGP;USRK;Kogalym (Kogalym International);Russia
KJA;UNKL;Krasnoyarsk (Emelyanovo);Russia
KGF;UAKK;Karaganda (Sary-Arka);Kazakhstan
URJ;USHU;Uraj;Russia
IWA;UUBI;Ivanovo (Yuzhny);Russia
CGQ;ZYCC;Changchun;China
KIJ;RJSN;Niigata;Japan
JON;PJON;Johnston Island (Johnston Atoll);Johnston Atoll
SMD;KSMD;Fort Wayne IN (Smith Fld);United States
ACV;KACV;Arcata CA (Arcata);United States
ATT;KATT;Austin TX (Camp Mabry Austin City);United States
OAJ;KOAJ;Jacksonville NC (Albert J Ellis);United States
TCL;KTCL;Tuscaloosa AL (Tuscaloosa Rgnl);United States
DBQ;KDBQ;Dubuque IA (Dubuque Rgnl);United States
PYJ;UERP;Yakutia (Poliarny Airport);Russia
NAJ;UBBN;Nakhchivan (Nakhchivan Airport);Azerbaijan
KVD;UBBG;Ganja (Ganja Airport);Azerbaijan
UKK;UASK;Ust Kamenogorsk (Ust Kamenogorsk Airport);Kazakhstan
PPK;UACP;Petropavlosk (Petropavlosk South Airport);Kazakhstan
GBJ;TFFM;Grand Bourg (Les Bases Airport);Guadeloupe
SFC;TFFC;St-François (St-François Airport);Guadeloupe
BBQ;TAPH;Codrington (Codrington Airport);Antigua and Barbuda
JPR;SWJI;Ji-Paraná (Ji-Paraná Airport);Brazil
MYC;SVBS;Maracay (Escuela Mariscal Sucre Airport);Venezuela
NZA;SPZA;Nazca (Maria Reiche Neuman Airport);Peru
CJA;SPJR;Cajamarca (Mayor General FAP Armando Revoredo Iglesias Airport);Peru
MVS;SNMU;Mucuri (Mucuri Airport);Brazil
ORG;SMZO;Paramaribo (Zorg en Hoop Airport);Suriname
REY;SLRY;Reyes (Reyes Airport);Bolivia
PUR;SLPR;Puerto Rico/Manuripi (Puerto Rico Airport);Bolivia
EYP;SKYP;Yopal (El Alcaraván Airport);Colombia
ESM;SETN;Esmeraldas (General Rivadeneira Airport);Ecuador
ZPC;SCPC;Pucon (Pucón Airport);Chile
TOW;SBTD;Toledo (Toledo Airport);Brazil
RIA;SBSM;Santa Maria (Santa Maria Airport);Brazil
LEC;SBLE;Lençóis (Chapada Diamantina Airport);Brazil
GUL;YGLB;Goulburn (Goulburn Airport);Australia
JDO;SBJU;Juazeiro Do Norte (Orlando Bezerra de Menezes Airport);Brazil
SST;SAZL;Santa Teresita (Santa Teresita Airport);Argentina
GGS;SAWR;Gobernador Gregores (Gobernador Gregores Airport);Argentina
OES;SAVN;San Antonio Oeste (Antoine De St Exupery Airport);Argentina
LHS;SAVH;Las Heras (Las Heras Airport);Argentina
TTG;SAST;Tartagal (General Enrique Mosconi Airport);Argentina
MBT;RPVJ;Masbate (Masbate Airport);Philippines
CRM;RPVF;Catarman (Catarman National Airport);Philippines
JOL;RPMJ;Jolo (Jolo Airport);Philippines
CGM;RPMH;Camiguin (Camiguin Airport);Philippines
CYU;RPLO;Cuyo (Cuyo Airport);Philippines
CJJ;RKTU;Chongju (Cheongju International Airport);South Korea
HIN;RKPS;Sacheon (Sacheon Air Base);South Korea
WJU;RKNW;Wonju (Wonju Airport);South Korea
MPK;RKJM;Mokpo (Mokpo Airport);South Korea
KUV;RKJK;Kunsan (Kunsan Air Base);South Korea
MYE;RJTQ;Miyakejima (Miyakejima Airport);Japan
SYO;RJSY;Shonai (Shonai Airport);Japan
ONJ;RJSR;Odate Noshiro (Odate Noshiro Airport);Japan
FKS;RJSF;Fukushima (Fukushima Airport);Japan
IWJ;RJOW;Iwami (Iwami Airport);Japan
NKM;RJNA;Nagoya (Nagoya Airport);Japan
HSG;RJFS;Saga (Saga Airport);Japan
OKD;RJCO;Sapporo (Okadama Airport);Japan
KUH;RJCK;Kushiro (Kushiro Airport);Japan
MFK;RCMT;Matsu Islands (Matsu Beigan Airport);Taiwan
HCN;RCKW;Hengchun (Hengchun Airport);Taiwan
LZN;RCFG;Matsu Islands (Matsu Nangan Airport);Taiwan
ENT;PKMA;Eniwetok Atoll (Eniwetok Airport);Marshall Islands
LUP;PHLU;Molokai (Kalaupapa Airport);United States
WRG;PAWG;Wrangell (Wrangell Airport);United States
VAK;PAVA;Chevak (Chevak Airport);United States
ANI;PANI;Aniak (Aniak Airport);United States
MOU;PAMO;Mountain Village (Mountain Village Airport);United States
MCG;PAMC;Mcgrath (McGrath Airport);United States
KLG;PALG;Kalskag (Kalskag Airport);United States
HNS;PAHN;Haines (Haines Airport);United States
HCR;PAHC;Holy Cross (Holy Cross Airport);United States
SGY;PAGY;Skagway (Skagway Airport);United States
GST;PAGS;Gustavus (Gustavus Airport);United States
ADK;PADK;Adak Island (Adak Airport);United States
GXF;OYSY;Sayun Intl (Sayun International Airport);Yemen
KAC;OSKL;Kamishly (Kamishly Airport);Syria
ISU;ORSU;Sulaymaniyah (Sulaymaniyah International Airport);Iraq
TUK;OPTU;Turbat (Turbat International Airport);Pakistan
SYW;OPSN;Sehwan Sharif (Sehwan Sharif Airport);Pakistan
KDU;OPSD;Skardu (Skardu Airport);Pakistan
PAJ;OPPC;Parachinar (Parachinar Airport);Pakistan
ORW;OPOR;Ormara Raik (Ormara Airport);Pakistan
KDD;OPKH;Khuzdar (Khuzdar Airport);Pakistan
HDD;OPKD;Hyderabad (Hyderabad Airport);Pakistan
JIW;OPJI;Jiwani (Jiwani Airport);Pakistan
DSK;OPDI;Dera Ismael Khan (Dera Ismael Khan Airport);Pakistan
DEA;OPDG;Dera Ghazi Khan (Dera Ghazi Khan Airport);Pakistan
DBA;OPDB;Dalbandin (Dalbandin Airport);Pakistan
CJL;OPCH;Chitral (Chitral Airport);Pakistan
BHV;OPBW;Bahawalpur (Bahawalpur Airport);Pakistan
BNP;OPBN;Bannu (Bannu Airport);Pakistan
AAN;OMAL;Al Ain (Al Ain International Airport);United Arab Emirates
OMH;OITR;Uromiyeh (Uromiyeh Airport);Iran
ADU;OITL;Ardabil (Ardabil Airport);Iran
LRR;OISL;Lar (Lar Airport);Iran
SRY;OINZ;Dasht-e-naz (Sari Dasht E Naz Airport);Iran
NSH;OINN;Noshahr (Noshahr Airport);Iran
AFZ;OIMS;Sabzevar (Sabzevar National Airport);Iran
BJB;OIMN;Bojnourd (Bojnourd Airport);Iran
RJN;OIKR;Rafsanjan (Rafsanjan Airport);Iran
BXR;OIKM;Bam (Bam Airport);Iran
KHD;OICK;Khorram Abad (Khoram Abad Airport);Iran
EWD;OEWD;Wadi-al-dawasir (Wadi Al Dawasir Airport);Saudi Arabia
AJF;OESK;Al-Jawf (Al-Jawf Domestic Airport);Saudi Arabia
DWD;OEDW;Dawadmi (Dawadmi Domestic Airport);Saudi Arabia
XAU;SOOS;Saul (Saul Airport);French Guiana
FBD;OAFZ;Faizabad (Faizabad Airport);Afghanistan
ILP;NWWE;Île des Pins (Île des Pins Airport);New Caledonia
BMY;NWWC;Waala (Belep Islands Airport);New Caledonia
TGJ;NWWA;Tiga (Tiga Airport);New Caledonia
IPA;NVVI;Ipota (Ipota Airport);Vanuatu
FTA;NVVF;Futuna Island (Futuna Airport);Vanuatu
DLY;NVVD;Dillon's Bay (Dillon's Bay Airport);Vanuatu
AWD;NVVB;Aniwa (Aniwa Airport);Vanuatu
AUY;NVVA;Anelghowhat (Anelghowhat Airport);Vanuatu
OLZ;NVSZ;Olpoi (North West Santo Airport);Vanuatu
SWJ;NVSX;Malekula Island (Southwest Bay Airport);Vanuatu
VLS;NVSV;Valesdir (Valesdir Airport);Vanuatu
ULB;NVSU;Ambryn Island (Uléi Airport);Vanuatu
TGH;NVST;Tongoa Island (Tongoa Island Airport);Vanuatu
SON;NVSS;Santo (Santo Pekoa International Airport);Vanuatu
RCL;NVSR;Redcliffe (Redcliffe Airport);Vanuatu
ZGU;NVSQ;Gaua Island (Gaua Island Airport);Vanuatu
NUS;NVSP;Norsup (Norsup Airport);Vanuatu
LNE;NVSO;Lonorore (Lonorore Airport);Vanuatu
MWF;NVSN;Maewo Island (Naone Airport);Vanuatu
LNB;NVSM;Lamen Bay (Lamen Bay Airport);Vanuatu
LPM;NVSL;Lamap (Lamap Airport);Vanuatu
PBJ;NVSI;Paama Island (Tavie Airport);Vanuatu
SSR;NVSH;Pentecost Island (Sara Airport);Vanuatu
LOD;NVSG;Longana (Longana Airport);Vanuatu
CCV;NVSF;Craig Cove (Craig Cove Airport);Vanuatu
EAE;NVSE;Sangafa (Sangafa Airport);Vanuatu
TOH;NVSD;Loh/Linua (Torres Airstrip);Vanuatu
SLH;NVSC;Sola (Sola Airport);Vanuatu
MTV;NVSA;Ablow (Mota Lava Airport);Vanuatu
UAH;NTMU;Ua Huka (Ua Huka Airport);French Polynesia
UAP;NTMP;Ua Pou (Ua Pou Airport);French Polynesia
AUQ;NTMN;Hiva-oa (Hiva Oa-Atuona Airport);French Polynesia
AHE;NTHE;Ahe (Ahe Airport);French Polynesia
APK;NTGD;Apataki (Apataki Airport);French Polynesia
MXS;NSMA;Savaii Island (Maota Airport);Samoa
FUT;NLWF;Futuna Island (Pointe Vele Airport);Wallis and Futuna
IUE;NIUE;Alofi (Niue International Airport);Niue
VBV;NFVB;Vanua Balavu (Vanua Balavu Airport);Fiji
NTT;NFTP;Niuatoputapu (Mata'aho Airport);Tonga
NFO;NFTO;Niuafoʻou Airport (Kuini Lavinia Airport); Tonga
HPA;NFTL;Lifuka (Lifuka Island Airport);Tonga
EUA;NFTE;Eua Island (Kaufana Airport);Tonga
SVU;NFNS;Savusavu (Savusavu Airport);Fiji
RTA;NFNR;Rotuma (Rotuma Airport);Fiji
KXF;NFNO;Koro Island (Koro Island Airport);Fiji
TVU;NFNM;Matei (Matei Airport);Fiji
LBS;NFNL;Lambasa (Labasa Airport);Fiji
LKB;NFNK;Lakeba Island (Lakeba Island Airport);Fiji
NGI;NFNG;Ngau (Ngau Airport);Fiji
MFJ;NFMO;Moala (Moala Airport);Fiji
MNF;NFMA;Mana Island (Mana Island Airport);Fiji
KDV;NFKD;Vunisea (Vunisea Airport);Fiji
PTF;NFFO;Malolo Lailai Island (Malolo Lailai Island Airport);Fiji
ICI;NFCI;Cicia (Cicia Airport);Fiji
PYE;NCPY;Penrhyn Island (Penrhyn Island Airport);Cook Islands
MOI;NCMR;Mitiaro Island (Mitiaro Island Airport);Cook Islands
MUK;NCMK;Mauke Island (Mauke Airport);Cook Islands
MHX;NCMH;Manihiki Island (Manihiki Island Airport);Cook Islands
MGS;NCMG;Mangaia Island (Mangaia Island Airport);Cook Islands
AIU;NCAT;Atiu Island (Atiu Island Airport);Cook Islands
PID;MYPI;Nassau (Nassau Paradise Island Airport);Bahamas
CRI;MYCI;Colonel Hill (Colonel Hill Airport);Bahamas
CAT;MYCB;Cat Island (New Bight Airport);Bahamas
ATC;MYCA;Arthur's Town (Arthurs Town Airport);Bahamas
COX;MYAK;Andros (Congo Town Airport);Bahamas
SCX;MM57;Salina Cruz (Salina Cruz Naval Air Station);Mexico
TND;MUTD;Trinidad (Alberto Delgado Airport);Cuba
CCC;MUOC;Cayo Coco (Cayo Coco Airport);Cuba
PAX;MTPX;Port-de-Paix (Port-de-Paix Airport);Haiti
JEE;MTJE;Jeremie (Jeremie Airport);Haiti
PLD;MRSR;Playa Samara (Playa Samara Airport);Costa Rica
SYQ;MRPV;San Jose (Tobias Bolanos International Airport);Costa Rica
PJM;MRPJ;Puerto Jimenez (Puerto Jimenez Airport);Costa Rica
PBP;MRIA;Nandayure (Islita Airport);Costa Rica
TNO;MRCV;Nicoya (Cabo Velas Airport);Costa Rica
BCL;MRBC;Pococi (Barra del Colorado Airport);Costa Rica
TTQ;MRAO;Roxana (Aerotortuguero Airport);Costa Rica
PLP;MPLP;La Palma (Captain Ramon Xatruch Airport);Panama
JQE;MPJE;Jaqué (Jaqué Airport);Panama
ONX;MPEJ;Colón (Enrique Adolfo Jimenez Airport);Panama
CTD;MPCE;Chitré (Alonso Valderrama Airport);Panama
JAL;MMJA;Jalapa (Lencero Airport);Mexico
GUB;MMGR;Guerrero Negro (Guerrero Negro Airport);Mexico
CUA;MMDA;Ciudad Constitución (Ciudad Constitución Airport);Mexico
CYW;MMCY;Celaya (Captain Rogelio Castillo National Airport);Mexico
MIJ;MLIP;Mili Island (Mili Island Airport);Marshall Islands
PEU;MHPL;Puerto Lempira (Puerto Lempira Airport);Honduras
AHS;MHAH;Ahuas (Ahuas Airport);Honduras
WTE;N36;Wotje Atoll (Wotje Atoll Airport);Marshall Islands
UIT;N55;Jabor Jaluit Atoll (Jaluit Airport);Marshall Islands
NDK;3N0;Namorik Atoll (Namorik Atoll Airport);Marshall Islands
MJB;Q30;Mejit Atoll (Mejit Atoll Airport);Marshall Islands
UTK;03N;Utirik Island (Utirik Airport);Marshall Islands
AAZ;MGQZ;Quezaltenango (Quezaltenango Airport);Guatemala
PBR;MGPB;Puerto Barrios (Puerto Barrios Airport);Guatemala
JBQ;MDJB;La Isabela (Dr Joaquin Balaguer International Airport);Dominican Republic
AZS;MDCY;Samana (Samaná El Catey International Airport);Dominican Republic
SLX;MBSY;Salt Cay (Salt Cay Airport);Turks and Caicos Islands
MDS;MBMC;Middle Caicos (Middle Caicos Airport);Turks and Caicos Islands
GDT;MBGT;Cockburn Town (JAGS McCartney International Airport);Turks and Caicos Islands
ILZ;LZZI;Žilina (Žilina Airport);Slovakia
SZF;LTFH;Samsun (Samsun-Çarşamba Airport);Turkey
EDO;LTFD;Balikesir Korfez (Balikesir Korfez Airport);Turkey
ISE;LTFC;Isparta (Isparta Süleyman Demirel Airport);Turkey
ADF;LTCP;Adiyaman (Adiyaman Airport);Turkey
AJI;LTCO;Agri (Agri Airport);Turkey
KCM;LTCN;Kahramanmaras (Kahramanmaras Airport);Turkey
SFQ;LTCH;Sanliurfa (Sanliurfa Airport);Turkey
KSY;LTCF;Kars (Kars Airport);Turkey
USQ;LTBO;Usak (Usak Airport);Turkey
BNX;LQBK;Banja Luka (Banja Luka International Airport);Bosnia and Herzegovina
CVU;LPCR;Corvo (Corvo Airport);Portugal
QSR;LIRI;Salerno (Salerno Pontecagnano Airport);Italy
AOT;LIMW;Aosta (Aosta Airport);Italy
SOB;LHSM;Sármellék (Sármellék International Airport);Hungary
QGY;LHPR;Győr (Győr-Pér International Airport);Hungary
PEV;LHPP;Pécs-Pogány (Pécs-Pogány Airport);Hungary
JSY;LGSO;Syros Island (Syros Airport);Greece
LTT;LFTZ;La Môle (La Môle Airport);France
ANE;LFJR;Angers/Marcé (Angers-Loire Airport);France
IDY;LFEY;Île d'Yeu (Île d'Yeu Airport);France
RJL;LELO;Logroño-Agoncillo (Logroño-Agoncillo Airport);Spain
ECN;LCEN;Nicosia (Ercan International Airport);Cyprus
YKM;KYKM;Yakima (Yakima Air Terminal McAllister Field);United States
WRL;KWRL;Worland (Worland Municipal Airport);United States
VLD;KVLD;Valdosta (Valdosta Regional Airport);United States
VCT;KVCT;Victoria (Victoria Regional Airport);United States
UIN;KUIN;Quincy (Quincy Regional Baldwin Field);United States
TUP;KTUP;Tupelo (Tupelo Regional Airport);United States
SMX;KSMX;Santa Maria (Santa Maria Pub Cpt G Allan Hancock Airport);United States
SLN;KSLN;Salina (Salina Municipal Airport);United States
SLK;KSLK;Saranac Lake (Adirondack Regional Airport);United States
SHR;KSHR;Sheridan (Sheridan County Airport);United States
SBP;KSBP;San Luis Obispo (San Luis County Regional Airport);United States
RUT;KRUT;Rutland (Rutland State Airport);United States
RKS;KRKS;Rock Springs (Rock Springs Sweetwater County Airport);United States
RHI;KRHI;Rhinelander (Rhinelander Oneida County Airport);United States
RDG;KRDG;Reading (Reading Regional Carl A Spaatz Field);United States
PSM;KPSM;Portsmouth (Pease International Tradeport);United States
PLN;KPLN;Pellston (Pellston Regional Airport of Emmet County Airport);United States
PIR;KPIR;Pierre (Pierre Regional Airport);United States
PIH;KPIH;Pocatello (Pocatello Regional Airport);United States
PIB;KPIB;Hattiesburg/Laurel (Hattiesburg Laurel Regional Airport);United States
OWB;KOWB;Owensboro (Owensboro Daviess County Airport);United States
OTH;KOTH;North Bend (Southwest Oregon Regional Airport);United States
MSL;KMSL;Muscle Shoals (Northwest Alabama Regional Airport);United States
MLS;KMLS;Miles City (Frank Wiley Field);United States
MKG;KMKG;Muskegon (Muskegon County Airport);United States
LYH;KLYH;Lynchburg (Lynchburg Regional Preston Glenn Field);United States
LWT;KLWT;Lewistown (Lewistown Municipal Airport);United States
LNS;KLNS;Lancaster (Lancaster Airport);United States
LMT;KLMT;Klamath Falls (Klamath Falls Airport);United States
LEB;KLEB;Lebanon (Lebanon Municipal Airport);United States
LBF;KLBF;North Platte (North Platte Regional Airport Lee Bird Field);United States
LBE;KLBE;Latrobe (Arnold Palmer Regional Airport);United States
LAR;KLAR;Laramie (Laramie Regional Airport);United States
JMS;KJMS;Jamestown (Jamestown Regional Airport);United States
IRK;KIRK;Kirksville (Kirksville Regional Airport);United States
KIO;Q51;Kili Island (Kili Airport);Marshall Islands
HTS;KHTS;Huntington (Tri State Milton J Ferguson Field);United States
HOT;KHOT;Hot Springs (Memorial Field);United States
GRI;KGRI;Grand Island (Central Nebraska Regional Airport);United States
GGW;KGGW;Glasgow (Wokal Field Glasgow International Airport);United States
FAY;KFAY;Fayetteville (Fayetteville Regional Grannis Field);United States
EWB;KEWB;New Bedford (New Bedford Regional Airport);United States
EKO;KEKO;Elko (Elko Regional Airport);United States
EAU;KEAU;Eau Claire (Chippewa Valley Regional Airport);United States
DUJ;KDUJ;Du Bois (DuBois Regional Airport);United States
DDC;KDDC;Dodge City (Dodge City Regional Airport);United States
CMX;KCMX;Hancock (Houghton County Memorial Airport);United States
CLM;KCLM;Port Angeles (William R Fairchild International Airport);United States
CKB;KCKB;Clarksburg (Harrison Marion Regional Airport);United States
CIU;KCIU;Sault Ste Marie (Chippewa County International Airport);United States
CGI;KCGI;Cape Girardeau (Cape Girardeau Regional Airport);United States
CEC;KCEC;Crescent City (Del Norte County Airport);United States
BRL;KBRL;Burlington (Southeast Iowa Regional Airport);United States
BQK;KBQK;Brunswick (Brunswick Golden Isles Airport);United States
BKW;KBKW;Beckley (Raleigh County Memorial Airport);United States
BFF;KBFF;Scottsbluff (Western Nebraska Regional Airport);United States
BFD;KBFD;Bradford (Bradford Regional Airport);United States
ATY;KATY;Watertown (Watertown Regional Airport);United States
APN;KAPN;Alpena (Alpena County Regional Airport);United States
ALW;KALW;Walla Walla (Walla Walla Regional Airport);United States
ALO;KALO;Waterloo (Waterloo Regional Airport);United States
ALM;KALM;Alamogordo (Alamogordo White Sands Regional Airport);United States
AHN;KAHN;Athens (Athens Ben Epps Airport);United States
ABY;KABY;Albany (Southwest Georgia Regional Airport);United States
ABR;KABR;Aberdeen (Aberdeen Regional Airport);United States
DIU;VA1P;Diu (Diu Airport);India
ULU;HUGU;Gulu (Gulu Airport);Uganda
RUA;HUAR;Arua (Arua Airport);Uganda
TBO;HTTB;Tabora (Tabora Airport);Tanzania
SHY;HTSY;Shinyanga (Shinyanga Airport);Tanzania
MUZ;HTMU;Musoma (Musoma Airport);Tanzania
LDI;HTLI;Lindi (Kikwetu Airport);Tanzania
TKQ;HTKA;Kigoma (Kigoma Airport);Tanzania
BKZ;HTBU;Bukoba (Bukoba Airport);Tanzania
PZU;HSPN;Port Sudan (Port Sudan New International Airport);Sudan
UYL;HSNN;Nyala (Nyala Airport);Sudan
ATB;HSAT;Atbara (Atbara Airport);Sudan
LAQ;HLLQ;Al Bayda' (La Abraq Airport);Libya
MJI;HLLM;Tripoli (Mitiga Airport);Libya
TOB;HLGN;Tobruk (Gamal Abdel Nasser Airport);Libya
SRX;HLGD;Sirt (Gardabya Airport);Libya
NYK;HKNY;Nanyuki (Nanyuki Civil Airport);Kenya
MYD;HKML;Malindi (Malindi Airport);Kenya
LKG;HKLK;Lokichoggio (Lokichoggio Airport);Kenya
ASV;HKAM;Amboseli National Park (Amboseli Airport);Kenya
ATZ;HEAT;Asyut (Asyut International Airport);Egypt
AAC;HEAR;El Arish (El Arish International Airport);Egypt
BUO;HCMV;Burao (Burao Airport);Somalia
GLK;HCMR;Galcaio (Galcaio Airport);Somalia
MGQ;HCMM;Mogadishu (Aden Adde International Airport);Somalia
BSA;HCMF;Bosaso (Bosaso Airport);Somalia
ALU;HCMA;Alula (Alula Airport);Somalia
TIE;HATP;Tippi (Tippi Airport);Ethiopia
MTF;HAMT;Mizan Teferi (Mizan Teferi Airport);Ethiopia
ABK;HAKD;Kabri Dehar (Kabri Dehar Airport);Ethiopia
GOR;HAGR;Gore (Gore Airport);Ethiopia
GDE;HAGO;Gode (Gode Airport);Ethiopia
DEM;HADD;Dembidollo (Dembidollo Airport);Ethiopia
DSE;HADC;Dessie (Combolcha Airport);Ethiopia
BEI;HABE;Beica (Beica Airport);Ethiopia
BCO;HABC;Baco (Baco Airport);Ethiopia
SFL;GVSF;Sao Filipe (Sao Filipe Airport); Fogo Island
RAI;GVNP;Praia (Praia International Airport); Santiago Island
NDR;GMMW;El Aroui (El Aroui Airport);Morocco
EUN;GMML;El Aaiún (Hassan I Airport);Western Sahara
ESU;GMMI;Essadouira (Mogador Airport);Morocco
VIL;GMMH;Dakhla (Dakhla Airport);Western Sahara
SMW;GMMA;Smara (Smara Airport);Western Sahara
OXB;GGOV;Bissau (Osvaldo Vieira International Airport);Guinea-Bissau
KEN;GFKE;Kenema (Kenema Airport);Sierra Leone
KBS;GFBO;Bo (Bo Airport);Sierra Leone
BTE;GFBN;Bonthe (Sherbro International Airport);Sierra Leone
JCU;GECT;Ceuta (Ceuta Heliport);Spain
GMZ;GCGM;La Gomera (La Gomera Airport);Spain
PFR;FZVS;Ilebo (Ilebo Airport);Congo (Kinshasa)
LJA;FZVA;Lodja (Lodja Airport);Congo (Kinshasa)
TSH;FZUK;Tshikapa (Tshikapa Airport);Congo (Kinshasa)
BSU;FZEN;Basankusu (Basankusu Airport);Congo (Kinshasa)
KRZ;FZBT;Kiri (Basango Mboliasa Airport);Congo (Kinshasa)
NIO;FZBI;Nioki (Nioki Airport);Congo (Kinshasa)
INO;FZBA;Inongo (Inongo Airport);Congo (Kinshasa)
MAT;FZAM;Matadi (Tshimpi Airport);Congo (Kinshasa)
BOA;FZAJ;Boma (Boma Airport);Congo (Kinshasa)
ERS;FYWE;Windhoek (Eros Airport);Namibia
SWP;FYSM;Swakopmund (Swakopmund Airport);Namibia
OMD;FYOG;Oranjemund (Oranjemund Airport);Namibia
OND;FYOA;Ondangwa (Ondangwa Airport);Namibia
LUD;FYLZ;Luderitz (Luderitz Airport);Namibia
CMK;FWCM;Club Makokola (Club Makokola Airport);Malawi
SRH;FTTA;Sarh (Sarh Airport);Chad
VPY;FQCH;Chimoio (Chimoio Airport);Mozambique
TCH;FOOT;Tchibanga (Tchibanga Airport);Gabon
MJL;FOGM;Mouila (Mouilla Ville Airport);Gabon
KOU;FOGK;Koulamoutou (Koulamoutou Airport);Gabon
MSZ;FNMO;Mocamedes (Namibe Airport);Angola
VPE;FNGI;Ondjiva (Ondjiva Pereira Airport);Angola
DUE;FNDU;Dundo (Dundo Airport);Angola
CBT;FNCT;Catumbela (Catumbela Airport);Angola
MJA;FMSJ;Manja (Manja Airport);Madagascar
WMA;FMNX;Mandritsara (Mandritsara Airport);Madagascar
TTS;FMNT;Tsaratanana (Tsaratanana Airport);Madagascar
WMP;FMNP;Mampikony (Mampikony Airport);Madagascar
DWB;FMNO;Soalala (Soalala Airport);Madagascar
IVA;FMNJ;Ambanja (Ambanja Airport);Madagascar
WPB;FMNG;Port Bergé (Port Bergé Airport);Madagascar
WAM;FMMZ;Ambatondrazaka (Ambatondrazaka Airport);Madagascar
WTS;FMMX;Tsiroanomandidy (Tsiroanomandidy Airport);Madagascar
WTA;FMMU;Tambohorano (Tambohorano Airport);Madagascar
TVA;FMMR;Morafenobe (Morafenobe Airport);Madagascar
MXT;FMMO;Maintirano (Maintirano Airport);Madagascar
BMD;FMML;Belo sur Tsiribihina (Belo sur Tsiribihina Airport);Madagascar
JVA;FMMK;Ankavandra (Ankavandra Airport);Madagascar
WAQ;FMMG;Antsalova (Antsalova Airport);Madagascar
YVA;FMCN;Moroni (Iconi Airport);Comoros
SLI;FLSW;Solwesi (Solwesi Airport);Zambia
CIP;FLCP;Chipata (Chipata Airport);Zambia
DIS;FCPL;Loubomo (Loubomo Airport);Congo (Brazzaville)
TLD;FBTL;Tuli Lodge (Limpopo Valley Airport);Botswana
SWX;FBSW;Shakawe (Shakawe Airport);Botswana
ORP;FBOR;Orapa (Orapa Airport);Botswana
GNZ;FBGZ;Ghanzi (Ghanzi Airport);Botswana
MBD;FAMM;Mafeking (Mmabatho International Airport);South Africa
AAM;FAMD;Malamala (Malamala Airport);South Africa
MQP;FAKN;Mpumalanga (Kruger Mpumalanga International Airport);South Africa
QRA;FAGM;Johannesburg (Rand Airport);South Africa
VTS;EVVA;Ventspils (Ventspils International Airport);Latvia
HMV;ESUT;Hemavan (Hemavan Airport);Sweden
SQO;ESUD;Mohed (Storuman Airport);Sweden
AGH;ESTA;Ängelholm (Ängelholm-Helsingborg Airport);Sweden
TYF;ESST;Torsby (Torsby Airport);Sweden
KSD;ESOK;Karlstad (Karlstad Airport);Sweden
HFS;ESOH;Hagfors (Hagfors Airport);Sweden
OSD;ESNZ;Östersund (Östersund Airport);Sweden
LCJ;EPLL;Lodz (Łódź Władysław Reymont Airport);Poland
BZG;EPBY;Bydgoszcz (Bydgoszcz Ignacy Jan Paderewski Airport);Poland
VRY;ENVR;Værøy (Værøy Heliport);Norway
VAW;ENSS;Vardø (Svartnes Airport);Norway
SOJ;ENSR;Sorkjosen (Sorkjosen Airport);Norway
SVJ;ENSH;Svolvær (Svolvær Helle Airport);Norway
SOG;ENSG;Sogndal (Sogndal Airport);Norway
SDN;ENSD;Sandane (Anda Airport);Norway
RET;ENRS;Røst (Røst Airport);Norway
RVK;ENRM;Rørvik (Ryum Airport);Norway
MQN;ENRA;Mo i Rana (Røssvoll Airport);Norway
OSY;ENNM;Namsos (Namsos Høknesøra Airport);Norway
LKN;ENLK;Leknes (Leknes Airport);Norway
CNL;EKSN;Sindal (Sindal Airport);Denmark
CFN;EIDL;Dongloe (Donegal Airport);Ireland
BRR;EGPR;Barra (Barra Airport);United Kingdom
HLY;EGOV;Angelsey (Anglesey Airport);United Kingdom
PZE;EGHK;Penzance (Penzance Heliport);United Kingdom
LEQ;EGHC;Land's End (Land's End / St. Just Airport);United Kingdom
WRY;EGEW;Westray (Westray Airport);United Kingdom
LWK;EGET;Lerwick (Lerwick / Tingwall Airport);United Kingdom
NDY;EGES;Sanday (Sanday Airport);United Kingdom
SOY;EGER;Stronsay (Stronsay Airport);United Kingdom
PPW;EGEP;Papa Westray (Papa Westray Airport);United Kingdom
NRL;EGEN;North Ronaldsay (North Ronaldsay Airport);United Kingdom
FIE;EGEF;Fair Isle (Fair Isle Airport);United Kingdom
EOI;EGED;Eday (Eday Airport);United Kingdom
CAL;EGEC;Campbeltown (Campbeltown Airport);United Kingdom
DSA;EGCN;Doncaster-Sheffield (Robin Hood Doncaster Sheffield Airport);United Kingdom
NQT;EGBN;Nottingham (Nottingham Airport);United Kingdom
SJY;EFSI;Seinäjoki / Ilmajoki (Seinäjoki Airport);Finland
HGL;EDXH;Helgoland (Helgoland-Düne Airport);Germany
HEI;EDXB;Büsum (Heide-Büsum Airport);Germany
HDF;EDAH;Heringsdorf (Heringsdorf Airport);Germany
KMS;DGSI;Kumasi (Kumasi Airport);Ghana
ELU;DAUO;Guemar (Guemar Airport);Algeria
BMW;DATM;Bordj Badji Mokhtar (Bordj Badji Mokhtar Airport);Algeria
CBH;DAOR;Béchar (Béchar Boudghene Ben Ali Lotfi Airport);Algeria
BLJ;DABT;Batna (Batna Airport);Algeria
ZWL;CZWL;Wollaston Lake (Wollaston Lake Airport);Canada
ZUM;CZUM;Churchill Falls (Churchill Falls Airport);Canada
ZTM;CZTM;Shamattawa (Shamattawa Airport);Canada
ZSJ;CZSJ;Sandy Lake (Sandy Lake Airport);Canada
ZRJ;CZRJ;Round Lake (Round Lake (Weagamow Lake) Airport);Canada
ZPB;CZPB;Sachigo Lake (Sachigo Lake Airport);Canada
ZMT;CZMT;Masset (Masset Airport);Canada
MSA;CZMD;Muskrat Dam (Muskrat Dam Airport);Canada
ZKE;CZKE;Kashechewan (Kashechewan Airport);Canada
ZJN;CZJN;Swan River (Swan River Airport);Canada
ZGI;CZGI;Gods River (Gods River Airport);Canada
ZFD;CZFD;Fond-Du-Lac (Fond-Du-Lac Airport);Canada
ZEM;CZEM;Eastmain River (Eastmain River Airport);Canada
ZBF;CZBF;Bathurst (Bathurst Airport);Canada
ILF;CZBD;Ilford (Ilford Airport);Canada
ZAC;CZAC;York Landing (York Landing Airport);Canada
YZG;CYZG;Salluit (Salluit Airport);Canada
YXN;CYXN;Whale Cove (Whale Cove Airport);Canada
YWP;CYWP;Webequie (Webequie Airport);Canada
YVZ;CYVZ;Deer Lake (Deer Lake Airport);Canada
YTL;CYTL;Big Trout Lake (Big Trout Lake Airport);Canada
YST;CYST;St. Theresa Point (St. Theresa Point Airport);Canada
YSK;CYSK;Sanikiluaq (Sanikiluaq Airport);Canada
YSF;CYSF;Stony Rapids (Stony Rapids Airport);Canada
YRL;CYRL;Red Lake (Red Lake Airport);Canada
YRA;CYRA;Gamètì (Rae Lakes Airport);Canada
YQN;CYQN;Nakina (Nakina Airport);Canada
YQD;CYQD;The Pas (The Pas Airport);Canada
YPW;CYPW;Powell River (Powell River Airport);Canada
YPO;CYPO;Peawanuck (Peawanuck Airport);Canada
YPM;CYPM;Pikangikum (Pikangikum Airport);Canada
YPH;CYPH;Inukjuak (Inukjuak Airport);Canada
YOH;CYOH;Oxford House (Oxford House Airport);Canada
YNL;CYNL;Points North Landing (Points North Landing Airport);Canada
YNE;CYNE;Norway House (Norway House Airport);Canada
YNC;CYNC;Wemindji (Wemindji Airport);Canada
YUD;CYMU;Umiujaq (Umiujaq Airport);Canada
YMT;CYMT;Chibougamau (Chapais Airport);Canada
YMH;CYMH;Mary's Harbour (Mary's Harbour Airport);Canada
XGR;CYLU;Kangiqsualujjuaq (Kangiqsualujjuaq (Georges River) Airport);Canada
YSG;CYLK;Lutselk'e (Lutselk'e Airport);Canada
YLH;CYLH;Lansdowne House (Lansdowne House Airport);Canada
YLC;CYLC;Kimmirut (Kimmirut Airport);Canada
YPJ;CYLA;Aupaluk (Aupaluk Airport);Canada
YKQ;CYKQ;Waskaganish (Waskaganish Airport);Canada
AKV;CYKO;Akulivik (Akulivik Airport);Canada
YIV;CYIV;Island Lake (Island Lake Airport);Canada
YIK;CYIK;Ivujivik (Ivujivik Airport);Canada
YHR;CYHR;Chevery (Chevery Airport);Canada
YHO;CYHO;Hopedale (Hopedale Airport);Canada
YNS;CYHH;Nemiscau (Nemiscau Airport);Canada
YHC;CYHC;Vancouver (Vancouver Harbour Water Airport);Canada
YQC;CYHA;Quaqtaq (Quaqtaq Airport);Canada
YGZ;CYGZ;Grise Fiord (Grise Fiord Airport);Canada
YGX;CYGX;Gillam (Gillam Airport);Canada
YGW;CYGW;Kuujjuarapik (Kuujjuarapik Airport);Canada
YGT;CYGT;Igloolik (Igloolik Airport);Canada
YGO;CYGO;Gods Lake Narrows (Gods Lake Narrows Airport);Canada
YGB;CYGB;Texada (Texada Gillies Bay Airport);Canada
YMN;CYFT;Makkovik (Makkovik Airport);Canada
YFH;CYFH;Fort Hope (Fort Hope Airport);Canada
YFA;CYFA;Fort Albany (Fort Albany Airport);Canada
YER;CYER;Fort Severn (Fort Severn Airport);Canada
YDP;CYDP;Nain (Nain Airport);Canada
YCS;CYCS;Chesterfield Inlet (Chesterfield Inlet Airport);Canada
YRF;CYCA;Cartwright (Cartwright Airport);Canada
YBX;CYBX;Lourdes-De-Blanc-Sablon (Lourdes De Blanc Sablon Airport);Canada
YBE;CYBE;Uranium City (Uranium City Airport);Canada
YAT;CYAT;Attawapiskat (Attawapiskat Airport);Canada
YKG;CYAS;Kangirsuk (Kangirsuk Airport);Canada
XKS;CYAQ;Kasabonika (Kasabonika Airport);Canada
YAG;CYAG;Fort Frances (Fort Frances Municipal Airport);Canada
YAC;CYAC;Cat Lake (Cat Lake Airport);Canada
ZLT;CTU5;La Tabatière (La Tabatière Airport);Canada
ZTB;CTB6;Tête-à-la-Baleine (Tête-à-la-Baleine Airport);Canada
YKU;CSU2;Chisasibi (Chisasibi Airport);Canada
YHP;CPV7;Poplar Hill (Poplar Hill Airport);Canada
YOG;CNT3;Ogoki Post (Ogoki Post Airport);Canada
KIF;CNM5;Kingfisher Lake (Kingfisher Lake Airport);Canada
XBE;CNE3;Bearskin Lake (Bearskin Lake Airport);Canada
YNO;CKQ3;North Spirit Lake (North Spirit Lake Airport);Canada
WNN;CKL3;Wunnumin Lake (Wunnumin Lake Airport);Canada
YAX;CKB6;Angling Lake (Wapekeka Airport);Canada
SUR;CJV7;Summer Beaver (Summer Beaver Airport);Canada
YLE;CEM3;Whatì (Whatì Airport);Canada
YCK;CEB3;Colville Lake (Colville Lake Airport);Canada
YRG;CCZ2;Rigolet (Rigolet Airport);Canada
YHA;CCP4;Port Hope Simpson (Port Hope Simpson Airport);Canada
YFX;CCK4;St. Lewis (St. Lewis (Fox Harbour) Airport);Canada
YWM;CCA6;Williams Harbour (Williams Harbour Airport);Canada
YAA;CAJ4;Anahim Lake (Anahim Lake Airport);Canada
YWS;CAE5;Whistler (Whistler/Green Lake Water Aerodrome);Canada
VPN;BIVO;Vopnafjörður (Vopnafjörður Airport);Iceland
THO;BITN;Thorshofn (Thorshofn Airport);Iceland
DRK;MRDK;Puntarenas (Drake Bay Airport);Costa Rica
GRY;BIGR;Grímsey (Grímsey Airport);Iceland
JQA;BGUQ;Uummannaq (Qaarsut Airport);Greenland
JUV;BGUK;Upernavik (Upernavik Airport);Greenland
JHS;BGSS;Sisimiut (Sisimiut Airport);Greenland
NAQ;BGQQ;Qaanaaq (Qaanaaq Airport);Greenland
JNS;BGNS;Narsaq (Narsaq Heliport);Greenland
JNN;BGNN;Nanortalik (Nanortalik Heliport);Greenland
JSU;BGMQ;Maniitsoq (Maniitsoq Airport);Greenland
JJU;BGJH;Qaqortoq (Qaqortoq Heliport);Greenland
JGO;BGGN;Qeqertarsuaq Airport (Qeqertarsuaq Heliport);Greenland
JFR;BGFH;Paamiut (Paamiut Heliport);Greenland
CNP;BGCO;Neerlerit Inaat (Neerlerit Inaat Airport);Greenland
LLU;BGAP;Alluitsup Paa (Alluitsup Paa Heliport);Greenland
WBM;AYWD;Wapenamanda (Wapenamanda Airport);Papua New Guinea
VAI;AYVN;Vanimo (Vanimo Airport);Papua New Guinea
RAB;AYTK;Tokua (Tokua Airport);Papua New Guinea
TBG;AYTB;Tabubil (Tabubil Airport);Papua New Guinea
TIZ;AYTA;Tari (Tari Airport);Papua New Guinea
MIS;AYMS;Misima Island (Misima Island Airport);Papua New Guinea
MXH;AYMR;Moro (Moro Airport);Papua New Guinea
MAS;AYMO;Momote (Momote Airport);Papua New Guinea
MDU;AYMN;Mendi (Mendi Airport);Papua New Guinea
KVG;AYKV;Kavieng (Kavieng Airport);Papua New Guinea
KMA;AYKM;Kerema (Kerema Airport);Papua New Guinea
KRI;AYKK;Kikori (Kikori Airport);Papua New Guinea
UNG;AYKI;Kiunga (Kiunga Airport);Papua New Guinea
HKN;AYHK;Hoskins (Kimbe Airport);Papua New Guinea
PNP;AYGR;Girua (Girua Airport);Papua New Guinea
GUR;AYGN;Gurney (Gurney Airport);Papua New Guinea
DAU;AYDU;Daru (Daru Airport);Papua New Guinea
CMU;AYCH;Kundiawa (Chimbu Airport);Papua New Guinea
BUA;AYBK;Buka Island (Buka Airport);Papua New Guinea
RBV;AGRM;Ramata (Ramata Airport);Solomon Islands
KGE;AGKG;Kagau Island (Kagau Island Airport);Solomon Islands
VAO;AGGV;Suavanao (Suavanao Airport);Solomon Islands
RUS;AGGU;Marau (Marau Airport);Solomon Islands
RNL;AGGR;Rennell Island (Rennell/Tingoa Airport);Solomon Islands
MNY;AGGO;Stirling Island (Mono Airport);Solomon Islands
GZO;AGGN;Gizo (Nusatupe Airport);Solomon Islands
MUA;AGGM;Munda (Munda Airport);Solomon Islands
SCZ;AGGL;Santa Cruz/Graciosa Bay/Luova (Santa Cruz/Graciosa Bay/Luova Airport);Solomon Islands
IRA;AGGK;Kirakira (Ngorangora Airport);Solomon Islands
MBU;AGGI;Mbambanakira (Babanakira Airport);Solomon Islands
FRE;AGGF;Fera Island (Fera/Maringe Airport);Solomon Islands
BAS;AGGE;Ballalae (Ballalae Airport);Solomon Islands
AKS;AGGA;Auki (Auki Airport);Solomon Islands
ATD;AGAT;Atoifi (Uru Harbour Airport);Solomon Islands
KXK;UHKK;Komsomolsk-on-Amur (Komsomolsk-on-Amur Airport);Russia
IKS;UEST;Tiksi (Tiksi Airport);Russia
CYX;UESS;Cherskiy (Cherskiy Airport);Russia
CKH;UESO;Chokurdah (Chokurdakh Airport);Russia
CNN;UELL;Neryungri (Chulman);Russia
KSN;UAUU;Kostanay (Kostanay West Airport);Kazakhstan
DZN;UAKD;Zhezkazgan (Zhezkazgan Airport);Kazakhstan
KOV;UACK;Kokshetau (Kokshetau Airport);Kazakhstan
UNI;TVSU;Union Island (Union Island International Airport);Saint Vincent and the Grenadines
BQU;TVSB;Bequia (J F Mitchell Airport);Saint Vincent and the Grenadines
VIJ;TUPW;Spanish Town (Virgin Gorda Airport);British Virgin Islands
NEV;TKPN;Charlestown (Vance Winkworth Amory International Airport);Saint Kitts and Nevis
BBR;TFFB;Basse Terre (Baillif Airport);Guadeloupe
DSD;TFFA;Grande Anse (La Désirade Airport);Guadeloupe
VIG;SVVG;El Vigía (Juan Pablo Pérez Alfonso Airport);Venezuela
MDO;SUPE;Maldonado (El Jaguel / Punta del Este Airport);Uruguay
SRA;SSZR;Santa Rosa (Santa Rosa Airport);Brazil
HUU;SPNC;Huánuco (Alferez Fap David Figueroa Fernandini Airport);Peru
LHC;SPBC;Caballococha (Caballococha Airport);Peru
SRJ;SLSB;San Borja (Capitán Av. German Quiroga G. Airport);Bolivia
RIB;SLRI;Riberalta (Capitán Av. Selin Zeitun Lopez Airport);Bolivia
GYA;SLGY;Guayaramerín (Capitán de Av. Emilio Beltrán Airport);Bolivia
PDA;SKPD;Puerto Inírida (Obando Airport);Colombia
NQU;SKNQ;Nuquí (Reyes Murillo Airport);Colombia
LPD;SKLP;La Pedrera (La Pedrera Airport);Colombia
LQM;SKLG;Puerto Leguízamo (Caucaya Airport);Colombia
GLJ;SKGZ;Garzón (La Jagua Airport);Colombia
CRC;SKGO;Cartago (Santa Ana Airport);Colombia
PSY;SFAL;Stanley (Stanley Airport);Falkland Islands
LOH;SETM;La Toma (Catamayo) (Camilo Ponce Enriquez Airport);Ecuador
SCY;SEST;San Cristóbal (San Cristóbal Airport);Ecuador
SOD;SDCO;Sorocaba (Sorocaba Airport);Brazil
ESR;SCES;El Salvador (Ricardo García Posada Airport);Chile
VDC;SBQV;Vitória Da Conquista (Vitória da Conquista Airport);Brazil
MII;SBML;Marília (Marília Airport);Brazil
MEA;SBME;Macaé (Macaé Airport);Brazil
NEC;SAZO;Necochea (Necochea Airport);Argentina
ING;SAWA;El Calafate (Lago Argentino Airport);Argentina
RXS;RPVR;Roxas City (Roxas Airport);Philippines
CYP;RPVC;Calbayog City (Calbayog Airport);Philippines
VRC;RPUV;Virac (Virac Airport);Philippines
TUG;RPUT;Tuguegarao (Tuguegarao Airport);Philippines
SFE;RPUS;San Fernando (San Fernando Airport);Philippines
BSO;RPUO;Basco (Basco Airport);Philippines
WNP;RPUN;Naga (Naga Airport);Philippines
TDG;RPMW;Tandag (Tandag Airport);Philippines
SUG;RPMS;Sangley Point (Surigao Airport);Philippines
SGS;RPMN;Sanga Sanga (Sanga Sanga Airport);Philippines
GES;RPMB;General Santos City (General Santos International Airport);Philippines
SFS;RPLB;Olongapo City (Subic Bay International Airport);Philippines
YNY;RKNY;Sokcho / Gangneung (Yangyang International Airport);South Korea
DYR;UHMA;Anadyr (Ugolny Airport);Russia
OHO;UHOO;Okhotsk (Okhotsk Airport);Russia
UJE;UJAP;Ujae Atoll (Ujae Atoll Airport);Marshall Islands
MPW;UKCM;Mariupol International (Mariupol International Airport);Ukraine
VSG;UKCW;Lugansk (Luhansk International Airport);Ukraine
OZH;UKDE;Zaporozhye (Zaporizhzhia International Airport);Ukraine
KWG;UKDR;Krivoy Rog (Lozuvatka International Airport);Ukraine
HRK;UKHH;Kharkov (Osnova International Airport);Ukraine
IFO;UKLI;Ivano-Frankivsk (Ivano Frankivsk International Airport);Ukraine
CWC;UKLN;Chernovtsk (Chernivtsi International Airport);Ukraine
RWN;UKLR;Rivne (Rivne International Airport);Ukraine
UDJ;UKLU;Uzhgorod (Uzhhorod International Airport);Ukraine
CSH;ULAS;Solovetsky Islands (Solovki Airport);Russia
CEE;ULBC;Cherepovets (Cherepovets Airport);Russia
AMV;ULDD;Amderma (Amderma Airport);Russia
KSZ;ULKK;Kotlas (Kotlas Airport);Russia
PES;ULPB;Petrozavodsk (Petrozavodsk Airport);Russia
GNA;UMMG;Hrodna (Hrodno Airport);Belarus
MVQ;UMOO;Mogilev (Mogilev Airport);Belarus
EIE;UNII;Yeniseysk (Yeniseysk Airport);Russia
KYZ;UNKY;Kyzyl (Kyzyl Airport);Russia
NOZ;UNWW;Novokuznetsk (Spichenkovo Airport);Russia
HTG;UOHH;Khatanga (Khatanga Airport);Russia
IAA;UOII;Igarka (Igarka Airport);Russia
GRV;URMG;Grozny (Grozny Airport);Russia
NAL;URMN;Nalchik (Nalchik Airport);Russia
OGZ;URMO;Beslan (Beslan Airport);Russia
ESL;URWI;Elista (Elista Airport);Russia
WKK;5A8;Aleknagik (Aleknagik Airport);United States
BKX;BKX;Brookings (Brookings Regional Airport);United States
BLF;BLF;Bluefield (Mercer County Airport);United States
EAR;EAR;Kearney (Kearney Municipal Airport);United States
GLH;GLH;Greenville (Mid Delta Regional Airport);United States
IFP;IFP;Bullhead (Laughlin-Bullhead Intl);United States
IGM;IGM;Kingman (Kingman Airport);United States
PSC;PSC;Pasco (Tri Cities Airport);United States
KQA;KQA;Akutan (Akutan Seaplane Base);United States
SVC;SVC;Silver City (Grant County Airport);United States
LPS;S31;Lopez (Lopez Island Airport);United States
SLY;USDD;Salekhard (Salekhard Airport);Russia
HMA;USHH;Khanty-Mansiysk (Khanty Mansiysk Airport);Russia
NYA;USHN;Nyagan (Nyagan Airport);Russia
OVS;USHS;Sovetskiy (Sovetsky Tyumenskaya Airport);Russia
IJK;USII;Izhevsk (Izhevsk Airport);Russia
KVX;USKK;Kirov (Pobedilovo Airport);Russia
NYM;USMM;Nadym (Nadym Airport);Russia
RAT;USNR;Raduzhnyi (Raduzhny Airport);Russia
NFG;USRN;Nefteyugansk (Nefteyugansk Airport);Russia
KRO;USUU;Kurgan (Kurgan Airport);Russia
LBD;UTDL;Khudzhand (Khudzhand Airport);Tajikistan
AZN;UTKA;Andizhan (Andizhan Airport);Uzbekistan
FEG;UTKF;Fergana (Fergana Airport);Uzbekistan
NMA;UTKN;Namangan (Namangan Airport);Uzbekistan
NCU;UTNN;Nukus (Nukus Airport);Uzbekistan
UGC;UTNU;Urgench (Urgench Airport);Uzbekistan
KSQ;UTSL;Khanabad (Karshi Khanabad Airport);Uzbekistan
TMJ;UTST;Termez (Termez Airport);Uzbekistan
RYB;UUBK;Rybinsk (Staroselye Airport);Russia
EGO;UUOB;Belgorod (Belgorod International Airport);Russia
URS;UUOK;Kursk (Kursk East Airport);Russia
LPK;UUOL;Lipetsk (Lipetsk Airport);Russia
VKT;UUYW;Vorkuta (Vorkuta Airport);Russia
UUA;UWKB;Bugulma (Bugulma Airport);Russia
JOK;UWKJ;Yoshkar-Ola (Yoshkar-Ola Airport);Russia
CSY;UWKS;Cheboksary (Cheboksary Airport);Russia
ULY;UWLW;Ulyanovsk (Ulyanovsk East Airport);Russia
OSW;UWOR;Orsk (Orsk Airport);Russia
PEZ;UWPP;Penza (Penza Airport);Russia
SKX;UWPS;Saransk (Saransk Airport);Russia
BWO;UWSB;Balakovo (Balakovo Airport);Russia
HBX;VAHB;Hubli (Hubli Airport);India
KCT;VCCK;Koggala (Koggala Airport);Sri Lanka
WRZ;VCCW;Wirawila (Wirawila Airport);Sri Lanka
BBM;VDBG;Battambang (Battambang Airport);Cambodia
SHL;VEBI;Shillong (Shillong Airport);India
GAU;VEGT;Guwahati (Lokpriya Gopinath Bordoloi International Airport);India
DMU;VEMR;Dimapur (Dimapur Airport);India
TEZ;VETZ;Tezpur (Tezpur Airport);India
BZL;VGBR;Barisal (Barisal Airport);Bangladesh
OUI;VLHS;Huay Xai (Ban Huoeisay Airport);Laos
BHR;VNBP;Bharatpur (Bharatpur Airport);Nepal
BDP;VNCG;Chandragarhi (Chandragadhi Airport);Nepal
MEY;VNMG;Meghauli (Meghauli Airport);Nepal
KEP;VNNG;Nepalgunj (Nepalgunj Airport);Nepal
GAN;VRMG;Gan Island (Gan Island Airport);Maldives
HAQ;VRMH;Haa Dhaalu Atoll (Hanimaadhoo Airport);Maldives
KDO;VRMK;Laamu Atoll (Kadhdhoo Airport);Maldives
MAQ;VTPM;Tak (Mae Sot Airport);Thailand
BMV;VVBM;Buonmethuot (Buon Ma Thuot Airport);Vietnam
HPH;VVCI;Haiphong (Cat Bi International Airport);Vietnam
CXR;VVCR;Nha Trang (Cam Ranh Airport);Vietnam
VCS;VVCS;Conson (Co Ong Airport);Vietnam
VCA;VVCT;Can Tho (Trà Nóc Airport);Vietnam
DIN;VVDB;Dienbienphu (Dien Bien Phu Airport);Vietnam
UIH;VVPC;Phucat (Phu Cat Airport);Vietnam
PXU;VVPK;Pleiku (Pleiku Airport);Vietnam
VII;VVVH;Vinh (Vinh Airport);Vietnam
BMO;VYBM;Banmaw (Banmaw Airport);Burma
TVY;VYDW;Dawei (Dawei Airport);Burma
KAW;VYKT;Kawthoung (Kawthoung Airport);Burma
LIW;VYLK;Loikaw (Loikaw Airport);Burma
MNU;VYMM;Mawlamyine (Mawlamyine Airport);Burma
BSX;VYPN;Pathein (Pathein Airport);Burma
PKK;VYPU;Pakhokku (Pakhokku Airport);Burma
SWQ;WADS;Sumbawa Island (Sumbawa Besar Airport);Indonesia
TMC;WADT;Waikabubak-Sumba Island (Tambolaka Airport);Indonesia
BUI;WAJB;Bokondini-Papua Island (Bokondini Airport);Indonesia
SEH;WAJS;Senggeh-Papua Island (Senggeh Airport);Indonesia
TJS;WALG;Tanjung Selor-Borneo Island (Tanjung Harapan Airport);Indonesia
DTD;WALJ;Datadawai-Borneo Island (Datadawai Airport);Indonesia
BEJ;WALK;Tanjung Redep-Borneo Island (Barau(Kalimaru) Airport);Indonesia
TJG;WAON;Tanjung-Borneo Island (Warukin Airport);Indonesia
SMQ;WAOS;Sampit-Borneo Island (Sampit(Hasan) Airport);Indonesia
LUV;WAPL;Langgur-Kei Islands (Dumatubun Airport);Indonesia
ARD;WATM;Alor Island (Mali Airport);Indonesia
BLG;WBGC;Belaga (Belaga Airport);Malaysia
LGL;WBGF;Long Datih (Long Lellang Airport);Malaysia
ODN;WBGI;Long Seridan (Long Seridan Airport);Malaysia
MKM;WBGK;Mukah (Mukah Airport);Malaysia
BKM;WBGQ;Bakalalan (Bakalalan Airport);Malaysia
LWY;WBGW;Lawas (Lawas Airport);Malaysia
BBN;WBGZ;Bario (Bario Airport);Malaysia
TMG;WBKM;Tomanggong (Tomanggong Airport);Malaysia
KUD;WBKT;Kudat (Kudat Airport);Malaysia
TKG;WICT;Bandar Lampung-Sumatra Island (Radin Inten II (Branti) Airport);Indonesia
HLP;WIHH;Jakarta (Halim Perdanakusuma International Airport);Indonesia
NTX;WION;Ranai-Natuna Besar Island (Ranai Airport);Indonesia
PSU;WIOP;Putussibau-Borneo Island (Pangsuma Airport);Indonesia
SQG;WIOS;Sintang-Borneo Island (Susilo Airport);Indonesia
PDO;WIPQ;Talang Gudang-Sumatra Island (Pendopo Airport);Indonesia
LSW;WITM;Lhok Seumawe-Sumatra Island (Malikus Saleh Airport);Indonesia
PKG;WMPA;Pangkor Island (Pulau Pangkor Airport);Malaysia
KBU;WRBK;Laut Island (Stagen Airport);Indonesia
LBW;WRLB;Long Bawan-Borneo Island (Long Bawan Airport);Indonesia
NNX;WRLF;Nunukan-Nunukan Island (Nunukan Airport);Indonesia
LPU;WRLP;Long Apung-Borneo Island (Long Apung Airport);Indonesia
ALH;YABA;Albany (Albany Airport);Australia
GYL;YARG;Argyle (Argyle Airport);Australia
AUU;YAUR;Aurukun (Aurukun Airport);Australia
BCI;YBAR;Barcaldine (Barcaldine Airport);Australia
BDD;YBAU;Badu Island (Badu Island Airport);Australia
BVI;YBDV;Birdsville (Birdsville Airport);Australia
BHQ;YBHI;Broken Hill (Broken Hill Airport);Australia
HTI;YBHM;Hamilton Island (Hamilton Island Airport);Australia
BEU;YBIE;Bedourie (Bedourie Airport);Australia
BRK;YBKE;Bourke (Bourke Airport);Australia
BUC;YBKT;Burketown (Burketown Airport);Australia
GIC;YBOI;Boigu (Boigu Airport);Australia
OKY;YBOK;Oakey (Oakey Airport);Australia
BQL;YBOU;Boulia (Boulia Airport);Australia
BHS;YBTH;Bathurst (Bathurst Airport);Australia
BLT;YBTR;Blackwater (Blackwater Airport);Australia
CVQ;YCAR;Carnarvon (Carnarvon Airport);Australia
CAZ;YCBA;Cobar (Cobar Airport);Australia
CPD;YCBP;Coober Pedy (Coober Pedy Airport);Australia
CNC;YCCT;Coconut Island (Coconut Island Airport);Australia
CNJ;YCCY;Cloncurry (Cloncurry Airport);Australia
CED;YCDU;Ceduna (Ceduna Airport);Australia
CTN;YCKN;Cooktown (Cooktown Airport);Australia
CMA;YCMU;Cunnamulla (Cunnamulla Airport);Australia
CNB;YCNM;Coonamble (Coonamble Airport);Australia
CUQ;YCOE;Coen (Coen Airport);Australia
OOM;YCOM;Cooma (Cooma Snowy Mountains Airport);Australia
DMD;YDMG;Doomadgee (Doomadgee Airport);Australia
NLF;YDNI;Darnley Island (Darnley Island Airport);Australia
DPO;YDPO;Devonport (Devonport Airport);Australia
ELC;YELD;Elcho Island (Elcho Island Airport);Australia
EPR;YESP;Esperance (Esperance Airport);Australia
FLS;YFLI;Flinders Island (Flinders Island Airport);Australia
GET;YGEL;Geraldton (Geraldton Airport);Australia
GLT;YGLA;Gladstone (Gladstone Airport);Australia
GTE;YGTE;Groote Eylandt (Groote Eylandt Airport);Australia
GFF;YGTH;Griffith (Griffith Airport);Australia
HID;YHID;Horn Island (Horn Island Airport);Australia
HOK;YHOO;Hooker Creek (Hooker Creek Airport);Australia
MHU;YHOT;Mount Hotham (Mount Hotham Airport);Australia
HGD;YHUG;Hughenden (Hughenden Airport);Australia
JCK;YJLC;Julia Creek (Julia Creek Airport);Australia
KAX;YKBR;Kalbarri (Kalbarri Airport);Australia
KNS;YKII;King Island (King Island Airport);Australia
KFG;YKKG;Kalkgurung (Kalkgurung Airport);Australia
KRB;YKMB;Karumba (Karumba Airport);Australia
KWM;YKOW;Kowanyama (Kowanyama Airport);Australia
KUG;YKUB;Kubin (Kubin Airport);Australia
LNO;YLEO;Leonora (Leonora Airport);Australia
LEL;YLEV;Lake Evella (Lake Evella Airport);Australia
LDH;YLHI;Lord Howe Island (Lord Howe Island Airport);Australia
IRG;YLHR;Lockhart River (Lockhart River Airport);Australia
LSY;YLIS;Lismore (Lismore Airport);Australia
LHG;YLRD;Lightning Ridge (Lightning Ridge Airport);Australia
LRE;YLRE;Longreach (Longreach Airport);Australia
LER;YLST;Leinster (Leinster Airport);Australia
LVO;YLTN;Laverton (Laverton Airport);Australia
UBB;YMAA;Mabuiag Island (Mabuiag Island Airport);Australia
MKR;YMEK;Meekatharra (Meekatharra Airport);Australia
MIM;YMER;Merimbula (Merimbula Airport);Australia
MGT;YMGB;Milingimbi (Milingimbi Airport);Australia
MNG;YMGD;Maningrida (Maningrida Airport);Australia
MCV;YMHU;McArthur River Mine (McArthur River Mine Airport);Australia
MQL;YMIA;Mildura (Mildura Airport);Australia
MMG;YMOG;Mount Magnet (Mount Magnet Airport);Australia
MRZ;YMOR;Moree (Moree Airport);Australia
MOV;YMRB;Moranbah (Moranbah Airport);Australia
MYA;YMRY;Moruya (Moruya Airport);Australia
MGB;YMTG;Mount Gambier (Mount Gambier Airport);Australia
ONG;YMTI;Mornington Island (Mornington Island Airport);Australia
MYI;YMUI;Murray Island (Murray Island Airport);Australia
MBH;YMYB;Maryborough (Maryborough Airport);Australia
NRA;YNAR;Narrandera (Narrandera Airport);Australia
NAA;YNBR;Narrabri (Narrabri Airport);Australia
NTN;YNTN;Normanton (Normanton Airport);Australia
ZNE;YNWN;Newman (Newman Airport);Australia
OLP;YOLD;Olympic Dam (Olympic Dam Airport);Australia
PUG;YPAG;Argyle (Port Augusta Airport);Australia
PMK;YPAM;Palm Island (Palm Island Airport);Australia
PBO;YPBO;Paraburdoo (Paraburdoo Airport);Australia
CCK;YPCC;Cocos Keeling Island (Cocos Keeling Island Airport);Cocos (Keeling) Islands
GOV;YPGV;Gove (Gove Airport);Australia
PKE;YPKS;Parkes (Parkes Airport);Australia
PLO;YPLC;Port Lincoln (Port Lincoln Airport);Australia
EDR;YPMP;Pormpuraaw (Pormpuraaw Airport);Australia
PQQ;YPMQ;Port Macquarie (Port Macquarie Airport);Australia
PTJ;YPOD;Portland (Portland Airport);Australia
ULP;YQLP;Quilpie (Quilpie Airport);Australia
RAM;YRNG;Ramingining (Ramingining Airport);Australia
RMA;YROM;Roma (Roma Airport);Australia
SGO;YSGE;St George (St George Airport);Australia
MJK;YSHK;Shark Bay (Shark Bay Airport);Australia
SBR;YSII;Saibai Island (Saibai Island Airport);Australia
SRN;YSRN;Strahan (Strahan Airport);Australia
XTG;YTGM;Thargomindah (Thargomindah Airport);Australia
TCA;YTNK;Tennant Creek (Tennant Creek Airport);Australia
VCD;YVRD;Victoria River Downs (Victoria River Downs Airport);Australia
SYU;YWBS;Sue Islet (Warraber Island Airport);Australia
WNR;YWDH;Windorah (Windorah Airport);Australia
WYA;YWHA;Whyalla (Whyalla Airport);Australia
WUN;YWLU;Wiluna (Wiluna Airport);Australia
WOL;YWOL;Wollongong (Wollongong Airport);Australia
WIN;YWTN;Winton (Winton Airport);Australia
BWT;YWYY;Burnie (Wynyard Airport);Australia
OKR;YYKI;Yorke Island (Yorke Island Airport);Australia
XMY;YYMI;Yam Island (Yam Island Airport);Australia
NAY;ZBBB;Beijing (Beijing Nanyuan Airport);China
CIF;ZBCF;Chifeng (Chifeng Airport);China
CIH;ZBCZ;Changzhi (Changzhi Airport);China
DAT;ZBDT;Datong (Datong Airport);China
HET;ZBHH;Hohhot (Baita Airport);China
BAV;ZBOW;Baotou (Baotou Airport);China
SJW;ZBSJ;Shijiazhuang (Shijiazhuang Daguocun International Airport);China
TGO;ZBTL;Tongliao (Tongliao Airport);China
HLH;ZBUL;Ulanhot (Ulanhot Airport);China
XIL;ZBXH;Xilinhot (Xilinhot Airport);China
BHY;ZGBH;Beihai (Beihai Airport);China
CGD;ZGCD;Changde (Changde Airport);China
DYG;ZGDY;Dayong (Dayong Airport);China
MXZ;ZGMX;Meixian (Meixian Airport);China
ZUH;ZGSD;Zhuhai (Zhuhai Airport);China
LZH;ZGZH;Liuzhou (Bailian Airport);China
ZHA;ZGZJ;Zhanjiang (Zhanjiang Airport);China
ENH;ZHES;Enshi (Enshi Airport);China
NNY;ZHNY;Nanyang (Nanyang Airport);China
XFN;ZHXF;Xiangfan (Xiangfan Airport);China
YIH;ZHYC;Yichang (Yichang Airport);China
AKA;ZLAK;Ankang (Ankang Airport);China
GOQ;ZLGM;Golmud (Golmud Airport);China
HZG;ZLHZ;Hanzhong (Hanzhong Airport);China
IQN;ZLQY;Qingyang (Qingyang Airport);China
XNN;ZLXN;Xining (Xining Caojiabu Airport);China
ENY;ZLYA;Yan'an (Yan'an Airport);China
UYN;ZLYL;Yulin (Yulin Airport);China
AVK;ZMAH;Arvaikheer (Arvaikheer Airport);Mongolia
LTI;ZMAT;Altai (Altai Airport);Mongolia
BYN;ZMBH;Bayankhongor (Bayankhongor Airport);Mongolia
DLZ;ZMDZ;Dalanzadgad (Dalanzadgad Airport);Mongolia
HVD;ZMKD;Khovd (Khovd Airport);Mongolia
MXV;ZMMN;Muren (Muren Airport);Mongolia
DIG;ZPDQ;Shangri-La (Diqing Airport);China
LUM;ZPLX;Luxi (Mangshi Airport);China
SYM;ZPSM;Simao (Simao Airport);China
ZAT;ZPZT;Zhaotong (Zhaotong Airport);China
KOW;ZSGZ;Ganzhou (Ganzhou Airport);China
JDZ;ZSJD;Jingdezhen (Jingdezhen Airport);China
JIU;ZSJJ;Jiujiang (Jiujiang Lushan Airport);China
JUZ;ZSJU;Quzhou (Quzhou Airport);China
LYG;ZSLG;Lianyungang (Lianyungang Airport);China
HYN;ZSLQ;Huangyan (Huangyan Luqiao Airport);China
LYI;ZSLY;Linyi (Shubuling Airport);China
JJN;ZSQZ;Quanzhou (Quanzhou Airport);China
TXN;ZSTX;Huangshan (Tunxi International Airport);China
WEF;ZSWF;Weifang (Weifang Airport);China
WEH;ZSWH;Weihai (Weihai Airport);China
WUX;ZSWX;Wuxi (Wuxi Airport);China
WUS;ZSWY;Wuyishan (Nanping Wuyishan Airport);China
WNZ;ZSWZ;Wenzhou (Wenzhou Yongqiang Airport);China
YNZ;ZSYN;Yancheng (Yancheng Airport);China
YIW;ZSYW;Yiwu (Yiwu Airport);China
HSN;ZSZS;Zhoushan (Zhoushan Airport);China
BPX;ZUBD;Bangda (Qamdo Bangda Airport);China
DAX;ZUDX;Dazhou (Dachuan Airport);China
GYS;ZUGU;Guangyuan (Guangyuan Airport);China
LZO;ZULZ;Luzhou (Luzhou Airport);China
MIG;ZUMY;Mianyang (Mianyang Airport);China
NAO;ZUNC;Nanchong (Nanchong Airport);China
LZY;ZUNZ;Nyingchi (Nyingchi Airport);China
WXN;ZUWX;Wanxian (Wanxian Airport);China
AKU;ZWAK;Aksu (Aksu Airport);China
IQM;ZWCM;Qiemo (Qiemo Airport);China
KCA;ZWKC;Kuqa (Kuqa Airport);China
KRL;ZWKL;Korla (Korla Airport);China
KRY;ZWKM;Karamay (Karamay Airport);China
YIN;ZWYN;Yining (Yining Airport);China
HEK;ZYHE;Heihe (Heihe Airport);China
JMU;ZYJM;Jiamusi (Jiamusi Airport);China
JNZ;ZYJZ;Jinzhou (Jinzhou Airport);China
NDG;ZYQQ;Qiqihar (Qiqihar Sanjiazi Airport);China
YNJ;ZYYJ;Yanji (Yanji Airport);China
WME;YMNE;Mount Keith;Australia
LRV;SVRS;Los Roques (Gran Roque Airport);Venezuela
IOR;EIIM;Inis Mor (Inishmore Airport);Ireland
NNR;EICA;Indreabhan (Connemara Regional Airport);Ireland
GTI;EDCG;Ruegen (Guettin MecklenburgVorpommern Germany);Germany
NBB;USHB;Berezovo;Russia
ORH;KORH;Worcester (Worcester Regional Airport);United States
AQG;ZSAQ;Anqing (Anqing Airport);China
SHP;ZBSH;Qinhuangdao (Shanhaiguan Airport);China
YCU;ZBYC;Yuncheng (Zhangxiao);China
LHW;ZLAN;Lanzhou (Lanzhou Airport);China
JGN;ZLJQ;Jiayuguan (Jiayuguan Airport);China
DDG;ZYDD;Dandong;China
DSN;ZBDS;Dongsheng (Ordos Ejin Horo);China
PZI;ZUZH;Panzhihua;China
PWT;KPWT;Bremerton (Bremerton National);United States
SPW;KSPW;Spencer (Spencer Muni);United States
JEF;KJEF;Jefferson City (Jefferson City Memorial Airport);United States
BLD;KBVU;Boulder City (Boulder City Municipal Airport);United States
UNT;EGPW;Unst (Unst Airport);United Kingdom
PVC;KPVC;Provincetown (Provincetown Muni);United States
LKE;KW55;Seattle (Kenmore Air Harbor Seaplane Base);United States
%u0;%u04;Nazran (Magas);Russia
SBH;TFFJ;Gustavia (Saint Barthelemy);France
SUI;UGSS;Sukhumi (Sukhumi Dranda);Georgia
TBW;UUOT;Tambow;Russia
OBN;EGEO;North Connel (Oban Airport);United Kingdom
FSZ;RJNS;Shizuoka (Mt. Fuji Shizuoka Airport);Japan
ERM;SSER;Erechim (Erechim Airport);Brazil
CVF;LFLJ;Courcheval (Courchevel Airport);France
FUL;KFUL;Fullerton (Fullerton Municipal Airport);United States
NVI;UTSA;Navoi (Navoi Airport);Uzbekistan
QSF;DAAS;Setif (Ain Arnat Airport);Algeria
LRH;LFBH;La Rochelle (La Rochelle-Ile de Re);France
SUN;KSUN;Hailey (Friedman Mem);United States
MCW;KMCW;Mason City (Mason City Municipal);United States
IAR;UUDL;Yaroslavl (Tunoshna);Russia
QNC;LSGN;Neuchatel (Neuchatel Airport);Switzerland
ZJI;LSZL;Locarno (Locarno Airport);Switzerland
RMT;NTAM;Rimatara;French Polynesia
UKX;UITT;Ust-Kut;Russia
RIN;AGRC;Ringi Cove (Ringi Cove Airport);Solomon Islands
ARE;TJAB;Arecibo (Antonio Juarbe Pol Airport);Puerto Rico
EAT;KEAT;Wenatchee (Pangborn Field);United States
AGE;EDWG;Wangerooge (Wangerooge Airport);Germany
OSH;KOSH;Oshkosh (Wittman Regional Airport);United States
BQT;UMBB;Brest;Belarus
TNL;UKLT;Ternopol;Ukraine
CEJ;UKRR;Chernigov;Ukraine
UKC;UKLC;Lutsk;Ukraine
NOA;YSNW;Nowra (Nowra Airport);Australia
KTR;YPTN;Katherine (Tindal Airport);Australia
COQ;ZMCD;Choibalsan (Choibalsan Airport);Mongolia
TRO;YTRE;Taree (Taree Airport);Australia
OAG;YORG;Orange (Orange Airport);Australia
GFN;YGFN;Grafton (Grafton Airport);Australia
MRQ;RPUW;Gasan (Marinduque Airport);Philippines
HDM;OIHH;Hamadan (Hamadan Airport);Iran
YIF;CYIF;St-Augustin (St Augustin Airport);Canada
VQS;TJCG;Vieques Island (Vieques Airport);Puerto Rico
KMV;VYKL;Kalemyo (Kalay Airport);Myanmar
LSS;TFFS;Les Saintes (Terre-de-Haut Airport);Guadeloupe
YEI;LTBR;Yenisehir (Yenisehir Airport);Turkey
TEQ;LTBU;Çorlu (Tekirdağ Çorlu Airport);Turkey
SIC;LTCM;Sinop (Sinop Airport);Turkey
MSR;LTCK;Mus (Mus Airport);Turkey
CKZ;LTBH;Canakkale (Canakkale Airport);Turkey
AOE;LTBY;Eskissehir (Anadolu Airport);Turkey
MPA;FYKM;Mpacha (Katima Mulilo Airport);Namibia
WVB;FYWB;Walvis Bay (Walvis Bay Airport);Namibia
PDP;SULS;Punta del Este (Capitan Corbeta C A Curbelo International Airport);Uruguay
SKT;OPST;Sialkot (Sialkot Airport);Pakistan
YVB;CYVB;Bonaventure (Bonaventure Airport);Canada
BHG;MHBL;Brus Laguna (Brus Laguna Airport);Honduras
UAS;HKSB;Samburu South (Samburu South Airport);Kenya
CHG;ZYCY;Chaoyang (Chaoyang Airport);China
WLH;NVSW;Walaha (Walaha Airport);Vanuatu
TGC;WBTM;Tanjung Manis (Tanjung Manis Airport);Malaysia
LKH;WBGL;Long Akah (Long Akah Airport);Malaysia
EGN;HSGN;Geneina (Geneina Airport);Sudan
TOG;PATG;Togiak Village (Togiak Airport);United States
PTH;PAPH;Port Heiden (Port Heiden Airport);United States
KVC;PAVC;King Cove (King Cove Airport);United States
KNW;PANW;New Stuyahok (New Stuyahok Airport);United States
IGG;PAIG;Igiugig (Igiugig Airport);United States
SLV;VISM;Shimla (Shimla Airport);India
NDC;VAND;Nanded (Nanded Airport);India
DHM;VIGG;Kangra (Kangra Airport);India
CQD;OIFS;Shahre Kord (Shahre Kord Airport);Iran
EGM;AGGS;Sege (Sege Airport);Solomon Islands
RGS;LEBG;Burgos (Burgos Airport);Spain
LEN;LELN;Leon (Leon Airport);Spain
DRG;PADE;Deering (Deering Airport);United States
AFS;UTSN;Zarafshan (Sugraly Airport);Uzbekistan
MQM;LTCR;Mardin (Mardin Airport);Turkey
TCG;ZWTC;Tacheng (Tacheng Airport);China
LGQ;SELA;Lago Agrio (Nueva Loja Airport);Ecuador
PFQ;OITP;Parsabad (Parsabade Moghan Airport);Iran
IIL;OICI;Ilam (Ilam Airport);Iran
GBT;OING;Gorgan (Gorgan Airport);Iran
ACP;OITM;Maragheh (Sahand Airport);Iran
TBH;RPVU;Romblon (Romblon Airport);Philippines
WUZ;ZGWZ;Wuzhou (Changzhoudao Airport);China
HMI;ZWHM;Hami (Hami Airport);China
SDP;PASD;Sand Point (Sand Point Airport);United States
GOP;VEGK;Gorakhpur (Gorakhpur Airport);India
ACR;SKAC;Araracuara (Araracuara Airport);Colombia
HGR;KHGR;Hagerstown (Hagerstown Regional Richard A Henson Field);United States
QBC;CYBD;Bella Coola (Bella Coola Airport);Canada
PJA;ESUP;Pajala (Pajala Airport);Sweden
KPC;PAPC;Port Clarence (Port Clarence Coast Guard Station);United States
GVR;SBGV;Governador Valadares (Governador Valadares Airport);Brazil
KVK;ULMK;Apatity (Kirovsk-Apatity Airport);Russia
CYZ;RPUY;Cauayan (Cauayan Airport);Philippines
TMU;MRTR;Nicoya (Tambor Airport);Costa Rica
FON;MRAN;La Fortuna/San Carlos (Arenal Airport);Costa Rica
QOW;DNIM;Imo (Imo Airport);Nigeria
ARC;PARC;Arctic Village (Arctic Village Airport);United States
YTQ;CYTQ;Tasiujaq (Tasiujaq Airport);Canada
YPX;CYPX;Puvirnituq (Puvirnituq Airport);Canada
OMC;RPVO;Ormoc City (Ormoc Airport);Philippines
WTK;PAWN;Noatak (Noatak Airport);United States
SVA;PASA;Savoonga (Savoonga Airport);United States
SHH;PASH;Shishmaref (Shishmaref Airport);United States
RBY;PARY;Ruby (Ruby Airport);United States
PHO;PPHO;Point Hope (Point Hope Airport);United States
MYU;PAMY;Mekoryuk (Mekoryuk Airport);United States
KVL;PAVL;Kivalina (Kivalina Airport);United States
KSM;PASM;St Mary's (St Marys Airport);United States
KAL;PAKV;Kaltag (Kaltag Airport);United States
HPB;PAHP;Hooper Bay (Hooper Bay Airport);United States
GAM;PAGM;Gambell (Gambell Airport);United States
ATK;PATQ;Atqasuk (Atqasuk Edward Burnell Sr Memorial Airport);United States
ANV;PANV;Anvik (Anvik Airport);United States
AKP;PAKP;Anaktuvuk Pass (Anaktuvuk Pass Airport);United States
AAT;ZWAT;Altay (Altay Airport);China
ELA;VYEL;Naypyidaw;Burma
ZXB;ENJA;Jan Mayen (Jan Mayensfield);Norway
WUA;ZBUH;Wuhai;China
GYY;KGYY;Gary (Gary Chicago International Airport);United States
BRD;KBRD;Brainerd (Brainerd Lakes Rgnl);United States
LWB;KLWB;Lewisburg (Greenbrier Valley Airport);United States
PGV;KPGV;Greenville (Pitt-Greenville Airport);United States
CYF;PACK;Chefornak (Chefornak Airport);United States
OXR;KOXR;Oxnard (Oxnard - Ventura County);United States
BKG;KBBG;Branson (Branson LLC);United States
TEN;ZUTR;Tongren;China
KNC;ZSJA;Jian (Jinggangshan);China
NIU;NTKN;Niau;French Polynesia
SCH;KSCH;Scotia NY (Stratton ANGB - Schenectady County Airpor);United States
NBC;UWKE;Nizhnekamsk (Begishevo);Russia
QRW;DNSU;Osubi (Warri Airport);Nigeria
LGO;EDWL;Langeoog (Langeoog Airport);Germany
NLP;FANS;Nelspruit (Nelspruit Airport);South Africa
CKC;UKKE;Cherkassy;Ukraine
UST;KSGJ;St. Augustine Airport;United States
NLV;UKON;Nikolayev (Mykolaiv International Airport);Ukraine
RHP;VNRC;Ramechhap;Nepal
STS;KSTS;Santa Rosa (Charles M Schulz Sonoma Co);United States
ISM;KISM;Kissimmee (Kissimmee Gateway Airport);United States
LCQ;KLCQ;Lake City (Lake City Municipal Airport);United States
LGU;KLGU;Logan (Logan-Cache);United States
BMC;KBMC;Brigham City;United States
MLD;KMLD;Malad City;United States
ASE;KASE;Aspen (Aspen Pitkin County Sardy Field);United States
HHH;KHHH;Hilton Head;United States
ULV;UWLL;Ulyanovsk (Barataevka);Russia
ERV;KERV;Kerrville (Kerrville Municipal Airport);United States
GED;KGED;Georgetown (Sussex Co);United States
ZSW;CZSW;Prince Rupert (Seal Cove Seaplane Base);Canada
GBN;KGBD;Great Bend (Great Bend Municipal);United States
HYS;KHYS;Hays (Hays Regional Airport);United States
SUS;KSUS;Null (Spirit Of St Louis);United States
LYU;KELO;Ely (Ely Municipal);United States
GPZ;KGPZ;Grand Rapids MN (Grand Rapids Itasca County);United States
TVF;KTVF;Thief River Falls;United States
EGV;KEGV;Eagle River;United States
ARV;KARV;Minocqua - Woodruff (Lakeland);United States
IKV;KIKV;Ankeny (Ankeny Regl Airport);United States
YBV;CYBV;Berens River;Canada
NGP;KNGP;Corpus Christi (Corpus Christi NAS);United States
YPI;CYPI;Port Simpson (Seaplane Base);Canada
AVX;KAVX;Catalina Island (Avalon);United States
MHV;KMHV;Mojave;United States
ZIN;LSMI;Interlaken (Air Base);Switzerland
INQ;EIIR;Inisheer;Ireland
SWT;UNSS;Strezhevoy;Russia
HUT;KHUT;Hutchinson (Hutchinson Municipal Airport);United States
BPM;OAIX;Kabul (Bagram AFB);Afghanistan
STJ;KSTJ;Rosecrans (Rosecrans Mem);United States
NDZ;KNDZ;Cuxhaven (Cuxhaven Airport);Germany
VOK;KVOK;Camp Douglas (Volk Fld);United States
BFT;KBFT;Beauford (BFT County Airport);United States
UAB;KUAB;Adana (Adana-Incirlik Airbase);Turkey
GUC;KGUC;Gunnison (Gunnison - Crested Butte);United States
SIA;ZLSN;Xi'AN (Xi'An Xiguan);China
TOA;KTOA;Torrance (Zamperini Field Airport);United States
MBL;KMBL;Manistee (Manistee County-Blacker Airport);United States
PGD;KPGD;Punta Gorda (Charlotte County-Punta Gorda Airport);United States
WFK;KFVE;Frenchville (Northern Aroostook Regional Airport);United States
JHW;KJHW;Jamestown (Chautauqua County-Jamestown);United States
YTM;CYFJ;Mont-Tremblant (Riviere Rouge - Mont-Tremblant International Inc. Airport);Canada
SME;KSME;Somerset (Lake Cumberland Regional Airport);United States
SHD;KSHD;Weyers Cave (Shenandoah Valley Regional Airport);United States
DVL;KDVL;Devils Lake (Devils Lake Regional Airport);United States
DIK;KDIK;Dickinson (Dickinson Theodore Roosevelt Regional Airport);United States
SDY;KSDY;Sidney (Sidney-Richland Municipal Airport);United States
CDR;KCDR;Chadron (Chadron Municipal Airport);United States
AIA;KAIA;Alliance (Alliance Municipal Airport);United States
MCK;KMCK;McCook (McCook Regional Airport);United States
MTH;KMTH;Marathon (Florida Keys Marathon Airport);United States
GDV;KGDV;Glendive (Dawson Community Airport);United States
OLF;KOLF;Wolf Point (LM Clayton Airport);United States
WYS;KWYS;West Yellowstone (Yellowstone Airport);United States
ALS;KALS;Alamosa (San Luis Valley Regional Airport);United States
CNY;KCNY;Moab (Canyonlands Field);United States
ELY;KELY;Ely (Ely Airport);United States
VEL;KVEL;Vernal (Vernal Regional Airport);United States
SRR;KSRR;Ruidoso (Sierra Blanca Regional Airport);United States
SOW;KSOW;Show Low (Show Low Regional Airport);United States
MYL;KMYL;McCall (McCall Municipal Airport);United States
SMN;KSMN;Salmon (Lemhi County Airport);United States
MMH;KMMH;Mammoth Lakes (Mammoth Yosemite Airport);United States
FRD;KFHR;Friday Harbor (Friday Harbor Airport);United States
ESD;KORS;Eastsound (Orcas Island Airport);United States
AST;KAST;Astoria (Astoria Regional Airport);United States
ONP;KNOP;Newport (Newport Municipal Airport);United States
EMK;PAEM;Emmonak (Emmonak Airport);United States
UNK;PAUN;Unalakleet (Unalakleet Airport);United States
UUK;PAKU;Kuparuk (Ugnu-Kuparuk Airport);United States
SHX;PAHX;Shageluk (Shageluk Airport);United States
CHU;PACH;Chuathbaluk (Chuathbaluk Airport);United States
NUI;PAQT;Nuiqsut (Nuiqsut Airport);United States
EEK;PAEE;Eek (Eek Airport);United States
KUK;PFKA;Kasigluk (Kasigluk Airport);United States
KWT;PFKW;Kwethluk (Kwethluk Airport);United States
KWK;PAGG;Kwigillingok (Kwigillingok Airport);United States
MLL;PADM;Marshall (Marshall Don Hunter Sr. Airport);United States
RSH;PARS;Russian Mission (Russian Mission Airport);United States
KGK;PAJZ;Koliganek (Koliganek Airport);United States
KMO;PAMB;Manokotak (Manokotak Airport);United States
CIK;PACI;Chalkyitsik (Chalkyitsik Airport);United States
EAA;PAEG;Eagle (Eagle Airport);United States
HUS;PAHU;Hughes (Hughes Airport);United States
HSL;PAHL;Huslia (Huslia Airport);United States
NUL;PANU;Nulato (Nulato Airport);United States
VEE;PAVE;Venetie (Venetie Airport);United States
WBQ;PAWB;Beaver (Beaver Airport);United States
CEM;PACE;Central (Central Airport);United States
SHG;PAGH;Shungnak (Shungnak Airport);United States
IYK;KIYK;Inyokern (Inyokern Airport);United States
VIS;KVIS;Visalia (Visalia Municipal Airport);United States
MCE;KMCE;Merced (Merced Municipal Airport);United States
CYR;SUCA;Colonia (Laguna de Los Patos International Airport);Uruguay
CPQ;SDAM;Campinas (Amarais Airport);Brazil
TWB;YTWB;Toowoomba;Australia
AYK;UAUR;Arkalyk (Arkalyk Airport);Kazakhstan
AGN;PAGN;Angoon (Angoon Seaplane Base);United States
ELV;PAEL;Elfin Cove (Elfin Cove Seaplane Base);United States
FNR;PANR;Funter Bay (Funter Bay Seaplane Base);United States
HNH;PAOH;Hoonah (Hoonah Airport);United States
AFE;PAFE;Kake (Kake Airport);United States
MTM;PAMM;Metakatla (Metlakatla Seaplane Base);United States
HYG;PAHY;Hydaburg (Hydaburg Seaplane Base);United States
EGX;PAII;Egegik (Egegik Airport);United States
KPV;PAPE;Perryville (Perryville Airport);United States
PIP;PAPN;Pilot Point (Pilot Point Airport);United States
WSN;PFWS;South Naknek (South Naknek Airport);United States
AKK;PAKH;Akhiok (Akhiok Airport);United States
KYK;PAKY;Karluk (Karuluk Airport);United States
KLN;PALB;Larsen Bay (Larsen Bay Airport);United States
ABL;PAFM;Ambler (Ambler Airport);United States
BKC;PABL;Buckland (Buckland Airport);United States
IAN;PAIK;Kiana (Bob Baker Memorial Airport);United States
OBU;PAOB;Kobuk (Kobuk Airport);United States
ORV;PFNO;Noorvik (Robert Curtis Memorial Airport);United States
WLK;PASK;Selawik (Selawik Airport);United States
KTS;PFKT;Brevig Mission (Brevig Mission Airport);United States
ELI;PFEL;Elim (Elim Airport);United States
GLV;PAGL;Golovin (Golovin Airport);United States
TLA;PATE;Teller (Teller Airport);United States
WAA;PAIW;Wales (Wales Airport);United States
WMO;PAWM;White Mountain (White Mountain Airport);United States
KKA;PAKK;Koyuk (Koyuk Alfred Adams Airport);United States
SMK;PAMK;St. Michael (St. Michael Airport);United States
SKK;PFSH;Shaktoolik (Shaktoolik Airport);United States
TNC;PATC;Tin City (Tin City LRRS Airport);United States
AKB;PAAK;Atka (Atka Airport);United States
IKO;PAKO;Nikolski (Nikolski Air Station);United States
CYT;PACY;Yakataga (Yakataga Airport);United States
AUK;PAUK;Alakanuk (Alakanuk Airport);United States
KPN;PAKI;Kipnuk (Kipnuk Airport);United States
KFP;PAKF;False Pass (False Pass Airport);United States
NLG;PAOU;Nelson Lagoon;United States
PML;PAAL;Cold Bay (Port Moller Airport);United States
KLW;PAKW;Klawock (Klawock Airport);United States
KWN;PAQH;Quinhagak (Quinhagak Airport);United States
KOT;PFKO;Kotlik (Kotlik Airport);United States
KYU;PFKU;Koyukuk (Koyukuk Airport);United States
SCM;PACM;Scammon Bay (Scammon Bay Airport);United States
NNL;PANO;Nondalton (Nondalton Airport);United States
KKH;PADY;Kongiganak (Kongiganak Airport);United States
NIB;PAFS;Nikolai (Nikolai Airport);United States
AKI;PFAK;Akiak (Akiak Airport);United States
AIN;PAWI;Wainwright (Wainwright Airport);United States
APZ;SAHZ;ZAPALA;Argentina
PNT;SCNT;Puerto Natales;Chile
SGV;SAVS;Sierra Grande;Argentina
IGB;SAVJ;Ingeniero Jacobacci;Argentina
NCN;PFCB;Chenega (Chenega Bay Airport);United States
6K8;PFTO;Tok (Tok Junction Airport);United States
IRC;PACR;Circle (Circle City Airport);United States
SLQ;PASL;Sleetmute (Sleetmute Airport);United States
HKB;PAHV;Healy (Healy River Airport);United States
AQC;PAQC;Klawock (Klawock Seaplane Base);United States
MHM;PAMH;Lake Minchumina (Minchumina Airport);United States
MLY;PAML;Manley Hot Springs (Manley Hot Springs Airport);United States
YSO;CCD4;Postville (Postville Airport);Canada
YWB;CYKG;Kangiqsujuaq (Kangiqsujuaq - Wakeham Bay Airport);Canada
YTF;CYTF;Alma (Alma Airport);Canada
YGV;CYGV;Havre-Saint-Pierre (Havre Saint-Pierre Airport);Canada
YXK;CYXK;Rimouski (Rimouski Airport);Canada
XTL;CYBQ;Tadoule Lake (Tadoule Lake Airport);Canada
XLB;CZWH;Lac Brochet (Lac Brochet Airport);Canada
XSI;CZSN;South Indian Lake (South Indian Lake Airport);Canada
YBT;CYBT;Brochet (Brochet Airport);Canada
ZGR;CZGR;Little Grand Rapids (Little Grand Rapids Airport);Canada
YCR;CYCR;Cross Lake (Cross Lake - Charlie Sinclair Memorial Airport);Canada
YRS;CYRS;Red Sucker Lake (Red Sucker Lake Airport);Canada
YOP;CYOP;Rainbow Lake (Rainbow Lake Airport);Canada
YBY;CYBF;Bonnyville (Bonnyville Airport);Canada
ZNA;CAC8;Nanaimo (Nanaimo Harbour Water Airport);Canada
YJM;CYJM;Fort St. James (Fort St. James - Perison Airport);Canada
YDT;CZBB;Boundary Bay (Boundary Bay Airport);Canada
ZEL;CYJQ;Bella Bella (Bella Bella Airport);Canada
RNI;MNCI;Corn Island (Corn Island Airport);Nicaragua
BZA;MNBZ;Bonanza (Bonanza Airport);Nicaragua
RFS;MNRT;Rosita (Rosita Airport);Nicaragua
SIU;MNSI;Siuna (Siuna Airport);Nicaragua
WSP;MNWP;Waspam (Waspam Airport);Nicaragua
RIK;MRCR;Carrillo (Carrillo Airport);Costa Rica
FUS;FUSS;Fussen;Germany
COZ;MDCZ;Constanza (Constanza Airport);Dominican Republic
NEG;MKNG;Negril (Negril Aerodrome);Jamaica
RVR;TJRV;Ceiba (Jose Aponte de la Torre Airport);Puerto Rico
ARR;SAVR;Alto Rio Senguer (Alto Rio Senguer Airport);Argentina
JSM;SAWS;Jose de San Martin (Jose de San Martin Airport);Argentina
UYU;SLUY;Uyuni (Uyuni Airport);Bolivia
RBQ;SLRQ;Rerrenabaque (Rurrenabaque Airport);Bolivia
ABF;NGAB;Abaiang Atoll (Abaiang Atoll Airport);Kiribati
ABN;SMBN;Albina (Albina Airstrip);Suriname
DRJ;SMDA;Drietabbetje (Drietabbetje Airstrip);Suriname
ICK;SMNI;Nieuw Nickerie (Majoor Henry Fernandes Airport);Suriname
OEM;SMPA;Paloemeu (Vincent Fayks Airport);Suriname
SMZ;SMST;Stoelmans Eiland (Stoelmans Eiland Airstrip);Suriname
TOT;SMCO;Totness (Totness Airstrip);Suriname
AGI;SMWA;Wageningen (Wageningen Airstrip);Suriname
KIA;PSKA;Kaieteur Falls (Kaieteur International Airport);Guyana
CSC;MRCA;Guapiles (Codela Airport);Costa Rica
ORJ;SYOR;Orinduik (Orinduik Airport);Guyana
NAI;SYAN;Annai (Annai Airport);Guyana
IMB;SYIB;Imbaimadai (Imbaimadai Airport);Guyana
KAR;SYKM;Kamarang (Kamarang Airport);Guyana
USI;SYMB;Mabaruma (Mabaruma Airport);Guyana
MHA;SYMD;Mahdia (Mahdia Airport);Guyana
PJC;SGPJ;Pedro Juan Caballero (Dr. Augusto Roberto Fuster International Airport);Paraguay
ACD;SKAD;Acandi (Alcides Fernandez Airport);Colombia
RVE;SKSA;Saravena (Los Colonizadores Airport);Colombia
VGZ;SKVG;Villa Garzon (Villa Garzon Airport);Colombia
EBG;SKEB;El Bagre (El Bagre Airport);Colombia
CAQ;SKCU;Caucasia (Juan H. White);Colombia
COG;SKCD;Condoto (Mandinga Airport);Colombia
TLU;SKTL;Tolu (Golfo de Morrosquillo Airport);Colombia
CFB;SBCB;Cabo Frio (Cabo Frio International Airport);Brazil
OPS;SWSI;Sinop (Sinop Airport);Brazil
GRP;SWGI;Gurupi (Gurupi Airport);Brazil
CMP;SNKE;Santana do Araguaia (Campo Alegre Airport);Brazil
BVS;SNVS;Breves (Breves Airport);Brazil
SFK;SNSW;Soure (Soure Airport);Brazil
PIN;SWPI;Parintins (Julio Belem Airport);Brazil
BRA;SNBR;Barreiras (Barreiras Airport);Brazil
STZ;SWST;Santa Terezinha (Confresa Airport);Brazil
MQH;SBMC;Minacu (Minacu Airport);Brazil
AUX;SWGN;Araguaina (Araguaina Airport);Brazil
NVP;SWNA;Novo Aripuana (Novo Aripuana Airport);Brazil
LVR;SWFE;Lucas do Rio Verde (Bom Futuro Airport);Brazil
FRC;SIMK;Franca (Franca Airport);Brazil
DOU;SSDO;Dourados (Dourados Airport);Brazil
LBR;SWLB;Labrea (Labrea Airport);Brazil
ROO;SWRD;Rondonopolis (Rondonopolis Airport);Brazil
GPB;SBGU;Guarapuava (Tancredo Thomaz de Faria Airport);Brazil
JCB;SSJA;Joacaba (Joacaba Airport);Brazil
RVD;SWLC;Rio Verde (General leite de Castro Airport);Brazil
AAX;SBAX;Araxa (Araxa Airport);Brazil
MBZ;SWMW;Maues (Maues Airport);Brazil
RBB;SWBR;Borba (Borba Airport);Brazil
CIZ;SWKO;Coari (Coari Airport);Brazil
BAZ;SWBC;Barcelos (Barcelos Airport);Brazil
DMT;SWDM;Diamantino (Diamantino Airport);Brazil
GNM;SNGI;Guanambi (Guanambi Airport);Brazil
QDJ;DAFI;Djelfa (Tsletsi Airport);Algeria
LBZ;FNLK;Lucapa (Lucapa Airport);Angola
KNP;FNCP;Kapanda (Kapanda Airport);Angola
AMC;FTTN;Am Timan (Am Timan Airport);Chad
GSQ;HEOW;Sharq Al-Owainat (Sharq Al-Owainat Airport);Egypt
MRB;KMRB;Martinsburg (Eastern WV Regional Airport);United States
AWA;HALA;Awasa (Awasa Airport);Ethiopia
JIJ;HAJJ;Jijiga (Jijiga Airport);Ethiopia
MKS;HAMA;Mekane Selam (Mekane Salam Airport);Ethiopia
DBM;HADM;Debre Marqos;Ethiopia
DBT;HADT;Debre Tabor (Debre Tabor Airport);Ethiopia
QHR;HAHM;Debre Zeyit (Harar Meda Airport);Ethiopia
GOB;HAGB;Goba (Robe Airport);Ethiopia
MYB;FOOY;Mayumba (Mayumba Airport);Gabon
RBX;HSMK;Rumbek (Rumbek Airport);Sudan
CPA;GLCP;Greenville (Cape Palmas Airport);Liberia
MAX;GOSM;Matam (Ouro Sogui Airport);Senegal
BDI;FSSB;Bird Island (Bird Island Airport);Seychelles
WHF;HSSW;Wadi Halfa (Wadi Halfa Airport);Sudan
NBE;DTNZ;Enfidha (Enfidha - Zine El Abidine Ben Ali International Airport);Tunisia
PAF;HUPA;Pakuba (Pakuba Airport);Uganda
HTY;LTDA;Hatay (Hatay Airport);Turkey
RVV;NTAV;Raivavae (Raivavae Airport);French Polynesia
ILD;LEDA;Lleida (Lleida-Alguaire Airport);Spain
BIU;BIBD;Bildudalur (Bildudalur Airport);Iceland
GJR;BIGJ;Gjogur (Gjogur Airport);Iceland
SAK;BIKR;Saudarkrokur;Iceland
IIA;EIMN;Inishmaan (Inishmaan Aerodrome);Ireland
TDK;UAAT;Taldykorgan (Taldykorgan Airport);Kazakhstan
ULG;ZMUL;Olgii (Olgii Airport);Mongolia
VGD;ULWW;Vologda (Vologda Airport);Russia
LDG;ULAL;Arkhangelsk (Leshukonskoye Airport);Russia
HSK;LEHC;Huesca (Huesca-Pirineos Airport);Spain
CQM;LERL;Ciudad Real (Ciudad Real Central Airport);Spain
NJF;ORNI;Najaf (Al Najaf International Airport);Iraq
CSA;EGEY;Colonsay (Colonsay Airport);United Kingdom
RKH;KUZA;Rock Hill (Rock Hill York Co Bryant Airport);United States
AGC;KAGC;Pittsburgh (Allegheny County Airport);United States
NZC;KVQQ;Jacksonville (Cecil Field);United States
FTY;KFTY;Atlanta (Fulton County Airport Brown Field);United States
TSO;EGHT;Tresco (Tresco Heliport);United Kingdom
TII;OATN;Tarin Kowt (Tarin Kowt Airport);Afghanistan
ZAJ;OAZJ;Zaranj (Zaranj Airport);Afghanistan
CCN;OACC;Chaghcharan (Chaghcharan Airport);Afghanistan
FUG;ZSFY;Fuyang (Fuyang Airport);China
BSD;ZPBS;Baoshan (Baoshan Airport);China
OSU;KOSU;Columbus (Ohio State University Airport);United States
ADS;KADS;Addison;United States
DTS;KDTS;Destin;United States
KHE;UKOH;Kherson (Chernobayevka Airport);Ukraine
SZS;NZRC;Stewart Island (Ryans Creek Aerodrome);New Zealand
YQI;CYQI;Yarmouth (Yarmouth Airport);Canada
ISO;KISO;Kinston (Kinston Regional Jetport);United States
FFA;KFFA;Kill Devil Hills (First Flight Airport);United States
LNJ;ZPLC;Lincang (Lincang Airport);China
CKS;SBCJ;Parauapebas (Carajas Airport);Brazil
MWK;WIOM;Anambas Islands (Matak Airport);Indonesia
PGU;OIBP;Khalije Fars (Persian Gulf Airport);Iran
YES;OISY;Yasuj (Yasuj Airport);Iran
OSB;ORBM;Mosul (Mosul International Airport);Iraq
TJH;RJBT;Toyooka (Tajima Airport);Japan
AXJ;RJDA;Amakusa (Amakusa Airfield);Japan
KKX;RJKI;Kikai (Kikai Airport);Japan
AGJ;RORA;Aguni (Aguni Airport);Japan
UGA;ZMBN;Bulgan (Bulgan Airport);Mongolia
ULO;ZMUG;Ulaangom (Ulaangom Airport);Mongolia
BPR;RPVW;Borongan (Borongan Airport);Philippines
LBX;RPLU;Lubang (Lubang Community Airport);Philippines
TJU;UTDK;Kulyab (Kulob Airport);Tajikistan
CMJ;RCCM;Cimei (Cimei Airport);Taiwan
TAZ;UTAT;Dasoguz (Dasoguz Airport);Turkmenistan
BWB;YBWX;Barrow Island (Barrow Island Airport);Australia
DRB;YDBY;Derby (Derby Airport);Australia
WGE;YWLG;Walgett (Walgett Airport);Australia
BRT;YBTI;Bathurst Island (Bathurst Island Airport);Australia
DKI;YDKI;Dunk Island (Dunk Island Airport);Australia
LZR;YLZI;Lizard Island (Lizard Island Airport);Australia
HLT;YHML;Hamilton (Hamilton Airport);Australia
HCQ;YHLC;Halls Creek (Halls Creek Airport);Australia
FIZ;YFTZ;Fitzroy Crossing (Fitzroy Crossing Airport);Australia
RVT;YNRV;Ravensthorpe (Ravensthorpe Airport);Australia
PVU;KPVU;Provo (Provo Municipal Airport);United States
SBS;KSBS;Steamboat Springs (Steamboat Springs Airport-Bob Adams Field);United States
DTA;KDTA;Delta (Delta Municipal Airport);United States
RIF;KRIF;Richfield (Richfield Minicipal Airport);United States
PUC;KPUC;Price (Carbon County Regional-Buck Davis Field);United States
LAM;KLAM;Los Alamos (Los Alamos Airport);United States
BXS;KBXS;Borrego Springs (Borrego Valley Airport);United States
HII;KHII;Lake Havasu City (Lake Havasu City Airport);United States
INW;KINW;Winslow (Winslow-Lindbergh Regional Airport);United States
DGL;KDGL;Douglas (Douglas Municipal Airport);United States
MZK;NGMK;Marakei (Marakei Airport);Kiribati
AEA;NGTB;Abemama (Abemama Atoll Airport);Kiribati
AAK;NGUK;Buariki (Aranuka Airport);Kiribati
KUC;NGKT;Kuria (Kuria Airport);Kiribati
AIS;NGTR;Arorae (Arorae Island Airport);Kiribati
TMN;NGTM;Tamana (Tamana Airport);Kiribati
BEZ;NGBR;Beru Island (Beru Island Airport);Kiribati
NIG;NGNU;Nikunau (Nikunau Airport);Kiribati
BBG;NGTU;Butaritari (Butaritari Atoll Airport);Kiribati
MTK;NGMN;Makin (Makin Airport);Kiribati
MNK;NGMA;Maiana (Maiana Airport);Kiribati
NON;NGTO;Nonouti (Nonouti Airport);Kiribati
TSU;NGTS;Tabiteuea (Tabiteuea South Airport);Kiribati
WTZ;NZWT;Whitianga (Whitianga Airport);New Zealand
KTF;NZTK;Takaka (Takaka Aerodrome);New Zealand
AFT;AGAF;Afutara (Afutara Airport);Solomon Islands
RNA;AGAR;Ulawa (Ulawa Airport);Solomon Islands
CHY;AGGC;Choiseul Bay (Choiseul Bay Airport);Solomon Islands
NNB;AGGT;Santa Ana (Santa Ana Airport);Solomon Islands
XYA;AGGY;Yandina (Yandina Airport);Solomon Islands
BPF;AGBT;Batuna (Batuna Airport);Solomon Islands
BOW;KBOW;Bartow (Bartow Municipal Airport);United States
KMW;UUBD;Kostroma (Sokerkino);Russia
FTI;NSFQ;Fiti'uta (Fitiuta Airport);American Samoa
LVK;KLVK;Livermore (Livermore Municipal);United States
MPI;KMPI;Mariposa (MariposaYosemite);United States
GFY;FYGF;Grootfontein;Namibia
NDU;FYRU;Rundu;Namibia
AGM;BGAM;Angmagssalik (Tasiilaq);Greenland
TRM;KTRM;Palm Springs (Jacqueline Cochran Regional Airport);United States
SMO;KSMO;Santa Monica (Santa Monica Municipal Airport);United States
UDD;KUDD;Palm Springs (Bermuda Dunes Airport);United States
ZSY;KSDL;Scottsdale (Scottsdale Airport);United States
OLM;KOLM;Olympia (Olympia Regional Airpor);United States
DWA;KDWA;Davis-Woodland-Winters (Yolo County Airport);United States
RIL;KRIL;Rifle (Garfield County Regional Airport);United States
SAA;KSAA;SARATOGA (Shively Field Airport);United States
PDK;KPDK;Atlanta (Dekalb-Peachtree Airport);United States
BMG;KBMG;Bloomington (Monroe County Airport);United States
SUA;KSUA;Stuart (Witham Field Airport);United States
MMU;KMMU;Morristown (Morristown Municipal Airport);United States
APC;KAPC;Napa (Napa County Airport);United States
SDM;KSDM;San Diego (Brown Field Municipal Airport);United States
ECP;KECP;Panama City (Panama City-NW Florida Bea.);United States
SBD;KSBD;San Bernardino (San Bernardino International Airport);United States
VAL;SNVB;Valenca (Valenca Airport);Brazil
MVF;SBMW;Mossoro (Dix Sept Rosado Airport);Brazil
CAU;SNRU;Caruaru (Caruaru Airport);Brazil
AWK;PWAK;Wake island (Wake Island Afld);Wake Island
QNV;SDNY;Nova Iguacu (Aeroclube de Nova Iguacu);Brazil
SQL;KSQL;San Carlos (San Carlos Airport);United States
OSZ;EPKO;Koszalin (Koszalin - Zegrze Pomorskie Airport);Poland
RWI;KRWI;Rocky Mount (Rocky Mount Wilson Regional Airport);United States
SEE;KSEE;El Cajon (Gillespie);United States
PHA;VVPR;Phan Rang (Phan Rang Airport);Vietnam
SQH;VVNS;Son-La (Na-San Airport);Vietnam
TKF;KTRK;Truckee (Truckee-Tahoe Airport);United States
FRJ;LFTU;Frejus (Frejus Saint Raphael);France
GEX;YGLG;Geelong (Geelong Airport);Australia
RYY;KRYY;Atlanta (Cobb County Airport-Mc Collum Field);United States
4U9;K4U9;Dell (Dell Flight Strip);United States
LVM;KLVM;Livingston-Montana (Mission Field Airport);United States
6S0;K6S0;Big Timber (Big Timber Airport);United States
BIV;KBIV;Holland (Tulip City Airport);United States
HEN;EFHE;Helsinki (Hernesaari Heliport);Finland
JRA;KJRA;New York (West 30th St. Heliport);United States
LAL;KLAL;Lakeland (Lakeland Linder Regional Airport);United States
SYH;VNSB;Syangboche;Nepal
IDL;KIDL;New York (Idlewild Intl);United States
RBK;KF70;Murrieta-Temecula (French Valley Airport);United States
FNU;LIER;Oristano (Fenosu);Italy
MYQ;VOMY;Mysore (Mysore Airport);India
PCW;KPCW;Port Clinton (Erie-Ottawa Regional Airport);United States
MGY;KMGY;Dayton (Dayton-Wright Brothers Airport);United States
RID;KRID;Richmond (Richmond Municipal Airport);United States
FDY;KFDY;Findley (Findlay Airport);United States
PEA;YPSH;Penneshaw (Penneshaw Airport);Australia
KFX;KAUF;Kaufbeuren (Kaufbeuren BF);Germany
MUQ;MUNI;Munich (Munich HBF);Germany
NUR;NURN;Nurnberg (Nurnberg HBF);Germany
EBE;EBEN;Ebenhofen (Ebenhofen BF);Germany
AUB;AUGS;Augsburg (Augsburg HBF);Germany
BIE;BIES;Biessenhofen (Biessenhofen BF);Germany
BUH;BUCH;Buchloe (Buchloe BF);Germany
FUX;FUSN;Fussen (Fussen BF);Germany
KEX;KEMP;Kempten (Kempten HBF);Germany
OAL;MARK;Marktoberdorf (Marktoberdorf BF);Germany
MOS;MARO;Marktoberdorf (Marktoberdorf Schule);Germany
ESX;ESSE;Essen (Essen HBF);Germany
BOX;BOCH;Bochum (Bochum HBF);Germany
KOX;KOLN;Koln (Koln HBF);Germany
BBP;EGHJ;Bembridge;United Kingdom
LES;LEUT;Leuterschach (Leuterschach BF);Germany
SPF;KSPF;Spearfish-South Dakota (Black Hills Airport-Clyde Ice Field);United States
KNO;EBKW;Knokke (Knokke-Heist Westkapelle Heliport);Belgium
QYD;EPOK;Gdynia;Poland
OLV;KOLV;Olive Branch (Olive Branch Muni);United States
ONQ;LTAS;Zonguldak;Turkey
BJC;KBJC;Broomfield-CO (Rocky Mountain Metropolitan Airport);United States
SLE;KSLE;Salem (McNary Field);United States
UTM;KUTA;Tunica (Tunica Municipal Airport);United States
ZKB;FLKY;Kasaba Bay (Kasaba Bay Airport);Zambia
LND;LIND;Lindau (Lindau HBF);Germany
MWC;KMWC;Milwaukee (Lawrence J Timmerman Airport);United States
JVL;KJVL;Janesville (Southern Wisconsin Regional Airport);United States
GKY;KGKY;Arlington (Arlington Municipal);United States
LZU;KLZU;Lawrenceville (Gwinnett County Airport-Briscoe Field);United States
BWG;KBWG;Bowling Green (Bowling Green-Warren County Regional Airport);United States
RVS;KRVS;Tulsa (Richard Lloyd Jones Jr Airport);United States
NHD;OMDM;Minhad AB (Minhad HB);United Arab Emirates
KGO;UKKG;Kirovograd;Ukraine
DBB;HEAL;Dabaa City (Alalamain Intl.);Egypt
BCE;KBCE;Bryce Canyon;United States
HDB;EDIU;Heidelberg;Germany
BUY;KBUY;Burlington (Burlington-Alamance Regional Airport);United States
CKL;UUMU;Shchyolkovo (Chkalovsky Airport);Russia
TCZ;ZUTC;Tengchong (Tengchong Tuofeng Airport);China
UKS;UKFB;Sevastopol (Belbek Sevastopol International Airport);Ukraine
JCI;KIXD;Olathe (New Century AirCenter Airport);United States
ESN;KESN;Easton (Easton-Newnam Field Airport);United States
HMR;ENHA;Hamar (Stafsberg Airport);Norway
MYV;KMYV;Yuba City (Yuba County Airport);United States
DUC;KDUC;Duncan (Halliburton Field Airport);United States
UVA;KUVA;Uvalde (Garner Field);United States
LOT;KLOT;Lockport (Lewis University Airport);United States
CCR;KCCR;Concord (Buchanan Field Airport);United States
OCA;07FA;Ocean Reef Club Airport (Key Largo);United States
YUS;ZLYS;Yushu (Yushu Batang);China
HIA;ZSSH;Huai An (Huai An Lianshui Airport);China
YOO;CYOO;Oshawa (Oshawa Airport);Canada
LHA;EDTL;Lahr (Lahr Airport);Germany
SGH;KSGH;Springfield (Springfield-Beckly Municipal Airport);United States
MSI;SIAM;South Aari Atoll (Sun Island Airport);Maldives
HEX;MDHE;Santo Domingo (Herrera International Airport);Dominican Republic
CDA;YCOO;Cooinda;Australia
JAB;YJAB;Jabiru;Australia
RGB;REGE;Regensburg (Regensburg HBF);Germany
TLG;TREU;Treuchtlingen (Treuchtlingen BF);Germany
HGS;GFHA;Freetown (Hastings Airport);Sierra Leone
TOP;KTOP;Topeka (Philip Billard Muni);United States
NGQ;ZUAL;Shiquanhe (Gunsa);China
CSO;EDBC;Cochstedt (Magdeburg-Cochstedt);Germany
WZB;WURZ;Wurzburg (Wurzburg HBF);Germany
TKI;KTKI;DALLAS (Collin County Regional Airport at Mc Kinney);United States
PWK;KPWK;Chicago-Wheeling (Chicago Executive);United States
KLS;KKLS;Kelso (Kelso Longview);United States
ZTA;NTGY;Tureia (Tureia Airport);French Polynesia
PUE;MPOA;Puerto Obaldia;Panama
KHC;UKFK;Kerch (Kerch Intl);Ukraine
UKA;HKUK;Ukunda (Ukunda Airport);Kenya
ILN;KILN;Wilmington (Wilmington Airborne Airpark);United States
AVW;KAVQ;Tucson (Marana Regional);United States
CGZ;KCGZ;Casa Grande (Casa Grande Municipal Airport);United States
BXK;KBXK;Buckeye (Buckeye Municipal Airport);United States
E63;KE63;Gila Bend (Gila Bend Municipal Airport);United States
MMI;KMMI;Athens (McMinn Co);United States
STK;KSTK;Sterling (Sterling Municipal Airport);United States
RWL;KFWL;Rawlins (Rawlins Municipal Airport-Harvey Field);United States
YZY;CYZY;Mackenzie British Columbia (Mackenzie Airport);Canada
CDW;KCDW;Caldwell (Caldwell Essex County Airport);United States
AIZ;KAIZ;Kaiser Lake Ozark (Lee C Fine Memorial Airport);United States
KIP;KIEV;KIEV (KIEV  INTERNATIONAL AIRPORT);Ukraine
BAM;BAMB;Bamberg (Bamberg BF);Germany
IGS;INGS;Ingolstadt (Ingolstadt BF);Germany
TVI;KTVI;Thomasville (Thomasville Regional Airport);United States
HSH;KHND;Henderson (Henderson Executive Airport);United States
GML;UKKM;Kiev (Gostomel Antonov);Ukraine
TMA;KTMA;Tifton (Henry Tift Myers Airport);United States
QXR;EPRA;RADOM;Poland
DVT;KDVT;Phoenix  (Deer Valley Municipal Airport);United States
YBW;CYBW;Calgary (Calgary Springbank Airport);Canada
YGE;CYGE;Golden (Golden Airport);Canada
YRV;CYRV;Revelstoke (Revelstoke Airport);Canada
HDO;KHDO;Hondo (Hondo Municipal Airport);United States
ZHY;ZLZW;Zhongwei (Zhongwei Xiangshan Airport);China
MCL;PAIN;McKinley Park (McKinley National Park Airport);United States
LHD;PALH;Anchorage (Lake Hood Seaplane Base);United States
PPC;PAPR;Prospect Creek (Prospect Creek Airport);United States
KHW;FBKR;Khwai River (Khwai River Lodge);Botswana
TXG;RCLG;Taichung (Taichung Airport);Taiwan
FZG;KFZG;Fitzgerald (Fitzgerald Municipal Airport);United States
XYE;VYYE;Ye;Burma
XEG;KGST;Kingston (Kingston Train Station);Canada
DWC;OMDW;Dubai (Dubai Al Maktoum);United Arab Emirates
RKP;KRKP;Rockport (Aransas County Airport);United States
MVV;LFHM;Verdun (Megeve Airport);France
MFX;LFKX;Ajaccio (Meribel Airport);France
OKF;FYOO;Okaukuejo (Okaukuejo Airport);Namibia
OKU;FYMO;Mokuti Lodge (Mokuti Lodge Airport);Namibia
PSH;EDXO;Sankt Peter-Ording (St. Peter-Ording Airport);Germany
CKF;KCKF;Cordele (Crisp County Cordele Airport);United States
OMN;KOMN;Ormond Beach (Ormond Beach municipal Airport);United States
TTD;KTTD;Troutdale (Portland Troutdale);United States
HIO;KHIO;Hillsboro (Portland Hillsboro);United States
KHT;OAKS;Khost (FOB Salerno);Afghanistan
NYT;VYNT;NAYPYITAW;Burma
GAI;KGAI;Gaithersburg (Montgomery County Airpark);United States
AZ3;OASA;Sharona;Afghanistan
YTA;CYTA;Pembroke (Pembroke Airport);Canada
TSB;FYTM;Tsumeb (Tsumeb Airport);Namibia
YSD;CYSD;Suffield (Suffield Heliport);Canada
BNU;SSBL;BLUMENAU (Aeroporto Blumenau);Brazil
CIO;LRT2;Timisoara (Aeroclub Cioca);Romania
CVX;KCVX;Charelvoix (Charlevoix Municipal Airport);United States
YCC;CYCC;Cornwall (Cornwall Regional Airport);Canada
IZA;SDZY;Juiz de Fora (Zona da Mata Regional Airport);Brazil
XFL;KXFL;Flagler (Flagler County Airport);United States
MVL;KMVL;Morrisville (Morrisville Stowe State Airport);United States
RBD;KRBD;Dallas (Dallas Executive Airport);United States
WST;KWST;Washington County (Westerly State Airport);United States
BID;KBID;Block Island (Block Island State Airport);United States
NME;PAGT;Nightmute (Nightmute Airport);United States
OOK;PAOO;Toksook Bay (Toksook Bay Airport);United States
OBY;BGSC;Ittoqqortoormiit (Ittoqqortoormiit Heliport);Greenland
VIN;UKWW;Vinnitsa;Ukraine
BGE;KBGE;Bainbridge (Decatur County Industrial Air Park);United States
ZKG;CTK6;Kegaska (Kegaska Airport);Canada
YBI;CCE4;Black Tickle (Black Tickle Airport);Canada
SPZ;KSPZ;Silver Springs (Silver Springs Airport);United States
WHP;KWHP;Los Angeles (Whiteman Airport);United States
MAE;KMAE;Madera (Madera Municipal Airport);United States
YZZ;CAD4;Trail (Trail Airport);Canada
YAB;CJX7;Arctic Bay (Arctic Bay Airport);Canada
MPY;SOOA;Maripasoula (Maripasoula Airport);French Guiana
LDX;SOOM;Saint-Laurent-du-Maroni (Saint-Laurent-du-Maroni Airport);French Guiana
KJI;ZWKN;Burqin (Kanas Airport);China
CPB;SKCA;Capurgana (Capurgana Airport);Colombia
HMB;HEMK;Sohag (Sohag International);Egypt
RVY;SURV;Rivera (Rivera International Airport);Uruguay
POJ;SNPD;Patos de Minas (Patos de Minas Airport);Brazil
JTC;SJTC;Bauru (Bauru-Arealva);Brazil
OIA;SDOW;Ourilandia do Norte (Ourilandia do Norte Airport);Brazil
RDC;SNDC;Redencao (Redencao Airport);Brazil
SXX;SNFX;Sao Felix do Xingu (Sao Felix do Xingu Airport);Brazil
BYO;SJDB;Bointo (Bonito Airport);Brazil
SXO;SWFX;Sao Felix do Araguaia (Sao Felix do Araguaia Airport);Brazil
CFC;SBCD;Cacador (Carlos Alberto da Costa Neves Airport);Brazil
CAF;SWCA;Carauari (Carauari Airport);Brazil
ERN;SWEI;Eirunepe (Amaury Feitosa Tomaz Airport);Brazil
CCI;SSCK;Concordia (Concordia Airport);Brazil
FBE;SSFB;Francisco Beltrao (Paulo Abdala Airport);Brazil
CFO;SJHG;Confresa (Confresa Airport);Brazil
AAF;KAAF;Apalachicola (Apalachicola Regional Airport);United States
UMU;SSUM;Umuarama (Orlando de Carvalho Airport);Brazil
DTI;SNDT;Diamantina (Diamantina Airport);Brazil
FBA;SWOB;Fonte Boa (Fonte Boa Airport);Brazil
OLC;SDCG;Sao Paulo de Olivenca (Senadora Eunice Micheles Airport);Brazil
HUW;SWHT;Humaita (Humaita Airport);Brazil
IRZ;SWTP;Santa Isabel do Rio Negro (Tapuruquara Airport);Brazil
ORX;SNOX;Oriximina (Oriximina Airport);Brazil
UNA;SBTC;Una (Una-Comandatuba Airport);Brazil
TEF;YTEF;Telfer (Telfer Airport);Australia
GZP;LTFG;Alanya (Gazipasa Airport);Turkey
DQH;KDQH;Douglas (Douglas Municipal Airport);United States
FRP;KFPR;Fort Pierce (St Lucie County International Airport);United States
ALX;ALX_;Alexandria;United States
TAN;KTAN;Taunton (Taunton Municipal Airport - King Field);United States
PYM;KPYM;Plymouth (Plymouth Municipal Airport);United States
OQU;KOQU;North Kingstown (Quonset State Airport);United States
OWD;KOWD;Norwood (Norwood Memorial Airport);United States
BAF;KBAF;Westfield (Barnes Municipal);United States
IJD;KIJD;Willimantic (Windham Airport);United States
MGJ;KMGJ;Montgomery (Orange County Airport);United States
CXY;KCXY;Harrisburg (Capital City Airport);United States
GHG;KGHG;Marshfield (Marshfield Municipal Airport);United States
DXR;KDXR;Danbury (Danbury Municipal Airport);United States
ASH;KASH;Nashua (Boire Field Airport);United States
LWM;KLWM;Lawrence (Lawrence Municipal Airport);United States
OXC;KOXC;Oxford (Waterbury-Oxford Airport);United States
FIT;KFIT;Fitchburg (Fitchburg Municipal Airport);United States
VPC;KVPC;Cartersville (Cartersville Airport);United States
PYP;KPYP;Centre (Centre-Piedmont-Cherokee County Regional Airport);United States
RMG;KRMG;Rome (Richard B Russell Airport);United States
GAD;KGAD;Gadsden (Northeast Alabama Regional Airport);United States
DKX;KDKX;Knoxville (Knoxville Downtown Island Airport);United States
WDR;KWDR;Winder (Barrow County Airport);United States
JYL;KJYL;Sylvania (Plantation Airpark);United States
DNN;KDNN;Dalton (Dalton Municipal Airport);United States
CTJ;KCTJ;Carrollton (West Georgia Regional Airport - O V Gray Field);United States
LGC;KLGC;LaGrange (LaGrange-Callaway Airport);United States
MLJ;KMLJ;Milledgeville (Baldwin County Airport);United States
PIM;KPIM;Pine Mountain (Harris County Airport);United States
FFC;KFFC;Atlanta (Atlanta Regional Airport - Falcon Field);United States
GVL;KGVL;Gainesville (Lee Gilmer Memorial Airport);United States
PHD;KPHD;New Philadelpha (Harry Clever Field Airport);United States
UDG;KUDG;Darlington (Darlington County Jetport);United States
HXD;KHXD;Hilton Head Island (Hilton Head Airport);United States
DNL;KDNL;Augusta (Daniel Field Airport);United States
MRN;KMRN;Morganton (Foothills Regional Airport);United States
PBX;KPBX;Pikeville (Pike County Airport - Hatcher Field);United States
TOC;KTOC;Toccoa (Toccoa RG Letourneau Field Airport);United States
PLV;UKHP;Poltava;Ukraine
WUU;HSWW;Wau (Wau Airport);Sudan
HUE;HAHU;Humera (Humera Airport);Ethiopia
OYL;HKMY;Moyale (Moyale Airport);Kenya
WYE;GFYE;Yengema (Yengema Airport);Sierra Leone
GBK;GFGK;Gbangbatok (Gbangbatok Airport);Sierra Leone
AFW;KAFW;Fort Worth (Fort Worth Alliance Airport);United States
57C;K57C;East Troy (East Troy Municipal Airport);United States
MYF;NULL;San Diego (Montgomery Field);United States
RMK;YREN;Renmark;Australia
KEW;CPV8;Keewaywin;Canada
YSP;CYSP;Marathon;Canada
YHF;CYHF;Hearst (Rene Fontaine);Canada
YHN;CYHN;Hornepayne;Canada
YKX;CYKX;Kirkland Lake;Canada
YMG;CYMG;Manitouwadge;Canada
YXZ;CYXZ;Wawa;Canada
YEM;CYEM;Manitowaning (Manitoulin East);Canada
YFD;CYFD;Brantford;Canada
LWC;KLWC;Lawrence (Lawrence Municipal);United States
EGT;KEGT;Wellington (Wellington Municipal);United States
PMP;KPMP;Pompano Beach (Pompano Beach Airpark);United States
XMC;YMCO;Mallacoota (Mallacoota Airport);Australia
ULH;OEAO;Al-Ula (Prince Abdul Majeed Airport);Saudi Arabia
EET;KEET;Alabaster (Shelby County Airport);United States
YUE;YYND;Yuendumu ;Australia
LOP;WADL;Praya (Lombok International Airport);Indonesia
OMM;OONR;Marmul;Oman
ZML;CZML;108 Mile Ranch (South Cariboo Regional Airport);Canada
HDG;ZBHD;Handan (Hebei Handan Airport);China
UMP;KUMP;Indianapolis (Indianapolis Metropolitan Airport);United States
LOZ;KLOZ;London (London-Corbin Airport-MaGee Field);United States
WMI;EPMO;Warsaw (Warsaw Modlin);Poland
JXA;ZYJX;Jixi (Jixi Airport);China
YGM;CYGM;Gimli (Gimli Industrial Park Airport);Canada
EYK;USHQ;Beloyarsky;Russia
RAC;KRAC;Racine (John H. Batten Airport);United States
RZP;RPSD;Taytay (Taytay Sandoval);Philippines
RKZ;ZURK;Shigatse (Shigatse Peace Airport);China
REI;KREI;Redlands (Redlands Municipal Airport);United States
RIR;KRIR;Riverside (Flabob Airport);United States
TIW;KTIW;Tacoma (Tacoma Narrows Airport);United States
JKA;KJKA;Gulf Shores (Jack Edwards Airport);United States
HMJ;UKLH;Khmeinitskiy;Ukraine
HIW;RJBH;Hiroshima (Hiroshima-Nishi);Japan
HZL;KHZL;Hazleton (Hazleton Municipal);United States
CBE;KCBE;Cumberland (Greater Cumberland Rgnl.);United States
YBO;CBW4;Bob Quinn Lake;Canada
KLF;UUBS;Kaluga (Grabtsevo);Russia
ELL;FAEA;Lephalale (Ellisras);South Africa
LNR;KLNR;Lone Rock (Tri-County Regional Airport);United States
JOT;KJOT;Joliet (Regional Airport);United States
VYS;KVYS;Peru (Illinois Valley Regional);United States
JXN;KJXN;Jackson (Reynolds Field);United States
BBX;KLOM;Philadelphia (Wings Field);United States
OBE;KOBE;Okeechobee (County);United States
SEF;KSEF;Sebring (Regional - Hendricks AAF);United States
AVO;KAVO;Avon Park (Executive);United States
GIF;KGIF;Winter Haven (Gilbert Airport);United States
ZPH;KZPH;Zephyrhills (Municipal Airport);United States
OCF;KOCF;Ocala (International Airport);United States
JES;KJES;Jesup (Jesup-Wayne County Airport);United States
52A;K52A;Madison (Madison GA Municipal Airport);United States
CCO;KCCO;Newnan (Coweta County Airport);United States
HQU;KHQU;Thomson (McDuffie County Airport);United States
AIK;KAIK;Aiken (Municipal Airport);United States
CDN;KCDN;Camden (Woodward Field);United States
LBT;KLBT;Lumberton (Municipal Airport);United States
SOP;KSOP;Pinehurst-Southern Pines (Moore County Airport);United States
RCZ;KRCZ;Rockingham (Richmond County Airport);United States
DLL;KDLL;Baraboo (Baraboo Wisconsin Dells Airport);United States
SVH;KSVH;Statesville (Regional Airport);United States
BUU;KBUU;Burlington (Municipal Airport);United States
LHV;KLHV;Lock Haven (William T. Piper Mem.);United States
LPR;KLPR;Lorain-Elyria (Lorain County Regional Airport);United States
BKL;KBKL;Cleveland (Burke Lakefront Airport);United States
DKK;KDKK;Dunkirk (Chautauqua County-Dunkirk Airport);United States
VAY;KVAY;Mount Holly (South Jersey Regional Airport);United States
LDJ;KLDJ;Linden (Linden Airport);United States
ANQ;KANQ;Angola (Tri-State Steuben County Airport);United States
VNW;KVNW;Van Wert (Van Wert County Airport);United States
RMY;KRMY;Marshall (Brooks Field Airport);United States
GVQ;KGVQ;Batavia (Genesee County Airport);United States
CLW;KCLW;Clearwater (Clearwater Air Park);United States
CGX;KCGX;Chicago (Meigs Field);United States
JZP;KJZP;Jasper (Pickens County Airport);United States
IGQ;KIGQ;Lansing (Lansing Municipal);United States
RNM;KRNM;Ramona (Ramona Airport);United States
BXO;LSZC;Buochs (Buochs Airport);Switzerland
OEB;KOEB;Coldwater (Branch County Memorial Airport);United States
WBW;KWBW;Wilkes-Barre (Wilkes-Barre Wyoming Valley Airport);United States
LNN;KLNN;Willoughby (Lost Nation Municipal Airport);United States
UMD;BGUM;Uummannaq (Uummannaq Heliport);Greenland
FFT;KFFT;Frankfort (Capital City Airport);United States
LEW;KLEW;Lewiston (Lewiston Maine);United States
NBU;KNBU;Glenview (Naval Air Station);United States
MRK;KMKY;Marco Island Airport (Marco Islands);United States
DRM;KDRM;Drummond Island (Drummond Island Airport);United States
GDW;KGDW;Gladwin (Gladwin Zettel Memorial Airport);United States
LWA;KLWA;South Haven (South Haven Area Regional Airport);United States
MFI;KMFI;Marshfield (Marshfield Municipal Airport);United States
ISW;KISW;Wisconsin Rapids (Alexander Field South Wood County Airport);United States
CWI;KCWI;Clinton (Clinton Municipal);United States
BVY;KBVY;Beverly (Beverly Municipal Airport);United States
YRQ;CYRQ;Trois Rivieres (Trois Rivieres Airport);Canada
POF;KPOF;Poplar Bluff (Poplar Bluff Municipal Airport);United States
EPM;KEPM;Eastport (Eastport Municipal Airport);United States
EOK;KEOK;Keokuk (Keokuk Municipal Airport);United States
PSL;EGPT;Perth (Perth Scone Airport);United Kingdom
SOO;ESCL;Soderhamn (Soderhamn Airport);Sweden
VNA;VLSV;Saravane (Saravane Airport);Laos
DKS;UODD;Dikson (Dikson Airport);Russia
BYT;EIBN;Bantry (Bantry Aerodrome);Ireland
ADY;FAAL;Alldays (Alldays Airport);South Africa
GAS;HKGA;Garissa;Kenya
HOA;HKHO;Hola;Kenya
KLK;HYFG;Kalokol;Kenya
KEY;HKKR;Kericho;Kenya
ILU;HKKL;Kilaguni;Kenya
ATJ;FMME;Antsirabe;Madagascar
OVA;FMSL;Bekily;Madagascar
RGK;UNBG;Gorno-Altaysk (Gorno-Altaysk Airport);Russia
FLD;KFLD;Fond du Lac (Fond Du Lac County Airport);United States
PCZ;KPCZ;Waupaca (Waupaca Municipal Airport);United States
STE;KSTE;Stevens Point (Stevens Point Municipal Airport);United States
ERY;KERY;Newberry (Luce County Airport);United States
PEF;EDCP;Peenemunde (Peenemunde Airfield);Germany
GQQ;KGQQ;Galion (Galion Municipal Airport);United States
TPN;SETI;Tiputini;Ecuador
PTZ;SESM;Pastaza (Shell Mera);Ecuador
CKV;KCKV;Clarksville (Clarksville-Montgomery County Regional Airport);United States
LPC;KLPC;Lompoc (Lompoc Airport);United States
CTH;KMQS;Coatesville (Chester County G O Carlson Airport);United States
BST;OABT;Lashkar Gah (Bost Airport);Afghanistan
LLK;UBBL;Lankaran (Lankaran International Airport);Azerbaijan
GBB;UBBQ;Qabala (Qabala Airport);Azerbaijan
ZTU;UBBY;Zaqatala (Zaqatala International Airport);Azerbaijan
LKP;KLKP;Lake Placid (Lake Placid Airport);United States
DEE;YXCM;Yuzhno-Kurilsk (Mendeleevo);Russia
AOH;KAOH;Lima (Lima Allen County Airport);United States
SSI;KSSI;Brunswick (McKinnon Airport);United States
BFP;KBFP;Beaver Falls;United States
GGE;KGGE;Georgetown (Georgetown County Airport);United States
HDI;KHDI;Cleveland (Hardwick Field Airport);United States
RNT;KRNT;Renton;United States
ZBM;CZBM;Bromont (Bromont Airport);Canada
POC;KPOC;La Verne (Brackett Field);United States
CDK;KCDK;Cedar Key (CedarKey);United States
CTY;KCTY;Cross City;United States
CEU;KCEU;Clemson;United States
BEC;KBEC;Wichita (Beech Factory Airport);United States
GTU;KGTU;Georgetown (Georgetown Municipal Airport);United States
QFO;EGSU;Duxford (Duxford Aerodrome);United Kingdom
SNY;KSNY;Sidney (Sidney Muni Airport);United States
GKL;YGKL;Great Keppel Island;Australia
RPB;YRRB;Roper Bar;Australia
IFL;YIFL;Innisfail;Australia
BIN;OABN;Bamyan (Bamyan Airport);Afghanistan
RGO;ZZ07;Chongjin (Chongjin Airport);North Korea
MOO;YOOM;Moomba;Australia
LUZ;EPLB;Lublin;Poland
JYO;KJYO;Leesburg (Leesburg Executive Airport);United States
VAM;VRMV;Maamigili (Maamigili Airport);Maldives
IRU;IRUF;Maldives Hilton Iru fushi (Hilton Iru fushi);Maldives
DHG;DHGU;Dhigurah (Dhigurah Centara Grand Maldives);Maldives
LSZ;LDLO;Mali Losinj (Losinj Airport);Croatia
ONS;YOLW;Onslow (Onslow );Australia
TDR;YTDR;Theodore;Australia
CZW;EPRU;RUDNIKI (RUDNIKI );Poland
SDC;KSDC;Williamson (Williamson-Sodus Airport);United States
CFX;SASC;Cafayate (Gilberto Lavaque);Argentina
WBU;KBDU;Boulder (Boulder Municipal);United States
TVS;ZBSN;Tangshan (Sannvhe);China
PAO;KPAO;Palo Alto (Palo Alto Airport of Santa Clara County);United States
FFZ;KFFZ;Mesa (Mesa Falcon Field);United States
P08;KP08;Cooldige (Coolidge Municipal Airport);United States
P52;KP52;Cottonwood (Cottonwood Airport);United States
A39;KA39;Phoenix (Phoenix Regional Airport);United States
E25;KE25;Wickenburg (Wickenburg Municipal Airport);United States
YTY;ZSYA;Yangzhou (Yangzhou Taizhou Airport);China
PTK;KPTK;Pontiac (Oakland Co. Intl);United States
KSI;GUKU;Kissidougou;Guinea
THQ;ZLTS;Tianshui (Tianshui Airport);China
GKK;VRMO;Kooddoo;Maldives
RCS;EGTO;Rochester (Rochester Airport);United Kingdom
KNN;GUXD;Kankan;Guinea
RHD;SANH;Rio Hondo (Termal);Argentina
KMP;FYKT;Keetmanshoop;Namibia
KGT;ZUKD;Kangding (Kangding Airport);China
VUS;ULWU;Veliky Ustyug;Russia
IOW;KIOW;Iowa City (Iowa City Municipal Airport);United States
TLQ;ZWTP;Turpan;China
MWM;KMWM;Windom (Windom Municipal Airport);United States
ANP;KANP;Annapolis (Lee Airport);United States
DUU;HTND;Ndutu;Tanzania
FXO;FQCB;Cuamba;Mozambique
ODO;UIKB;Bodaibo;Russia
ZTR;UKKV;Zhytomyr;Ukraine
HRI;VCRI;Mattala (Mattala Rajapaksa Intl.);Sri Lanka
PEQ;KPEQ;Pecos (Pecos Municipal Airport);United States
HBG;KHBG;Hattiesburg (Hattiesburg Bobby L. Chain Municipal Airport);United States
QCJ;SDBK;Botucatu;Brazil
QSC;SDSC;Sao Carlos (Sao Carlos TAM);Brazil
YKN;KYKN;Yankton (Chan Gurney);United States
SES;KSES;Selma Alabama (Selfield Airport);United States
KTI;VDKT;Kratie (Kratie Airport);Cambodia
GYU;ZLGY;Guyuan;China
CNI;ZYCH;Changhai;China
KRH;EGKR;Redhill (Redhill Aerodrome);United Kingdom
CCL;YCCA;Chinchilla;Australia
HWD;KHWD;Hayward (Hayward Executive Airport);United States
MZP;NZMK;Motueka;New Zealand
JHQ;YSHR;Shute Harbour;Australia
ARB;KARB;Ann Arbor (Ann Arbor Municipal Airport);United States
SHT;YSHT;Shepparton;Australia
TEM;YTEM;Temora;Australia
GAH;YGAY;Gayndah;Australia
WIO;YWCA;Wilcannia;Australia
ULK;UERL;Lensk;Russia
IGD;LTCT;Igdir;Turkey
GNY;LTCS;Sanliurfa (Sanliurfa GAP);Turkey
KZR;LTBZ;Kutahya (Zafer);Turkey
VLU;ULOL;Velikiye Luki;Russia
YLK;VOYK;Bangalore (Yelahanka AFB);India
BEO;YPEC;Lake Macquarie (Belmont Airport);Australia
4A7;K4A7;Hampton (Clayton County Tara Field);United States
BMP;YBPI;Brampton Island;Australia
NGZ;KNGZ;Alameda (NAS Alameda);United States
YCN;CYCN;Cochrane;Canada
BJP;SBBP;Braganca Paulista (Aeroporto Estadual Arthur Siqueira);Brazil
BQB;YBLN;Brusselton;Australia
IVR;YIVL;Inverell;Australia
GLI;YGLI;Glen Innes;Australia
IMM;KIMM;Immokalee ;United States
YIC;ZSYC;Yichun (Yichun Mingyueshan Airport);China
PTB;KPTB;Petersburg (Dinwiddie County Airport);United States
KGN;FZOK;Kasongo (Kasongo Lunda);Congo (Kinshasa)
SBM;KSBM;Sheboygan (Sheboygan County Memorial Airport);United States
KFE;YFDF;Cloudbreak (Dave Forest Airport);Australia
BJU;VNBR;Bajura (Bajura Airport);Nepal
MZJ;KMZJ;Marana (Pinal Airpark);United States
GEU;KGEU;Glendale (Glendale Municipal Airport);United States
SAD;KSAD;Safford (Safford Regional Airport);United States
SLJ;YSOL;Solomon (Solomon Airport);Australia
KJP;ROKR;Kerama (Kerama Airport);Japan
SIK;KSIK;Sikeston (Sikeston Memorial Municipal);United States
L52;KL52;Oceano (Oceano County Airport);United States
TTI;NTTE;Tetiaroa (Tetiaroa Airport);French Polynesia
GFL;KGFL;Queensbury (Floyd Bennett Memorial Airport);United States
5B2;K5B2;Ballston Spa (Saratoga County Airport);United States
MJU;WAWJ;Mamuju (Tampa Padang);Indonesia
CGC;KCGC;Crystal River;United States
MTN;KMTN;Baltimore (Martin State);United States
LHM;KLHM;Lincoln (Lincoln Regional Airport Karl Harder Field);United States
FZI;KFZI;Fostoria (Fostoria Metropolitan Airport);United States
IZG;KIZG;Fryeburg (Eastern Slopes Regional);United States
NEW;KNEW;New Orleans (Lakefront);United States
COE;KCOE;Coeur d'Alene (Pappy Boyington);United States
BMT;KBMT;Beaumont (Beaumont Municipal);United States
DNV;KDNV;Danville (Vermilion Regional);United States
COJ;YCBB;Coonabarabran;Australia
NYE;HKNI;NYERI;Kenya
1H2;KEFH;Effingham (Effingham Memorial Airport);United States
AAP;KAAP;Houston (Andrau Airport);United States
FCM;KFCM;Eden Prairie (Flying Cloud Airport);United States
LIX;FWLK;Likoma Island (Likoma Island Airport);Malawi
OJC;KOJC;Olathe (Johnson County Airport);United States
GIU;VCCS;Sigiriya (Sigiriya Airport);Sri Lanka
EUM;EDHN;Neumuenster;Germany
TKT;VTPT;Tak;Thailand
YLS;CYLS;Barrie-Orillia (Lake Simcoe);Canada
YEE;CYEE;Midland (Huronia);Canada
NU8;CNU8;Markham;Canada
ND4;CND4;Haliburton (Stanhope);Canada
NF4;CNF4;Lindsay;Canada
YCM;CYSN;Saint Catherines (Niagara District);Canada
XCM;CNZ3;Chatham (Kent);Canada
YPD;CNK4;Parry Sound;Canada
OQN;KOQN;West Goshen Township (Brandywine Airport);United States
MNZ;KHEF;Manassas;United States
BGG;LTCU;Bingol;Turkey
KFS;LTAL;Kastamonu (Uzunyazi);Turkey
2H0;K2H0;Shelbyville (Shelby County Airport);United States
DCY;ZUDC;Daocheng (Yading Daocheng);China
GXH;ZLXH;Xiahe city (Gannan);China
CIY;LICB;Comiso;Italy
KVM;UHMO;Markovo (Markovo Airport);Russia
ZKP;UESU;Zyryanka (Zyryanka West Airport);Russia
UMS;UEMU;Ust-Maya (Ust-Maya Airport);Russia
ADH;UEEA;Aldan (Aldan Airport);Russia
NLT;ZWNL;Xinyuan (Nalati);China
PTA;PALJ;Port alsworth (Port Alsworth Airport);United States
BOR;LFSQ;Belfort (Fontaine Airport);France
FDW;KFDW;Winnsboro (Fairfield County Airport);United States
OBC;HDOB;Obock;Djibouti
TDJ;HDTJ;Tadjoura;Djibouti
AQB;MGQC;Santa Cruz des Quiche (Santa Cruz des Quiche Airport);Guatemala
NOR;BINF;Nordfjordur (Nordfjordur Airport);Iceland
BTZ;LTBE;Bursa (Bursa Airport);Turkey
DAW;KDAW;Rochester (Skyhaven Airport);United States
WAR;WAJR;Waris-Papua Island (Waris Airport);Indonesia
EWK;KEWK;Newton (Newton City-County Airport);United States
BSJ;YBNS;Bairnsdale (Bairnsdale Airport);Australia
TZR;KTZR;Columbus (Bolton Field);United States
W04;KW04;Ocean Shores (Ocean Shores Municipal);United States
55S;K55S;Packwood;United States
FBR;KFBR;Fort Bridger;United States
S40;KS40;Prosser;United States
CLS;KCLS;Chehalis (Chehalis-Centralia);United States
M94;KM94;Mattawa (Desert Aire);United States
S30;KS30;Lebanon (Lebanon State);United States
EVW;KEVW;Evanston (Evanston-Uinta CO Burns Fld);United States
K83;KK83;Sabetha (Sabetha Municipal);United States
LRO;KLRO;Mount Pleasant (Mount Pleasant Regional-Faison Field);United States
ACJ;KACJ;Americus (Jimmy Carter Regional);United States
EUF;KEUF;Eufala (Weedon Field);United States
6J4;K6J4;Saluda (Saluda County);United States
MQI;KMQI;Manteo (Dare County Regional);United States
AUO;KAUO;Auburn (Auburn University Regional);United States
CZG;KCZG;Endicott (Tri-Cities);United States
EKY;KEKY;Bessemer;United States
A50;KA50;Ellicott (Colorado Springs East);United States
ИКУ;UCFL;Tamchy (Issyk-Kul International Airport);Kyrgyzstan
MIC;KMIC;Crystal (Crystal Airport);United States
23M;K23M;Quitman (Clarke CO);United States
DBN;KDBN;Dublin (W H Barron Field);United States
PUK;NTGQ;Pukarua (Pukarua Airport);French Polynesia
CVO;KCVO;Corvallis (Corvallis Muni);United States
PXH;YPMH;Prominent Hill;Australia`
	
	return adlist, help
}
