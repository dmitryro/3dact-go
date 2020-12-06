import axios from 'axios';

axios.defaults.baseURL = 'https://api.github.com';

var front_articles = Array();

var whyus = Array();

var features = Array();
features.push({"title":"Feature1",
               "description": "Description1"});
features.push({"title":"Feature2",
               "description": "Description2"});
features.push({"title":"Feature3",
                "description": "Description3"});

var about = Array();
about.push({"title":"About Us",
            "story": "3D Act is an AI and automation platform that is transforming the way businesses deliver customer support. Use 3D Act to automate answering support inquiries, give reps everything they need at their fingertips, and ensure your information is always accurate. 3D Act integrates with your existing systems and workflows to build machine learning models of common tasks, answer known questions, and make every rep more productive. Start using the most accurate AI-powered performance on the market to reduce resolution times, eliminate unnecessary escalations, and improve customer satisfaction. 3D Act was founded in 2015 and is based in New York City, NY."});

var howworks = Array();
howworks.push({"index":"1", "title":"Itegrate", "description":"Integrate your Knowledge Base, Dashboard, API and third party tools to read conversations from."});
howworks.push({"index":"2","title":"Load", "description":"Sanitize your content - remove sensitive private data using AI/NLP and store in Knowledge Base."});
howworks.push({"index":"3","title":"Monitor", "description":"Use dashboard to define rules to follow and events to respond, actions to run and shortcuts to use."});
howworks.push({"index":"4","title":"Use", "description":"Get envolned when it's really needed, not all the time."});
howworks.push({"index":"5","title":"Feedback", "description":"Provide feedback when it's needed automatically."});

var choose = Array();


var faqs = Array();
faqs.push({"index":"1", "question":"How does it work", "answer": "It just works."});
faqs.push({"index":"2", "question":"How do I install", "answer": "You install."});
faqs.push({"index":"3", "question":"Where do I learn", "answer": "You learn."});
faqs.push({"index":"4", "question":"Where do I start", "answer": "You start."});
faqs.push({"index":"5", "question":"How do I solve it", "answer": "You solve it."});

var usecases = Array();
usecases.push({"index":"1", "title":"Utilize internal knowledge", 
               "description": "Build your internal knowledge base. Get the high quality responses that are sanitized, individualized and designed for the  best experience. Your questions may be non-trivial, ambiguous or open-ended, but they will get the best answer possible."});
usecases.push({"index":"2", "title":"Push content to external venues", 
               "description": "Have your answers published in a number of venues - Slack/Teams/Jira/Salesforce/Confluence/FAQ sections/online documentation/public sites. Turn some of your discussions into learning experience, pull requests, Jira tickets, online materials."});
usecases.push({"index":"3", "title":"Provide interactive experience.", 
               "description": "Integrate your favorite chat patform with knowledge base and Slack/Teams. Manage your customer support on a way that minimizes the need for live assistance. Motivate your customers to use as much self-service as they can. Syncronize with external resources."});


var team = Array();
team.push({"username":"Dmitry",
           "name": "Dmitry Roitman",
           "title": "CTO and Co-Founder",
           "avatar": "../assets/png/icon-avatar-default.png",
           "profile": "some",
           "bio": "Dmitry is an experienced developer and architect. He's building our platform."});
team.push({"username":"Jimin",
           "avatar": "../assets/png/icon-avatar-default.png",
           "profile": "some",
           "name": "Jimin Kim",
           "title": "CEO and Co-Founder",
           "bio": "Description2"});
team.push({"username":"dmitryro",
           "avatar": "../assets/png/icon-avatar-default.png",
           "profile": "some",
           "name": "Dmitry Roitman",
           "title": "COO and Co-Founder",
           "bio": "Description3"});

var team_members = Array();

team_members.push({"username":"dmitryro",
                   "profile": "some story",
                   "bio": "some cool story",
                   "avatar":"../assets/png/icon-avatar-default.png"});

team_members.push({"username":"jiminkim",
                   "profile": "some story",
                   "bio": "some cool story",
                   "avatar":"../assets/png/icon-avatar-default.png"});

team_members.push({"username":"thirdperson",
                   "profile": "some story",
                   "bio": "some cool story",
                   "avatar":"../assets/png/icon-avatar-default.png"});

