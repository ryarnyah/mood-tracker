<template>
    <div>
        <div v-if="title != null && content != null">
            <div>
                <h1>{{ title }}</h1>
                <h2>{{ content }}</h2>
            </div>
            <div v-if="!entrySaved">
                <div class="md-layout md-alignment-top-center">
                    <md-card class="md-layout-item md-size-50 md-small-size-100">
                        <md-card-content>
                            <div class="md-layout md-gutter">
                                <div class="md-layout-item md-small-size-100">
                                    <md-field>
                                        <label>Comment</label>
                                        <md-textarea md-autogrow name="comment" id="comment" v-model="comment" maxlength="128" />
                                    </md-field>
                                </div>
                            </div>
                        </md-card-content>
                    </md-card>
                </div>
                <div class="md-layout">
                    <div class="md-layout-item" v-on:click="saveEntry(1)"><md-button class="md-raised md-accent">Bad</md-button></div>
                    <div class="md-layout-item" v-on:click="saveEntry(2)"><md-button class="md-raised">Normal</md-button></div>
                    <div class="md-layout-item" v-on:click="saveEntry(3)"><md-button class="md-raised md-primary">Perfect</md-button></div>
                </div>
            </div>
            <div v-if="entrySaved">
                Thank you for your time.
            </div>
            <md-snackbar :md-active.sync="showError">{{ errorMessage }}</md-snackbar>
        </div>
        <div v-if="title == null && content == null">
            You have already answered this mood.
        </div>
    </div>
</template>

<script>
import { grpc } from '@improbable-eng/grpc-web';
import { Mood } from '../proto/mood_pb_service';
import { AddEntryRequest, GetMoodFromEntryRequest, Entry } from '../proto/mood_pb';

export default {
    name: 'AddEntry',
    props: [
        'moodId',
        'entryAccessCode'
    ],
    data: () => ({
      title: null,
      content: null,
      comment: null,
      sending: false,
      entrySaved: false,
      showError: false,
      errorMessage: null
    }),
    created () {
        this.getMood();
    },
    methods: {
      getValidationClass (fieldName) {
        const field = this.$v.form[fieldName]

        if (field) {
          return {
            'md-invalid': field.$invalid && field.$dirty
          }
        }
      },
      getMood () {
        let request = new GetMoodFromEntryRequest();
        request.setMoodId(this.moodId);
        request.setEntryAccessCode(this.entryAccessCode);

        let v = this;
        grpc.unary(Mood.GetMoodFromEntry, {
            request: request,
            host: '/grpc',
            onEnd: function(res) {
                const { status, statusMessage, message } = res;
                if (status === grpc.Code.OK && message) {
                    v.title = message.getTitle();
                    v.content = message.getContent();
                } else if (status !== grpc.Code.OK) {
                    v.showError = true;
                    v.errorMessage = statusMessage;
                }
            }
        });
      },
      saveEntry (record) {
          this.sending = true;
          let request = new AddEntryRequest();
          let entry = new Entry();
          entry.setRecord(record);
          entry.setComment(this.comment);

          request.setMoodId(this.moodId);
          request.setEntryAccessCode(this.entryAccessCode);
          request.setEntry(entry);

          let v = this;
          grpc.unary(Mood.AddEntry, {
              request: request,
              host: '/grpc',
              onEnd: function(res) {
                  const { status, statusMessage, message } = res;
                  if (status === grpc.Code.OK && message) {
                      v.entrySaved = true;
                  } else if (status !== grpc.Code.OK) {
                      v.showError = true;
                      v.errorMessage = statusMessage;
                  }

                  v.sending = false;
              }
          });
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
