<template>
  <b-container
    class="my-5">
    <b-row align-h="center">
      <b-col v-for="(article, i) in news" :key="i"
        class="my-2"
        :md="(article.important ? '12' : '6')"
        :lg="(article.important ? '6' : '3')">
        <news-card :article="article" />
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
  import NewsCard from '@/components/NewsCard.vue';

  export default {
    components: {
      NewsCard
    },
    data() {
      return {
        news: []
      }
    },
    mounted() {
      let data = `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" media="screen" href="/~d/styles/rss2full.xsl"?><?xml-stylesheet type="text/css" media="screen" href="http://feeds.ign.com/~d/styles/itemcontent.css"?><rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:media="http://search.yahoo.com/mrss/" version="2.0">
  <channel>
    <title>IGN E3 Videos</title>
    <link>http://www.ign.com/videos</link>
    <description>IGN has the latest press conferences, video game videos, game trailers, gameplay clips and interviews from E3</description>
    <atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="self" type="application/rss+xml" href="http://feeds.ign.com/ign/e3-videos" /><feedburner:info xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" uri="ign/e3-videos" /><atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="hub" href="http://pubsubhubbub.appspot.com/" /><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="https://add.my.yahoo.com/rss?url=http%3A%2F%2Ffeeds.ign.com%2Fign%2Fe3-videos" src="http://us.i1.yimg.com/us.yimg.com/i/us/my/addtomyyahoo4.gif">Subscribe with My Yahoo!</feedburner:feedFlare><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="http://fusion.google.com/add?feedurl=http%3A%2F%2Ffeeds.ign.com%2Fign%2Fe3-videos" src="http://buttons.googlesyndication.com/fusion/add.gif">Subscribe with Google</feedburner:feedFlare><item>
      <title>Fireteam Chat Ep. 209 - Destiny 2 Loot vs. Borderlands Loot and Our E3 Hopes - IGN's Destiny Show</title>
      <link>http://www.ign.com/videos/2019/05/04/fireteam-chat-ep-209-destiny-2-loot-vs-borderlands-loot-and-our-e3-hopes-igns-destiny-show</link>
      <description>After the Borderlands 3 reveal loot is on all of our minds and we discuss what Destiny does well vs. the Gearbox shooter. Then we chat a bit about our E3 2019 hopes for Destiny.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/05/03/640/3166ed11f9680209cd2e8fdc2ca61967-347000-1556918058.mp4" type="video/mp4" />
      <pubDate>Sat, 04 May 2019 00:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/05/04/fireteam-chat-ep-209-destiny-2-loot-vs-borderlands-loot-and-our-e3-hopes-igns-destiny-show</guid>
      <media:content bitrate="346000.0" duration="2341" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/05/03/640/3166ed11f9680209cd2e8fdc2ca61967-347000-1556918058.mp4">
        <media:keywords>destiny 2, borderlands 3, borderlands vs destiny, borderlands, e3, loot, destiny loot, weapons, Talkshow, PC, Destiny 2, games, Shooter, Activision, Bungie Software, Xbox One, PS4, Destiny 2 – Expansion III: Forsaken, DLC / Expansion</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/5/3/destiny209-1556918121103.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-05-04T00:00:00Z</dc:date>
    </item>
    <item>
      <title>Square Enix Announced E3 Conference Details - IGN News</title>
      <link>http://www.ign.com/videos/2019/04/23/square-enix-announced-e3-conference-details</link>
      <description>Final Fantasy VII Remake? The Avengers Project? Square Enix announces its plans for E3 2019.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/04/23/640/ee2196da8e1f1abbc5c1786878e4922b-347000-1556040828.mp4" type="video/mp4" />
      <pubDate>Tue, 23 Apr 2019 17:31:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/04/23/square-enix-announced-e3-conference-details</guid>
      <media:content bitrate="341000.0" duration="55" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/04/23/640/ee2196da8e1f1abbc5c1786878e4922b-347000-1556040828.mp4">
        <media:keywords>games, gaming, ign, ign news, breaking news, ps4, xbox one, nintendo, video games, Feature, switch, nintendo switch, News, e3, e32019, e3 2019, Square Enix, companies, PlayStation 4, Final Fantasy VII Remake</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/4/23/042319-squaree3thumb-1556040709495.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-04-23T17:31:00Z</dc:date>
    </item>
    <item>
      <title>Best Cosplay of PAX East 2019</title>
      <link>http://www.ign.com/videos/2019/04/03/best-cosplay-of-pax-east-2019</link>
      <description>PAX East 2019 is over, but don't be sad! We have a round up of some of the best cosplay we found while on floor.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/04/02/640/ba8c1a859515c1ebef7d2adff6303092-347000-1554254802.mp4" type="video/mp4" />
      <pubDate>Wed, 03 Apr 2019 13:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/04/03/best-cosplay-of-pax-east-2019</guid>
      <media:content bitrate="346000.0" duration="119" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/04/02/640/ba8c1a859515c1ebef7d2adff6303092-347000-1554254802.mp4">
        <media:keywords>overwatch, borderlands, borderlands 3, cosplay, costume, jessica nigri, Sneaky Zebra, PAX East, IGN Access, Pax east 2019, pax, pax 2019, best cosplay, top cosplayers, best cosplayers, anthem cosplay, top cosplay, meg turney, anime expo, san diego comic con, new york comic con, wondercon, master chief, smash bros cosplay, dragon ball z cosplay, cosplay girls, pax south, pax west, pax aus, penny arcade, cosplay compilation, cosplay tips, c2e2, e3, cosplay makeup, blizzcon, sdcc, nycc, Feature, PAXEast, Overwatch, episodes, The Legend of Zelda, comic, Manga, Viz, PC, League of Legends, games, Action, MOBA (Multiplayer Online Battle Arena), Riot Games, THQ, Mac, PS4, Death Stranding, Sony Computer Entertainment, Kojima Productions, GBA, Mario Bros., Platformer, Mario, Nintendo, Arcade, Nintendo R&amp;D1, Wii, Apple-II, Atarisoft (PC division), 400/800, Atari Inc. (1972-1984), NES, Nintendo R&amp;D4, 7800, Atari Corp. (1984-2003), Wii U, Switch, 2600, 3DS, 5200, Hamster Corporation, God of War, Adventure, Santa Monica Studio</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/4/2/bestcosplaypax2019blogroll-1554254084230.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-04-03T13:00:00Z</dc:date>
    </item>
    <item>
      <title>WTF PAX?! The 5 Weirdest Games at Pax East 2019</title>
      <link>http://www.ign.com/videos/2019/04/01/wtf-pax-the-5-weirdest-games-at-pax-east-2019</link>
      <description>Pax East 2019 is chock full o' every type of games your heart desires! BUT Sydnee Goodman found her favorite bonkers, different, and just plain off-the-beaten video games on the convention floor.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/04/01/640/0832c3ba3e3303a2fd2b9ef6529596d5-347000-1554114234.mp4" type="video/mp4" />
      <pubDate>Mon, 01 Apr 2019 13:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/04/01/wtf-pax-the-5-weirdest-games-at-pax-east-2019</guid>
      <media:content bitrate="343000.0" duration="252" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/04/01/640/0832c3ba3e3303a2fd2b9ef6529596d5-347000-1554114234.mp4">
        <media:keywords>wtf games, weirdest video games ever, wtf video games, grossest video games, video game fail, unique video games, best video games, indie games, best indie games 2019, speaking simulator, simulator games, pax 2019, pax, e3, what the golf, dating sim, dating simulator, later daters, hatoful dating sim, japanese game, japanese video game, wtf japan, buildings have feelings too, nintendo switch, city management game, sim city, sims, merge games, freedom finger, wide right interactive, steam, steam games, mobile games, Let's Play, Feature, PAXEast, PC, Freedom Finger, games, Shooter, Music, SHMUP, Wide Right Interactive, PS4, Xbox One, Speaking Simulator, Simulation, Affable Games, Buildings Have Feelings Too, Merge Games, Black Staff Games, Switch</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/3/31/wtf-paxblogroll-1554079910914.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-04-01T13:00:00Z</dc:date>
    </item>
    <item>
      <title>2 Minutes of PAX 2019 Cosplay!</title>
      <link>http://www.ign.com/videos/2019/03/30/2-minutes-of-pax-2019-cosplay</link>
      <description>We searched the convention center floor to find the best, craziest, and coolest cosplay we could find! Check out these amazing cosplayers who hit the floor day 1 of Pax East 2019.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/03/30/640/6560d8e4a1c2ac05fcb1c18a5c5af717-347000-1553961904.mp4" type="video/mp4" />
      <pubDate>Sat, 30 Mar 2019 17:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/03/30/2-minutes-of-pax-2019-cosplay</guid>
      <media:content bitrate="346000.0" duration="104" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/03/30/640/6560d8e4a1c2ac05fcb1c18a5c5af717-347000-1553961904.mp4">
        <media:keywords>borderlands 3, cosplay. costume, jessica nigri, Sneaky Zebra, PAX East, IGN Access, Pax east 2019, pax, cosplay contest, best cosplay, top cosplayers, best cosplayers, meg turney, anime expo, san diego comic con, new york comic con, wondercon, master chief, halo cosplay, dragon ball z cosplay, cosplay girls, cosplay guys, pax south, pax west, pax aus, penny arcade, cosplay posing tips, cosplay compilation, cosplay tips, c2e2, e3, cosplay makeup, blizzcon, cosplay wig, sdcc, nycc, disney cosplay, none, PAXEast</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/3/30/pax2019-cosplayblogroll-1553957467324.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-03-30T17:00:00Z</dc:date>
    </item>
    <item>
      <title>Everything at PAX East 2019 in 3 Minutes!</title>
      <link>http://www.ign.com/videos/2019/03/30/everything-at-pax-east-2019-in-3-minutes</link>
      <description>PAX East has it all:  tabletop gaming, arcade gaming, video gaming, and more- don't believe us? Fine, Sydnee Goodman will show you EVERYTHING!</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/03/29/640/b19c3b4e682328ab857fc51786b3a0c1-347000-1553914342.mp4" type="video/mp4" />
      <pubDate>Sat, 30 Mar 2019 13:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/03/30/everything-at-pax-east-2019-in-3-minutes</guid>
      <media:content bitrate="343000.0" duration="180" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/03/29/640/b19c3b4e682328ab857fc51786b3a0c1-347000-1553914342.mp4">
        <media:keywords>PAX East, IGN Access, Pax east 2019, pax, pax 2019, nintendo, cosplay, occulus, oculus, tabletop, gameplay, pax announcement, sydnee goodman, ign, arcade gaming, pax boston, days gone, e3, gdc, none, PAXEast, Nintendo, companies, PS3, Tekken, games, Fighting, Sony Computer Entertainment, Namco, Arcade, Unknown, PSP, PS, Fillpoint, Switch, Nintendo Switch, Hardware, PS4, Days Gone, Adventure, Bend Studio</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/3/29/ignaccess-exploringpaxeastblogroll-1553901936295.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-03-30T13:00:00Z</dc:date>
    </item>
    <item>
      <title>Bethesda Won't Discuss Elder Scrolls 6, Starfield at E3 2019 - IGN News</title>
      <link>http://www.ign.com/videos/2019/03/29/bethesda-wont-discuss-elder-scrolls-6-starfield-at-e3-2019</link>
      <description>Bethesda has confirmed that it will not be discussing either The Elder Scrolls 6 or Starfield at E3 2019.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/03/29/640/7751d4cc3fd74bb87e0898a2562d95b3-347000-1553901157.mp4" type="video/mp4" />
      <pubDate>Fri, 29 Mar 2019 23:08:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/03/29/bethesda-wont-discuss-elder-scrolls-6-starfield-at-e3-2019</guid>
      <media:content bitrate="343000.0" duration="55" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/03/29/640/7751d4cc3fd74bb87e0898a2562d95b3-347000-1553901157.mp4">
        <media:keywords>games, gaming, ign, ign news, breaking news, ps4, xbox one, nintendo, video games, Feature, switch, nintendo switch, News, E3, e3 2019, Bethesda Game Studios, companies, TBA, The Elder Scrolls VI, Starfield, PAX East</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/3/29/032919-bethesdae3thumb-1553900966445.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-03-29T23:08:00Z</dc:date>
    </item>
    <item>
      <title>What Can Sony Learn From Nintendo Direct and Inside Xbox? - Beyond Highlight</title>
      <link>http://www.ign.com/videos/2019/02/15/what-can-sony-learn-from-nintendo-direct-and-inside-xbox-beyond-highlight</link>
      <description>Sony is pulling out of E3, but what lessons about messaging can they take from their biggest competitors?</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/02/14/640/86f18f77ce3c26c1ec54bdbc94c448bd-347000-1550173086.mp4" type="video/mp4" />
      <pubDate>Fri, 15 Feb 2019 20:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/02/15/what-can-sony-learn-from-nintendo-direct-and-inside-xbox-beyond-highlight</guid>
      <media:content bitrate="344000.0" duration="393" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/02/14/640/86f18f77ce3c26c1ec54bdbc94c448bd-347000-1550173086.mp4">
        <media:keywords>shawn, layden, shawn layden, transparency, PS4, nintendo direct, inside xbox, e3, sony e3, Talkshow, Switch, Nintendo Switch, games, Hardware, Nintendo, Xbox One, Xbox One X, Microsoft</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/2/15/beyond-whatsonycanlearn-1550258367133.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-02-15T20:00:00Z</dc:date>
    </item>
    <item>
      <title>How Can Sony Control the Conversation Without E3? - Beyond 578</title>
      <link>http://www.ign.com/videos/2019/02/13/how-can-sony-control-the-conversation-without-e3-beyond-578</link>
      <description>IGN's PlayStation show discusses Shawn Layden's comments regarding E3, as well as Apex Legends, and whether or not too many games are coming out this week. Beyond!</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2019/02/13/640/54afb8b4f05adb2104b83bed2507dc47-347000-1550087890.mp4" type="video/mp4" />
      <pubDate>Wed, 13 Feb 2019 23:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2019/02/13/how-can-sony-control-the-conversation-without-e3-beyond-578</guid>
      <media:content bitrate="346000.0" duration="3313" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2019/02/13/640/54afb8b4f05adb2104b83bed2507dc47-347000-1550087890.mp4">
        <media:keywords>e3, nintendo direct, direct, inside xbox, layden, shawn layden, Talkshow, PS4, Apex Legends, games, Shooter, Battle Royale, Free-to-Play, Electronic Arts, Respawn Entertainment, Xbox One, PC, Resident Evil 2, Action, Adventure, Resident Evil, Metro Exodus, 4A Games, Deep Silver, Kingdom Hearts III, RPG, Kingdom Hearts, Square Enix, Far Cry: New Dawn, Ubisoft, Ubisoft Montreal, Ubisoft Shanghai, Ubisoft Kiev, Ubisoft Bucharest</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2019/2/13/beyondsonye3blogroll-1550085667462.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2019-02-13T23:00:00Z</dc:date>
    </item>
    <item>
      <title>Behind The Scenes at IGN Fan Fest 2018</title>
      <link>http://www.ign.com/videos/2018/12/15/ign-fan-fest-2018</link>
      <description>Jacki Jing parties it up in Downtown LA at IGN Fan Fest, playing Fortnite with YouTubers runJDrun and Sam Bashor, and a bunch of IGN fans.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/12/14/640/eed0298c1b3bd93cae22eb38eb1ff50d-347000-1544838045.mp4" type="video/mp4" />
      <pubDate>Sat, 15 Dec 2018 03:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/12/15/ign-fan-fest-2018</guid>
      <media:content bitrate="344000.0" duration="168" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/12/14/640/eed0298c1b3bd93cae22eb38eb1ff50d-347000-1544838045.mp4">
        <media:keywords>IGN, fortnite, early access, kingdom hearts 3, KH3, Just cause 4, dragon quest XI, Hitman 2, streamer showdown, fortnite tonite, fortnite news, fan fest, E3, 2018, Jacki Jing, Sam Bashor, rude gamer, only stupid answers, JD witherspoon, runJDrun, none, Feature, PS4, Fortnite, games, Action, Battle Royale, Epic Games, Epic Games -- Poland, Mac, PC, Switch, iPhone, Xbox One, Android</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/12/14/thumb2-1544831672689.png" width="640" height="360" />
      </media:content>
      <dc:date>2018-12-15T03:00:00Z</dc:date>
    </item>
    <item>
      <title>What PlayStation Skipping E3 Means for 2019</title>
      <link>http://www.ign.com/videos/2018/11/15/what-playstation-skipping-e3-means-for-2019</link>
      <description>Jonathon Dornbush and Max Scoville discuss PlayStation's big move to skip E3 next year and what that means for Sony and the yearly convention.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/11/15/640/c942db8f48659ea3a5f79834859c9416-347000-1542328868.mp4" type="video/mp4" />
      <pubDate>Thu, 15 Nov 2018 23:59:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/11/15/what-playstation-skipping-e3-means-for-2019</guid>
      <media:content bitrate="343000.0" duration="205" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/11/15/640/c942db8f48659ea3a5f79834859c9416-347000-1542328868.mp4">
        <media:keywords>PS5, PlayStation, E3 2019, E3, Video Games, Opinion, PS4, PlayStation 4 Pro, games, Hardware, Sony Computer Entertainment</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/11/15/36190303ps4-ignjpg-57aad11280wjpg-4856ff-1542326364125.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2018-11-15T23:59:00Z</dc:date>
    </item>
    <item>
      <title>19 Minutes of Anthem Gameplay (with Developer Commentary) in 4K</title>
      <link>http://www.ign.com/videos/2018/07/05/19-minutes-of-anthem-gameplay-with-developer-commentary-in-4k</link>
      <description>BioWare take us on a guided tour through the E3 2018 gameplay of Anthem from the Colossus Javelin.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/07/03/640/be9309970f7ac7c9f84edbc73f9de536-347000-1530671168.mp4" type="video/mp4" />
      <pubDate>Thu, 05 Jul 2018 16:29:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/07/05/19-minutes-of-anthem-gameplay-with-developer-commentary-in-4k</guid>
      <media:content bitrate="345000.0" duration="1165" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/07/03/640/be9309970f7ac7c9f84edbc73f9de536-347000-1530671168.mp4">
        <media:keywords>anthem, javelin, gameplay, e3, developer, bioware, Gameplay, PS4, Anthem, games, Shooter, RPG, Electronic Arts, BioWare, Xbox One, PC</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/7/3/javelin1-1530672238644.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2018-07-05T16:29:00Z</dc:date>
    </item>
    <item>
      <title>5 Minutes of Left 4 Dead-like Earthfall Gameplay on Xbox One X</title>
      <link>http://www.ign.com/videos/2018/07/02/5-minutes-of-left-4-dead-like-earthfall-gameplay-on-xbox-one-x</link>
      <description>Teaming up to defend a base in the Left 4 Dead-like Earthfall on Xbox One X.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/07/02/640/781549b0f9b2b85ba5caa48a8de69a40-347000-1530568286.mp4" type="video/mp4" />
      <pubDate>Mon, 02 Jul 2018 21:20:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/07/02/5-minutes-of-left-4-dead-like-earthfall-gameplay-on-xbox-one-x</guid>
      <media:content bitrate="345000.0" duration="310" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/07/02/640/781549b0f9b2b85ba5caa48a8de69a40-347000-1530568286.mp4">
        <media:keywords>e3, e3 2018, Gameplay, Xbox One, EarthFall, games, Shooter, Holospark, PC, PS4</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/7/2/earthfallthumb-1530566405327.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2018-07-02T21:20:00Z</dc:date>
    </item>
    <item>
      <title>For Honor: Marching Fire Shaolin Breach Gameplay on Walled City - E3 2018</title>
      <link>http://www.ign.com/videos/2018/07/02/for-honor-marching-fire-shaolin-breach-gameplay-on-walled-city-e3-2018</link>
      <description>Running into battle as the Shaolin in For Honor: Marching Fire.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/07/02/640/d2e5163425e8901163bcc21b79ddc7d6-500000-1530558630.mp4" type="video/mp4" />
      <pubDate>Mon, 02 Jul 2018 19:01:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/07/02/for-honor-marching-fire-shaolin-breach-gameplay-on-walled-city-e3-2018</guid>
      <media:content bitrate="502000.0" duration="513" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/07/02/640/d2e5163425e8901163bcc21b79ddc7d6-500000-1530558630.mp4">
        <media:keywords>e3, e3 2018, Gameplay, E3, PS4, For Honor: Marching Fire, games, Action, DLC / Expansion, Ubisoft, Ubisoft Montreal, PC, Xbox One, For Honor</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/7/2/shaolingp-1530558264603.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2018-07-02T19:01:00Z</dc:date>
    </item>
    <item>
      <title>13 Minutes of For Honor: Marching Fire Tiandi Breach Gameplay - E3 2018</title>
      <link>http://www.ign.com/videos/2018/07/02/13-minutes-of-for-honor-marching-fire-tiandi-breach-gameplay-e3-2018</link>
      <description>Taking down the lord as Tiandi in For Honor's upcoming Marching Fire release.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/07/02/640/6542fde3119b2b83b448dd0b54d7aecf-500000-1530558041.mp4" type="video/mp4" />
      <pubDate>Mon, 02 Jul 2018 19:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/07/02/13-minutes-of-for-honor-marching-fire-tiandi-breach-gameplay-e3-2018</guid>
      <media:content bitrate="501000.0" duration="800" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/07/02/640/6542fde3119b2b83b448dd0b54d7aecf-500000-1530558041.mp4">
        <media:keywords>tiandi, for honor, breach, e3, e3 2018, Gameplay, E3, Xbox One, For Honor, games, Action, Ubisoft, Ubisoft Montreal, PC, PS4, For Honor: Marching Fire, DLC / Expansion</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/7/2/tiandiwallpaper-1530557624914.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2018-07-02T19:00:00Z</dc:date>
    </item>
    <item>
      <title>Halo Infinite: Location, Armor, and AI Theories</title>
      <link>http://www.ign.com/videos/2018/06/29/halo-infinite-location-armor-and-ai-theories-from-the-trailer</link>
      <description>There are a lot of hints in the Halo Infinite trailer about where it takes place, how the armor changed in unexpected ways and which non-Cortana AI he will be using.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/06/29/640/f476242aad4a90855f5c0c6ae58387db-500000-1530308000.mp4" type="video/mp4" />
      <pubDate>Fri, 29 Jun 2018 22:00:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/06/29/halo-infinite-location-armor-and-ai-theories-from-the-trailer</guid>
      <media:content bitrate="501000.0" duration="325" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/06/29/640/f476242aad4a90855f5c0c6ae58387db-500000-1530308000.mp4">
        <media:keywords>e3, e3 2018, halo, halo infinite, armor, Mjolnir, ai, cortana, Feature, Xbox One, Halo Infinite, games, Shooter, Microsoft, 343 Industries</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/6/29/haloinfiniteanalysisyt-1530307794050.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2018-06-29T22:00:00Z</dc:date>
    </item>
    <item>
      <title>15 Huge Games Missing From E3 2018</title>
      <link>http://www.ign.com/videos/2018/06/29/15-huge-games-missing-from-e3-2018</link>
      <description>From the Final Fantasy VII Remake to Red Dead Redemption 2, here are the games we were bummed to see not make it to E3.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/06/29/640/968cc16672a5bbc8f9913c6dc6b189b6-347000-1530305161.mp4" type="video/mp4" />
      <pubDate>Fri, 29 Jun 2018 20:45:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/06/29/15-huge-games-missing-from-e3-2018</guid>
      <media:content bitrate="343000.0" duration="267" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/06/29/640/968cc16672a5bbc8f9913c6dc6b189b6-347000-1530305161.mp4">
        <media:keywords>none, Lists, E3, PS4, Final Fantasy VII Remake, games, RPG, Final Fantasy, Square Enix, CyberConnect2, Red Dead Redemption 2, Action, Adventure, Rockstar  Games, Xbox One, PC, Marvel's The Avengers, Crystal Dynamics, Square Enix Montreal, Switch, Shenmue III, Deep Silver, Ys Net, iPhone, Artifact, Card, Battle, Card Collecting Game (CCG), Valve, Android, Yoshi for Nintendo Switch, Platformer, Mario, Nintendo, Metroid Prime 4, Shooter, Pokemon for Nintendo Switch, Pokemon, Game Freak, Bayonetta 3, PlatinumGames, Deep Down, Capcom, Dead Island 2, Sumo Digital, Wild, Wild Sheep Studio, Witchfire, The Astronauts, Mac, The Last Night</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/6/29/36190303missingfrome31280-1530305153650.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2018-06-29T20:45:00Z</dc:date>
    </item>
    <item>
      <title>New Features Coming to World of Warcraft: Battle for Azeroth - E3 2018</title>
      <link>http://www.ign.com/videos/2018/06/26/new-features-coming-to-world-of-warcraft-battle-for-azeroth-e3-2018</link>
      <description>Blizzard's John Haight joins us to discuss the all-new features coming to World of Warcraft's Battle for Azeroth expansion!</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/06/26/640/627541c6d38ed7d7d2e25753d746c65e-500000-1530055884.mp4" type="video/mp4" />
      <pubDate>Tue, 26 Jun 2018 23:45:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/06/26/new-features-coming-to-world-of-warcraft-battle-for-azeroth-e3-2018</guid>
      <media:content bitrate="498000.0" duration="802" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/06/26/640/627541c6d38ed7d7d2e25753d746c65e-500000-1530055884.mp4">
        <media:keywords>Interview, E3, PC, World of Warcraft, games, RPG, Vivendi Games, Blizzard Entertainment, Mac</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/6/26/bfa-1530055841989.jpg" width="640" height="360" />
      </media:content>
      <dc:date>2018-06-26T23:45:00Z</dc:date>
    </item>
    <item>
      <title>Every 3rd-party Nintendo Switch Game Announced at E3 - Game Scoop! 488 Highlight</title>
      <link>http://www.ign.com/videos/2018/06/23/every-3rd-party-nintendo-switch-game-announced-at-e3-game-scoop-488-highlight</link>
      <description>23 new games are on the way to Nintendo's latest.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/06/22/640/5f2117475e48788fea51b5bbe1af2470-347000-1529701521.mp4" type="video/mp4" />
      <pubDate>Sat, 23 Jun 2018 15:57:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/06/23/every-3rd-party-nintendo-switch-game-announced-at-e3-game-scoop-488-highlight</guid>
      <media:content bitrate="345000.0" duration="253" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/06/22/640/5f2117475e48788fea51b5bbe1af2470-347000-1529701521.mp4">
        <media:keywords>none, Talkshow, E3</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/6/22/tsr-15276422716831280w-1529701504623.png" width="640" height="360" />
      </media:content>
      <dc:date>2018-06-23T15:57:00Z</dc:date>
    </item>
    <item>
      <title>Cyberpunk 2077: All the Gameplay Details You Might Have Missed in the E3 Trailer</title>
      <link>http://www.ign.com/videos/2018/06/23/cyberpunk-2077-all-the-gameplay-details-you-might-have-missed-in-the-e3-trailer</link>
      <description>We break down everything in the trailer that we learned from a demo behind-closed-doors at E3.</description>
      <enclosure url="http://assets14.ign.com/videos/zencoder/2018/06/22/640/21fa81de9b521b8f8a179388417db499-347000-1529712690.mp4" type="video/mp4" />
      <pubDate>Sat, 23 Jun 2018 00:25:00 GMT</pubDate>
      <guid isPermaLink="false">http://www.ign.com/videos/2018/06/23/cyberpunk-2077-all-the-gameplay-details-you-might-have-missed-in-the-e3-trailer</guid>
      <media:content bitrate="341000.0" duration="709" height="360" type="video/mp4" width="640" url="http://assets14.ign.com/videos/zencoder/2018/06/22/640/21fa81de9b521b8f8a179388417db499-347000-1529712690.mp4">
        <media:keywords>gameplay, trailer, detail, rewind, Feature, E3, PC, Cyberpunk 2077, games, RPG, CD Projekt, PS4, Xbox One</media:keywords>
        <media:thumbnail url="https://assets1.ignimgs.com/thumbs/userUploaded/2018/6/22/36190303screenshot-2-1529712675874.png" width="640" height="360" />
      </media:content>
      <dc:date>2018-06-23T00:25:00Z</dc:date>
    </item>
  </channel>
</rss>
`;

      var parseString = require('xml2js').parseString;
      parseString(data, (err, res) => {
        let num = 0;
        for(let article of res.rss.channel[0].item) {
          if(num >= 12) continue;

          let post = {
            important: num == 2 || num == 4,
            url: article.link[0],
            thumbnail: article['media:content'][0]['media:thumbnail'][0]['$'].url,
            cover: {
              type: article['media:content'][0]['$'].type,
              src: article['media:content'][0]['$'].url
            },
            title: article.title[0],
            preview: article.description[0]
          };

          this.news.push(post);

          num += post.important ? 2 : 1;

          console.log(post);
        }
      });
    }
  }
</script>