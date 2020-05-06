<template>
    <div>
        <div v-if="title != null && content != null">
            <div>
                <h1>{{ title }}</h1>
                <h2>{{ content }}</h2>
            </div>
            <div class="md-layout md-alignment-top-center">
                <MoodStatChart v-bind:chartData="chartData" cssClasses="md-layout-item md-size-50 md-small-size-100" />
            </div>
            <div class="md-layout md-alignment-top-center">
                <md-card class="md-layout-item md-size-50 md-small-size-100">
                    <md-card-content>
                        <md-table>
                            <md-table-row>
                                <md-table-head>Record</md-table-head>
                                <md-table>
                                    <md-table-row>
                                        <md-table-head>Date</md-table-head>
                                        <md-table-head>Count</md-table-head>
                                    </md-table-row>
                                </md-table>
                            </md-table-row>
                            <md-table-row v-for="stat in stats" :key="stat.getRecord()">
                                <md-table-cell>{{ getLabelForRecord(stat.getRecord()) }}</md-table-cell>
                                <md-table>
                                    <md-table-row v-for="recordStat in stat.getRecordStatsList()" :key="recordStat.getRecordEntry().getSeconds()">
                                        <md-table-cell>{{ recordStat.getRecordEntry().toDate().toLocaleDateString() }}</md-table-cell>
                                        <md-table-cell>{{ recordStat.getCount() }}</md-table-cell>
                                    </md-table-row>
                                </md-table>
                            </md-table-row>
                        </md-table>
                        <md-table>
                            <md-table-row>
                                <md-table-head>Record</md-table-head>
                                <md-table-head>Date</md-table-head>
                                <md-table-head>Comment</md-table-head>
                            </md-table-row>
                            <md-table-row v-for="entry in entries" :key="entry.record">
                                <md-table-cell>{{ getLabelForRecord(entry.getRecord()) }}</md-table-cell>
                                <md-table-cell>{{ entry.getRecordEntry().toDate().toLocaleDateString() }}</md-table-cell>
                                <md-table-cell>{{ entry.getComment() }}</md-table-cell>
                            </md-table-row>
                        </md-table>
                    </md-card-content>
                </md-card>
            </div>
        </div>
        <md-snackbar :md-active.sync="showError">{{ errorMessage }}</md-snackbar>
    </div>
</template>

<script>
import { grpc } from '@improbable-eng/grpc-web';
import { Mood } from '../proto/mood_pb_service';
import { GetMoodRequest } from '../proto/mood_pb';
import MoodStatChart from './MoodStatChart.vue';
import moment from 'moment';

export default {
    name: 'GetMood',
    props: [
        'moodId',
        'moodAccessCode'
    ],
    components: {
        MoodStatChart,
    },
    data: () => ({
      sending: false,
      title: null,
      content: null,
      entries: [],
      stats: null,
      showError: false,
      errorMessage: null,
      chartData: null
    }),
    created () {
        this.getMood();
    },
    methods: {
      dynamicColors() {
          var r = Math.floor(Math.random() * 255);
          var g = Math.floor(Math.random() * 255);
          var b = Math.floor(Math.random() * 255);
          return "rgb(" + r + "," + g + "," + b + ")";
      },
      getChartData() {
          var datasets = [];
          for (const stat of this.stats) {
              var label = this.getLabelForRecord(stat.getRecord());
              var data = [];
              for (const recordStat of stat.getRecordStatsList()) {
                  data.push({
                      t: new moment(recordStat.getRecordEntry().toDate()),
                      y: recordStat.getCount()
                  })
              }
              datasets.push({
                  backgroundColor: this.dynamicColors(),
                  label: label,
                  data: data,
              })
          }
          return {
              datasets: datasets
          };
      },
      getLabelForRecord(record) {
          switch (record) {
              case 0:
                  return 'Not answered';
              case 1:
                  return 'Bad';
              case 2:
                  return 'Normal';
              case 3:
                  return 'Perfect';
          }
      },
      getMood () {
        let request = new GetMoodRequest();
        request.setMoodId(this.moodId);
        request.setMoodAccessCode(this.moodAccessCode);

        let v = this;
        grpc.unary(Mood.GetMood, {
            request: request,
            host: '/grpc',
            onEnd: function(res) {
                const { status, statusMessage, message } = res;
                if (status === grpc.Code.OK && message) {
                    v.title = message.getTitle();
                    v.content = message.getContent();
                    v.entries = message.getEntriesList();
                    v.stats = message.getStatsList();
                    v.chartData = v.getChartData(v.stats);
                } else if (status !== grpc.Code.OK) {
                    v.showError = true;
                    v.errorMessage = statusMessage;
                }
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