var packages = Array();
packages.push({"title": "Free", 
               "description":"Free Package",
               "pricenote": "for teems of 1-10 users*",
               "packagenote": "*Teams of 11+ are $7 per user, per month",
               "period": ["month","year"],
               "items": ["Slack Bot", "2 Knowledge Integrations", "Analytics and Reporting", "Verification",  "Browser Extension", "Basic Support"],
               "price": ["Free", "Free"]});
packages.push({"title": "Startup",  
               "description":"Startup Package",
               "period": ["month", "year"],
               "pricenote": "for teems of 1-10 users*",
               "items": ["Slack Bot", "3 Knowledge Integrations", 
                                      "Analytics and Reporting", 
                                      "Verification", "Browser Extension", "Priority Support"],
               "packagenote": "*Teams of 11+ are $60 per user, per month",
               "price": ["40 USD","450 USD"]});
packages.push({"title": "Enterprise",  
               "description":"Enterprise Package",
               "period": ["month","year"],
               "pricenote": "for teems of 100-1000 users*",
               "items": ["Slack Bot", "5 Knowledge Integrations", "Analytics and Reporting", "Verification", "Advanced AI", "Browser Extension", "Priority Support"],
               "packagenote": "*Teams of 1000+ are $150 per user, per month",
               "price": ["100 USD", "1000 USD"]});

front_articles.push({"title": "Time is lost", "description": "Companies have multiple communication channels - Slack, Jira, Salesforce, Dropbox, Bitbucket, Github. Customers do initial research online but frequently do not find a satifying answer, so they resort to Slack, or Jira. Support teams, developers and consultants then spend time answering same queistons over and over again. There's a way to replace this repeated, tedious and counter-productive process with more efficient one.", "points":[]});
front_articles.push({"title": "Knowledge unutilized", "description": "Most content can't be just stored in a knowledge base for further reuse because it contains sensitive data - names, dates, addresses, emails, other sensitive private information. For the same reason this content can't be made publically available. This leads to missed opportunities and missed leads. To be properly reused, this content needs to be sanitized and stored in a template-like format in the knowledge base (as an ontology, a graph, a tree, a taxonomy).", "points":[]});
front_articles.push({"title": "Customer success management is inefficient",
                    "description":"When customers come to your online resources, such as FAQ sections, Q/A sections, documentation portals, or blogs, and are looking for an answer or an example, they frequently don't know where to find the suitable answer, if any exists for them at all. Sometimes there's none for them, and they have to either resort to live conversation - in Slack/Teams or online chat, or just give you a call. Many times they're just not motivated enough to go through this whole experience and give up the process altogether. To keep your customers on your side with loyalty and lasting relationship, you need tools that bring everthing in sync and allow them to get a quality response to their most  irregular inquiries.", "points":[]});
front_articles.push({"title": "Impact is not measured", "description": "There's currently very limited number of platforms efficiently reusing knowledge and not just storing it in knowledge base, but taking into consideration many facts and events happening in background. An answer provided to each specific customer may be different one in a matter of minutes, as things change quickly. What we do we carefully analyze facts, we bring our knowledge from our knowledge base and we adjust it to reflect the context and provide the most educated response possible. We also provide automation and extensive insigtful analytics to measure success.", "points":[]});

front_articles.push({"title": "Analytics unutilized", "description": "The observations you make for charts - Google Analytics, Kibana, CloudWatch, Graphana, Mixpanel - require you to make decisions based on limited visual evaluations and within limited time, but you still have to do it manually. The way it's done many facts are not taken into consideration and decision making is based on judgements and thus is prone to potential errors. You never get an educated recommendation based on real, noise-free data. We provide a convinient way to store data, sanitize it, turn facts into actions.","points":[]});

front_articles.push({"title": "Have your conversations streamed, sanitized and actionalized from:", "description":"Integrate with third party tools:","points":["Atlassian (Jira, Confluence, Bitbucket)", "ELK", "Slack", "SAP", "Dropbox", "Salesforce", "Github", "Trello", "Kubernetes"]});

front_articles.push({"title": "Have your knowledge base and Slack/Teams integrated and synchronized with:", "description":"Online chat tools:","points":["Drift", "Freshdesk", "HubSpot Marketing Hub", "Intercom", "LivehCat", "Olark", "Pure Chat", "SnapEngage", "Zendesk"]});


