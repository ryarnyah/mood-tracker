<template>
  <div>
      <form novalidate class="md-layout md-alignment-top-center" @submit.prevent="validateForm">
          <md-card class="md-layout-item md-size-50 md-small-size-100">
              <md-card-header>
                  <div class="md-title">Create Mood</div>
              </md-card-header>
              <md-card-content>
                  <div class="md-layout md-gutter">
                      <div class="md-layout-item md-small-size-100">
                          <md-field :class="getValidationClass('title')">
                              <label for="title">Title</label>
                              <md-input name="title" id="title" v-model="form.title" :disabled="sending" maxlength="128" />
                              <span class="md-error" v-if="!$v.form.title.required">The title is required</span>
                              <span class="md-error" v-else-if="!$v.form.title.minlength">Invalid title</span>
                              <span class="md-error" v-else-if="!$v.form.title.maxlength">Invalid title</span>
                          </md-field>
                      </div>
                  </div>
                  <div>
                      <div class="md-layout-item md-small-size-100">
                          <md-field :class="getValidationClass('content')">
                              <label for="content">Content</label>
                              <md-textarea md-autogrow name="content" id="content" v-model="form.content" :disabled="sending" maxlength="512" />
                              <span class="md-error" v-if="!$v.form.title.maxlength">Invalid content</span>
                          </md-field>
                      </div>
                  </div>
                  <div>
                      <div class="md-layout-item md-small-size-100">
                          <md-field :class="getValidationClass('numberOfEntries')">
                              <label for="numberOfEntries">Number of answer needed</label>
                              <md-input name="numberOfEntries" id="numberOfEntries" v-model="form.numberOfEntries" :disabled="sending" />
                              <span class="md-error" v-if="!$v.form.numberOfEntries.required">The Number of entries is required</span>
                              <span class="md-error" v-else-if="!$v.form.numberOfEntries.beetween">Invalid content</span>
                          </md-field>
                      </div>
                  </div>
              </md-card-content>
              <md-progress-bar md-mode="indeterminate" v-if="sending" />
              <md-card-actions>
                  <md-button type="submit" class="md-primary" :disabled="sending">Create mood</md-button>
              </md-card-actions>
              <md-snackbar :md-active.sync="moodSaved">The mood was created with success!</md-snackbar>
              <md-snackbar :md-active.sync="showError">{{ errorMessage }}</md-snackbar>
          </md-card>
      </form>
      <div class="md-layout md-alignment-top-center" v-if="moodHasData">
          <md-card class="md-layout-item md-size-50 md-small-size-100">
              <md-card-content>
                  <md-table>
                      <md-table-row>
                          <md-table-head md-numeric>Mood Url (must save it somewhere)</md-table-head>
                      </md-table-row>
                      <md-table-row>
                          <md-table-cell md-numeric><a v-bind:href="moodUrl">{{ moodUrl }}</a></md-table-cell>
                      </md-table-row>
                  </md-table>
                  <md-table>
                      <md-table-row>
                          <md-table-head md-numeric>Entries URLs</md-table-head>
                      </md-table-row>
                      <md-table-row v-for="entry in entriesAccessCodesList" :key="entry">
                          <md-table-cell md-numeric><a v-bind:href="entryUrl(entry)">{{ entryUrl(entry) }}</a></md-table-cell>
                      </md-table-row>
                  </md-table>
              </md-card-content>
          </md-card>
      </div>
  </div>
</template>

<script>
import { grpc } from '@improbable-eng/grpc-web';
import { Mood } from '../proto/mood_pb_service';
import { CreateMoodRequest } from '../proto/mood_pb';

import { validationMixin } from 'vuelidate';
import {
    required,
    minLength,
    maxLength,
    between
  } from 'vuelidate/lib/validators';

export default {
    name: 'CreateMood',
    mixins: [validationMixin],
    computed: {
        moodUrl: function() {
            return window.location.protocol + '//' + window.location.host + '/mood/' + this.moodId + '/' + this.moodAccessCode;
        }
    },
    data: () => ({
      form: {
          title: null,
          content: null,
          numberOfEntries: null
      },
      sending: false,
      moodSaved: false,
      moodHasData: false,
      entriesAccessCodesList: null,
      moodId: null,
      moodAccessCode: null,
      showError: false,
      errorMessage: null
    }),
    validations: {
      form: {
        title: {
            required,
            minLength: minLength(1),
            maxLength: maxLength(128)
        },
        content: {
            minLength: minLength(1),
            maxLength: maxLength(512)
        },
        numberOfEntries: {
            required,
            between: between(1, 20)
        }
      }
    },
    methods: {
      entryUrl(entryAccessCode) {
          return window.location.protocol + '//' + window.location.host + '/entry/' + this.moodId + '/' + entryAccessCode;
      },
      getValidationClass (fieldName) {
        const field = this.$v.form[fieldName]

        if (field) {
          return {
            'md-invalid': field.$invalid && field.$dirty
          }
        }
      },
      saveMood () {
          this.sending = true;
          let request = new CreateMoodRequest();
          request.setTitle(this.form.title);
          request.setContent(this.form.content);
          request.setNumberOfRecordsNeeded(this.form.numberOfEntries);

          let v = this;
          grpc.unary(Mood.CreateMood, {
              request: request,
              host: '/grpc',
              onEnd: function(res) {
                  const { status, statusMessage, message } = res;
                  if (status === grpc.Code.OK && message) {
                      v.entriesAccessCodesList = message.getEntriesAccessCodesList();
                      v.moodId = message.getMoodId();
                      v.moodAccessCode = message.getMoodAccessCode();
                      v.moodSaved = true;
                      v.moodHasData = true;
                  } else if (status !== grpc.Code.OK) {
                      v.showError = true;
                      v.errorMessage = statusMessage;
                  }

                  v.sending = false;
              }
          });
      },
      validateForm () {
        this.$v.$touch()

        if (!this.$v.$invalid) {
          this.saveMood()
        }
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
