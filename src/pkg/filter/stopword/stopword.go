package stopword

import (
    "fetch/filter"
    "fetch/tokenizer"
)

type Stopword struct {}

func (a *Stopword) Process(in tokenizer.TokenChan) tokenizer.TokenChan {
    // http://armandbrahaj.blog.al/2009/04/14/list-of-english-stop-words/
    words := map[string]bool{"toward":true,"right":true,"everyone":true,"sometimes":true,"yourselves":true,"overall":true,"associated":true,"goes":true,"whos":true,"were":true,"than":true,"arent":true,"first":true,"inc":true,"quite":true,"enough":true,"my":true,"off":true,"non":true,"aint":true,"best":true,"are":true,"only":true,"hers":true,"elsewhere":true,"usually":true,"theyve":true,"howbeit":true,"gets":true,"someone":true,"whole":true,"theyll":true,"becomes":true,"latter":true,"another":true,"new":true,"nor":true,"that":true,"keep":true,"used":true,"take":true,"merely":true,"indicate":true,"had":true,"not":true,"but":true,"given":true,"some":true,"need":true,"am":true,"therein":true,"an":true,"several":true,"rather":true,"now":true,"taken":true,"against":true,"having":true,"twice":true,"get":true,"via":true,"unlikely":true,"there":true,"entirely":true,"also":true,"saying":true,"just":true,"gives":true,"anyways":true,"as":true,"else":true,"at":true,"asking":true,"known":true,"whither":true,"thus":true,"downwards":true,"despite":true,"cannot":true,"might":true,"further":true,"outside":true,"together":true,"except":true,"got":true,"among":true,"ones":true,"apart":true,"concerning":true,"somebody":true,"nothing":true,"un":true,"serious":true,"hardly":true,"available":true,"knows":true,"tried":true,"thereupon":true,"nevertheless":true,"able":true,"whether":true,"up":true,"our":true,"onto":true,"beyond":true,"has":true,"regarding":true,"uses":true,"everybody":true,"selves":true,"six":true,"until":true,"following":true,"out":true,"seeing":true,"plus":true,"come":true,"zero":true,"us":true,"says":true,"most":true,"although":true,"sensible":true,"either":true,"which":true,"wed":true,"thats":true,"above":true,"looking":true,"hopefully":true,"corresponding":true,"eight":true,"once":true,"indeed":true,"id":true,"done":true,"hasnt":true,"whereby":true,"using":true,"too":true,"ie":true,"same":true,"wish":true,"hence":true,"one":true,"if":true,"never":true,"comes":true,"according":true,"nd":true,"necessary":true,"probably":true,"ive":true,"thoroughly":true,"such":true,"many":true,"others":true,"this":true,"seems":true,"greetings":true,"ill":true,"often":true,"tries":true,"theirs":true,"behind":true,"com":true,"wants":true,"aside":true,"followed":true,"seeming":true,"whatever":true,"edu":true,"do":true,"would":true,"viz":true,"im":true,"sure":true,"youre":true,"in":true,"shall":true,"gotten":true,"been":true,"ours":true,"thereby":true,"thereafter":true,"the":true,"hereupon":true,"tell":true,"other":true,"maybe":true,"seemed":true,"think":true,"dont":true,"no":true,"example":true,"whoever":true,"theres":true,"is":true,"came":true,"what":true,"both":true,"it":true,"so":true,"namely":true,"shouldnt":true,"wouldnt":true,"therefore":true,"less":true,"every":true,"immediate":true,"him":true,"definitely":true,"lest":true,"wasnt":true,"unfortunately":true,"anyway":true,"havent":true,"since":true,"noone":true,"considering":true,"lately":true,"beside":true,"and":true,"per":true,"wheres":true,"theyre":true,"must":true,"be":true,"his":true,"see":true,"appropriate":true,"few":true,"herself":true,"besides":true,"from":true,"later":true,"yours":true,"hereafter":true,"actually":true,"hither":true,"for":true,"old":true,"respectively":true,"under":true,"reasonably":true,"relatively":true,"happens":true,"awfully":true,"will":true,"itd":true,"ignored":true,"where":true,"specify":true,"placed":true,"beforehand":true,"allows":true,"okay":true,"contain":true,"specified":true,"two":true,"towards":true,"consequently":true,"keeps":true,"ask":true,"know":true,"use":true,"currently":true,"did":true,"become":true,"que":true,"should":true,"regards":true,"ought":true,"anyhow":true,"follows":true,"itself":true,"below":true,"alone":true,"went":true,"name":true,"who":true,"whereafter":true,"werent":true,"these":true,"perhaps":true,"go":true,"along":true,"yourself":true,"herein":true,"wont":true,"very":true,"try":true,"nine":true,"kept":true,"gone":true,"wonder":true,"those":true,"became":true,"everything":true,"getting":true,"presumably":true,"provides":true,"welcome":true,"ltd":true,"formerly":true,"nobody":true,"whats":true,"took":true,"even":true,"hello":true,"value":true,"their":true,"sent":true,"near":true,"any":true,"by":true,"last":true,"its":true,"useful":true,"tends":true,"always":true,"help":true,"appear":true,"whereupon":true,"ever":true,"could":true,"why":true,"unto":true,"truly":true,"doing":true,"seriously":true,"changes":true,"inner":true,"vs":true,"unless":true,"themselves":true,"mainly":true,"however":true,"them":true,"let":true,"said":true,"amongst":true,"qv":true,"weve":true,"thence":true,"then":true,"better":true,"myself":true,"obviously":true,"can":true,"etc":true,"she":true,"well":true,"away":true,"though":true,"eg":true,"afterwards":true,"almost":true,"believe":true,"causes":true,"with":true,"somewhat":true,"self":true,"least":true,"all":true,"please":true,"described":true,"cmon":true,"meanwhile":true,"around":true,"somehow":true,"nearly":true,"instead":true,"more":true,"of":true,"containing":true,"inward":true,"everywhere":true,"willing":true,"whence":true,"brief":true,"oh":true,"hadnt":true,"regardless":true,"insofar":true,"want":true,"does":true,"consider":true,"different":true,"they":true,"ok":true,"looks":true,"indicated":true,"particular":true,"look":true,"th":true,"youd":true,"was":true,"novel":true,"still":true,"itll":true,"thank":true,"secondly":true,"have":true,"thorough":true,"on":true,"contains":true,"forth":true,"across":true,"yes":true,"something":true,"may":true,"really":true,"et":true,"none":true,"yet":true,"ourselves":true,"accordingly":true,"over":true,"way":true,"to":true,"throughout":true,"far":true,"into":true,"or":true,"likely":true,"former":true,"lets":true,"ex":true,"didnt":true,"hereby":true,"clearly":true,"especially":true,"whereas":true,"her":true,"you":true,"ts":true,"trying":true,"hes":true,"exactly":true,"already":true,"appreciate":true,"neither":true,"before":true,"inasmuch":true,"wherever":true,"your":true,"within":true,"thanx":true,"indicates":true,"mostly":true,"moreover":true,"while":true,"he":true,"sorry":true,"four":true,"like":true,"liked":true,"anybody":true,"much":true,"mean":true,"otherwise":true,"furthermore":true,"anyone":true,"me":true,"certain":true,"being":true,"couldnt":true,"whenever":true,"possible":true,"after":true,"sub":true,"thanks":true,"hi":true,"rd":true,"seem":true,"specifying":true,"how":true,"whose":true,"cant":true,"re":true,"somewhere":true,"next":true,"seen":true,"fifth":true,"isnt":true,"certainly":true,"co":true,"we":true,"third":true,"five":true,"particularly":true,"normally":true,"here":true,"when":true,"theyd":true,"cs":true,"becoming":true,"course":true,"youve":true,"thru":true,"about":true,"youll":true,"without":true,"allow":true,"down":true,"saw":true,"going":true,"nowhere":true,"himself":true,"anywhere":true,"seven":true,"soon":true,"each":true,"say":true,"second":true,"because":true,"own":true,"whom":true,"between":true,"needs":true,"various":true,"upon":true,"during":true,"cause":true,"sup":true,"three":true,"sometime":true,"heres":true,"anything":true,"latterly":true,"little":true,"wherein":true,"through":true,"doesnt":true,"again":true,"a":true,"put":true,"i":true}

    return filter.BuildFilter(in, func (t *tokenizer.Token) *tokenizer.Token {
        if words[t.Backing()] {
            return nil
        }
        return t
    })
}

func Build() *Stopword {
    return new(Stopword)
}
