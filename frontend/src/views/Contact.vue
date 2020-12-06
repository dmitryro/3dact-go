<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container has-text-left">
          <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-11 has-text-left">
                  <h1 class="title">Contact Us</h1>
              </div>
          </div>
             <div class="columns">
                <div class="column is-1">
                </div>
                <div class="column is-4">
                   <div class="container">
                       <figure class="image is-fullwidth">
                           <img src="../assets/svg/contact.svg">
                       </figure>
                   </div>
                   <div class="container mt-6">

                        <div class="columns">

                          <div class="column is-half">
                                 <header>
                                   <span class="title" style="font-size:1.2em;">
                                      NYC Office
                                   </span>
                                 </header>
                                 <div class="mt-2" style="font-size:0.95em;">
                                          <ul class="has-text-left bulletless">
                                              <li>Union Square 15</li>
                                              <li>New York, NY</li>
                                              <li><time datetime="2016-1-1">9:00 AM - 6:00 PM EST</time></li>
                                          </ul>
                                 </div>
                                 <footer class="card-footer">
                                    Phone: (212) 111-2222
                                 </footer> 
                          </div> <!-- end half -->
                          <div class="column is-half">
                                 <header>
                                   <span class="title"  style="font-size:1.2em;">
                                      SF Office
                                   </span>
                                 </header>
                                 <div class="mt-2" style="font-size:0.95em;">
                                          <ul class="has-text-left bulletless">
                                              <li>Some Street, #no</li>
                                              <li>San-Francisco, CA</li>
                                              <li><time datetime="2016-1-1">9:00 AM - 6:00 PM PST</time></li>
                                          </ul>
                                 </div>
                                 <footer class="card-footer">
                                    Phone: (875) 111-2222
                                 </footer>
                          </div> <!-- end half -->

                         </div> <!-- end columns -->                      
                   </div> <!-- end container -->
                </div> <!-- end is-4-->
                <div class="column is-1">
                </div>
                <div class="column is-4">




                <section id="new-item-form">
                   
                    <b-field is-full label="Full Name" :type="isFullNameDanger" :message="FullNameMessage">
                          <b-input v-model="fullname" placeholder="Full Name" v-hammer:tap="onFullTap"></b-input>                        
                    </b-field>


                    <b-field is-full label="Email" :type="isEmailDanger" :message="EmailMessage">
                          <b-input v-model="email" placeholder="Email" v-hammer:tap="onEmailTap"></b-input>
                    </b-field>

                    <b-field is-full label="Title" :type="isTitleDanger" :message="TitleMessage">
                          <b-input v-model="title" placeholder="Title" v-hammer:tap="onTitleTap">


                          </b-input>
                    </b-field>
                  
                    <label class="label">Subject</label> 
                    <b-field is-full :type="isSubjectDanger" :message="SubjectMessage"> 
                        <b-select is-full class="is-fullwidth" v-hammer:tap="onSubjectTap" placeholder="Select Subject ..." v-model="subject" style="min-width:100% !important;">
                          <option is-full class="is-fullwidth"  v-hammer:tap="onSubjectTap"
                              v-for="option in subjects"
                              :value="option"
                              :key="option">
                              {{ option }}
                          </option>
                        </b-select>
                    </b-field>

                    <b-field is-full label="Message" :type="isMessageDanger" :message="MessageMessage">
                        <b-input maxlength="200" v-hammer:tap="onMessageTap" placeholder="Message..." type="textarea" v-model="message"></b-input>
                    </b-field>

                   <div class="columns">
                        <div class="column is-one-quarter">
                            <div class="field is-grouped">
                                <div class="control">
                                       <b-button class="button is-medium is-primary"
                                        :disabled="isNewFullNameFormDisabled"
                                         @click="submitForm">Send</b-button>
                                </div>
                             </div>
                        </div>
                        <div class="column is-three-quarters">
                            <b-notification v-model="isSent" style="margin-top:-1.0em;" auto-close type="is-success" aria-close-label="Close notification">
                               We will be in touch soon!!!
                            </b-notification>
                        </div>
                    </div>
                </section>
             </div>
             <div class="column is-2">
             </div>


          </div>

</div> <!-- close column-->
</div> <!-- close columns -->

  </section>
</template>
<script>
import client from '../services/client'

