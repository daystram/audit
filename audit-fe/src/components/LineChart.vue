<script lang="ts">
import { Component, Mixins } from "vue-property-decorator";
import { Line, mixins } from "vue-chartjs";

@Component({
  extends: Line,
  mixins: [mixins.reactiveProp],
})
export default class LineChart extends Mixins(mixins.reactiveProp, Line) {
  mounted(): void {
    this.renderChart(this.chartData, {
      spanGaps: true,
      elements: {
        line: {
          borderColor: "#00ADFF",
          backgroundColor: "#00ADFF10",
        },
        point: {
          radius: 0,
          hitRadius: 8,
          hoverRadius: 0,
        },
      },
      scales: {
        xAxes: [
          {
            type: "time",
            time: {
              unit: "hour",
              stepSize: 0.25,
              displayFormats: {
                hour: "h:mm A",
              },
            },
            ticks: {
              major: {
                enabled: true,
              },
              fontColor: "#FFFFFFA0",
            },
            gridLines: {
              zeroLineColor: "#FFFFFF08",
              color: "#FFFFFF08",
            },
          },
        ],
        yAxes: [
          {
            ticks: {
              beginAtZero: true,
              callback: (value) => {
                return `${value} ms `;
              },
              fontColor: "#FFFFFFA0",
            },
            scaleLabel: {
              display: true,
              labelString: "Response Time",
              fontColor: "#FFFFFF80",
            },
            gridLines: {
              zeroLineColor: "#FFFFFF08",
              color: "#FFFFFF08",
            },
          },
        ],
      },
      legend: {
        display: false,
      },
      tooltips: {
        callbacks: {
          label: (tooltipItem) => {
            return (tooltipItem.yLabel as number) < 1e-3
              ? `No response`
              : `${tooltipItem.yLabel} ms `;
          },
        },
      },
      responsive: true,
      maintainAspectRatio: false,
    });
  }
}
</script>
