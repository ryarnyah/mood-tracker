<script>
import { Line } from 'vue-chartjs'

export default {
    extends: Line,
    props: {
        chartData: {},
        cssClasses: {
            type: String,
            default: ''
        },
        width: {
            default: null,
            type: Number
        },
        height: {
            default: null,
            type: Number
        },
    },
    data: () => ({
        options: {
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        unit: 'day'
                    }
                }],
                yAxes: [{
                    ticks: {
                        beginAtZero: true,
                        min: 0,
                        max: 3,
                        stepSize: 1
                    }
                }]
            }
        }
    }),
    mounted () {
        this.renderChart(this.chartData, this.options);
    }
}
</script>