export default {
  name: 'Contact',
  props: {
  },
  components: {
  },
  data() {
    return {
      isTitleLoaded: false,
      isEmailLoaded: false,
      isSubjectLoaded: false,
      isMessageLoaded: false,
      isFullLoaded: false,
      isSent: false,
      subject: null,
      message: null,
      fullname: "",
      title: "",
      email:"",
      subjects: ["Product                           ", 
                 "Partnership                       ", 
                 "Demo                              ", 
                 "Features                          ", 
                 "Pricing                           ", 
                 "Other                             "],
      infos: [],
    };
  },
  methods: {
    onFullTap() {
       this.isFullLoaded = true;
    },
    onTitleTap() {
       this.isTitleLoaded = true;
    },
    onSubjectTap() {
       this.isSubjectLoaded = true;
    },
    onMessageTap() {
       this.isMessageLoaded = true;
    },
    onEmailTap() {
       this.isEmailLoaded = true;
    },
    submitForm(event) {
            event.preventDefault(),
            //process...             

            //this.$buefy.notification.open('We will be in touch soon!')
            client.post('api',  {"full_name":this.fullname, "email": this.email, "message": this.message, "title":this.title, "subject":this.subject});
            this.isSent = true;
            this.isMessageLoaded = false;
            this.isTitleLoaded = false;
            this.isFullLoaded = false;
            this.isSubjectLoaded = false;
            this.isEmailLoaded = false;
            this.message = "";
            this.title = "";
            this.subject = "";
            this.email = "";
            this.fullname = "";
            this.isFullNameDanger = "is-success";
            this.isTitleDanger = "is-success";
            this.isEmailDanger = "is-success";
            this.isSubjectDanger = "is-success";
            this.isMessageDanger = "is-success";
            event.target.reset();
    },
    adjust() {
      var element = document.getElementsByTagName("select")[0];
      if (element) {
          var parent = element.parentElement;
          parent.classList.add("is-fullwidth");
      }
    },
    hasFullName() {
        if (this.fullname) {
            return true;
        } else {
            return false;
        } 
    },
    validEmail: function (email) {
      var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    }
  },
  onSubjectChange: function () {
      var element = document.getElementsByTagName("select")[0];
      var parent = element.parentElement;
      parent.classList.add("is-fullwidth");
  },
  created: function () {
  },
  mounted() {
      var element = document.getElementsByTagName("select")[0];
      var parent = element.parentElement;
      parent.classList.add("is-fullwidth");
  },
  computed: {
    isFullNameDanger() {
        if (!this.isFullLoaded) {
            return "is-success";
        }
        if (this.fullname !== "") {
            return "is-success";
        } else {
            return "is-danger";
        }
    },
    isTitleDanger() {
        if (!this.isTitleLoaded) {
            return "is-success";
        }
        if (this.title !== "") {
            return "is-success";
        } else {
            return "is-danger";
        }
    },
    isSubjectDanger() {
        if (!this.isSubjectLoaded) {
            return "is-success";
        } 
        if (this.subject) {
            this.adjust();
            return "is-success";
        } else {
            this.adjust();
            return "is-danger";
        }
    },
    isEmailDanger() {
        if (!this.isEmailLoaded) {
            return "is-success";
        }
        if (!this.validEmail(this.email)) {
            return "is-danger";
        } else if (this.email !== "") {
            return "is-success";
        } else {
            return "is-danger";
        }
    },
    isMessageDanger() {
        if (!this.isMessageLoaded) {
            return "is-success";
        }
        if (this.message !== "") {
            return "is-success";
        } else {
            return "is-danger";
        }
    },
    FullNameMessage() {
        if (!this.isFullLoaded) {
            return "";
        }
        if (this.fullname !== "") {
            return "";
        } else {
            return "Please enter an FullName";
        }
    },
    MessageMessage() {
        if (!this.isMessageLoaded) {
            return "";
        }
        if (this.message !== null) {
            return "";
        } else {
            return "Please enter a Message";
        }
    },
    TitleMessage() {
        if (!this.isTitleLoaded) {
            return "";
        }
        if (this.title !== "") {
            return "";
        } else {
            return "Please enter an Title";
        }
    },
    SubjectMessage() {
        if (!this.isMessageLoaded) {
            return "";
        }
        if (this.subjects.includes(this.subject)) {
            this.adjust();
            return "";
        } else {
            this.adjust();
            return "Please select a valid Subject!";
        }
    },
    EmailMessage() {
        if (!this.isEmailLoaded) {
            return "";
        }
        if (this.email !== "") {
            if (!this.validEmail(this.email)) {
                return "Please enter a valid Email!"
            } else {
                return "";
            }
        }
        else {
            return "Please enter an Email";
        }
    },
    formIsInDanger() {
      if (this.isFullNameDanger === "is-danger" 
       || this.isMessageDanger === "is-danger" 
       || this.isEmailDanger === "is-danger" 
       || this.isSubjectDanger === "is-danger"
       || this.isTitleDanger === "is-danger") {
        return true;
      } else {
        return false;
      }
    },
    isNewFullNameFormDisabled() {
      if (this.formIsInDanger) return true;
      else return false;
    }
  },
};
</script>

<style>
  footer {
    margin-top: 50px;
  }
</style>