whyus.push({"index":"1",
            "title":"Improved customer experience",
            "list":["Customers find more answers online and can do self-service as much as needed.",
                    "Content is depersonalized and is reusable, then it gets personalized later as per need.",
                    "Targeted response in Slack or Teams becomes more focused and more human. More trust is built. ",
                    "You're able to get the help just in place if you wish."]});


whyus.push({"index":"2",
            "title":"Improved employee productivity",
            "list":["Busy engineers spend less time doing repeated things and addressing repeated issues.",
                    "Customer Success productivity increased - more time and better focus.",
                    "Friction eliminated, wellbeing improves and burnout rate falls."]});


whyus.push({"index":"3",
            "title":"Improved brand competitiveness",
            "list":["Happier customers mean more competitive brand and more revenue.",   
                    "More productive and less stressed employees mean better business and stronger growth.",
                    "Now your team has more time to innovate and less need to worry about being outcompeted."]});


whyus.push({"index":"4",
            "title":"Cutting edge technology advantage",
            "list": ["Being up to date with the science and the data means being ahead of the competitors.",
                     "Keeping it human-yet-advanced, automated-yet-friendly means being ahead with AI. ",
                     "The platform is there for your customers and you when you needed and to the full capacity possible."]});

choose.push({"index":"1",
             "title":"If you  want to evaluate:",
             "list": ["Choose our free version to start with.", "This will give you the taste of the technology", 
                      "You can always to uprade to a more advanced or a different product package.",
                      "If you know exactly what feature set you need, and need an advice, please contact us!",
                      "If you've been still undecided about what's best, let's chat and you'll help you make the right choice!"]});

choose.push({"index":"2",
             "title":"If you are an early stage startup:",
             "list": ["Our initial suggestion is to start using our Start Up package.", 
                      "If you have multiple teams, and they need different feature sets, let's connect and discuss.", 
                      "If you have some unique use cases which we have not mentioned, or feel like they are related but not coverd, let's chat!",
                      "If you've been still undecided about what's best, let's chat and we'll help you make the right choice!"]});

choose.push({"index":"3",
             "title":"You are funded startup (50+ employees):",
             "list": ["Our initial suggestion is to start using our Start Up package.", 
                      "If you have multiple teams, and they need different feature sets, let's connect and discuss.", 
                      "If you have some unique use cases which we have not mentioned, or feel like they are related but not coverd, let's chat!",
                      "If you've been still undecided about what's best, let's chat and we'll help you make the right choice!"]});

choose.push({"index":"4",
             "title":"You are an enterprise (150+ employees):",
             "list": ["Our suggestion is to buy our enterprise package.", 
                      "If you want to start small, you can also use our Start Up package, but we'll need to discuss some use cases.", 
                      "If your enterprise has multiple departments, teams and has unique need in connecting them, let's discuss our extra functionalities for enterprises.",
                      "If you've been still undecided about what's best, let's chat and we'll help you make the right choice!"]});

choose.push({"index":"5",
             "title":"You are an need of automation:",
             "list": ["Our suggestion is to buy our enterprise package.",
                      "If you want to start small, you can also use our Start Up package, but we'll need to discuss some use cases.",
                      "If your enterprise has multiple departments, teams and has unique need in connecting them, let's discuss our extra functionalities for enterprises.",
                      "If you've been still undecided about what's best, let's chat and we'll help you make the right choice!"]});




export default {
  getUsecases() {
      return usecases;
  },
  getChoose() {
      return choose;
  },
  getFaqs() {
      return faqs;
  },
  getWhyus() {
      return whyus;
  },
  getHowworks() {
      return howworks;
  },
  getFrontArticles() {
      return front_articles;
  },
  getPackages() {
      return packages;
  },
  getFeatures() {
      return features;
  },
  getAbout() {
      return about;
  },
  getTeam() {
      return team;
  },
  getRepos(repo) {
    return axios.get(
      `/repos/${repo}`,
      { headers: { Authorization: process.env.GITHUB_TOKEN } },
    );
  },

  getMembers() {
    return axios.get(
      '/orgs/french-ai/members',
      { headers: { Authorization: process.env.GITHUB_TOKEN } },
    );
  },
};
