<template>
  <div class="page-block">
    <h2 class="page-block__header">{{formatMonth(statistic.date)}}</h2>

    <div class="table__wrapper">
    <table class="table__table" cellspacing="0">
      <!-- TABLE HEAD -->
      <tr class="table__tr">
        <th class="table__th" rowspan="2">Оператор</th>
        <th class="table__th" rowspan="2">Сумма Ч.</th>
        <th class="table__th" rowspan="2">Сумма З.</th>
        <th class="table__th" colspan="4">Средние данные за месяц</th>
        <th class="table__th" colspan="5" 
          v-for="day in getDays()"
          v-bind:key="day">
          {{formatDay(day)}}
        </th>
      </tr>
        
      <tr class="table__tr">
        <th class="table__th">Начало</th>
        <th class="table__th">Конец</th>
        <th class="table__th">Перерыв</th>
        <th class="table__th">Работа</th>
        <template 
          v-for="day in getDays()"
          v-bind:key="day">
          <th class="table__th">Начало</th>
          <th class="table__th">Конец</th>
          <th class="table__th">Раб. Ч.</th>
          <th class="table__th">Пер.</th>
          <th class="table__th">Ч.З.</th>
        </template>
      </tr>

      <!-- TABLE DATA -->
      <tr class="table__tr"
        v-for="operatorStatistic in statistic.statistic"
        v-bind:key="operatorStatistic.operator.operator.id">
        <td class="table__td">{{operatorStatistic.operator.operator.worker.login}}</td>
        <td class="table__td">{{formatDuration(getOperatorShiftsTime(operatorStatistic.operator_shifts))}}</td>
        <td class="table__td">{{getOperatorShiftsDelays(operatorStatistic.operator_shifts)}}</td>
        <td class="table__td">{{formatTime(getMiddleOperatorShiftsStart(operatorStatistic.operator_shifts))}}</td>
        <td class="table__td">{{formatTime(getMiddleOperatorShiftsEnd(operatorStatistic.operator_shifts))}}</td>
        <td class="table__td">{{formatDuration(getMiddleOperatorShiftsTimeoutsTime(operatorStatistic.operator_shifts))}}</td>
        <td class="table__td">{{formatDuration(getMiddleOperatorShiftsTime(operatorStatistic.operator_shifts))}}</td>
        <template 
          v-for="day in getDays()"
          v-bind:key="day">
          <td class="table__td">{{formatTime(getOperatorShiftStartAtDay(operatorStatistic.operator_shifts, day))}}</td>
          <td class="table__td">{{formatTime(getOperatorShiftEndAtDay(operatorStatistic.operator_shifts, day))}}</td>
          <td class="table__td">{{formatDuration(getOperatorShiftTimeAtDay(operatorStatistic.operator_shifts, day))}}</td>
          <td class="table__td">{{getOperatorShiftTimeoutsCountAtDay(operatorStatistic.operator_shifts, day)}}</td>
          <td class="table__td">{{getOperatorShiftDelaysAtDay(operatorStatistic.operator_shifts, day)}}</td>
        </template>
      </tr>
    </table>
    </div>
  </div>
</template>


<script>
  import moment from 'moment';

  export default {
    name: "statistic-component",
    props: ['statistic'],
    methods: {
      formatMonth(date) {
        console.log(moment(date))

        const dateString = moment(date).format("MMMM YYYY")
        return dateString.charAt(0).toUpperCase() + dateString.slice(1)
      },
      formatDay(date) {
        const dateString = moment(date).format("dddd DD.MM.YYYY")
        return dateString.charAt(0).toUpperCase() + dateString.slice(1)
      },
      formatTime(date) {
        if (typeof date === "string") date = new Date(date)
        if (date <= 0) return "-"
        return moment(date).format("HH:mm")
      },
      formatDuration(seconds) {
        if (seconds === 0) return "-"

        const hours = Math.floor(seconds / 60 / 60)
        const minutes = Math.floor(seconds / 60 % 60)
        const secondsRem = Math.floor(seconds % 60)
        if (hours > 0) return `${hours}ч. ${minutes}м.`
        return `${minutes}м. ${secondsRem}с.`
      },
      getDays() {
        const statisticDate = new Date(this.statistic.date)
        const monthDate = new Date(statisticDate.getFullYear(), statisticDate.getMonth() + 1, 0);

        return Array.from({length: monthDate.getDate()}, (_, k) => new Date(
          monthDate.getFullYear(),
          monthDate.getMonth(),
          k+1,
          12
        ))
      },
      getOperatorShiftsTime(shifts) {
        if (shifts.length === 0) return 0 

        return shifts.reduce(
          (shiftsAcc, shift) => {
            if (new Date(shift.end_date).getTime() < 0) return shiftsAcc

            const timeouts = shift.timeouts.reduce(
              (timeoutsAcc, timeout) => {
                return timeoutsAcc + ((new Date(timeout.end_date) - new Date(timeout.start_date)) / 1000)
              }, 0
            )

            console.log(timeouts)

            return shiftsAcc + ((new Date(shift.end_date) - new Date(shift.start_date)) / 1000) - timeouts
          }, 0
        )
      },
      getOperatorShiftsDelays(shifts) {
        if (shifts.length === 0) return 0 

        return shifts.reduce((delaysAcc, shift) => {
          return delaysAcc + shift.delays 
        }, 0)
      },
      getOperatorShiftStartAtDay(shifts, day) {
        for (const shift of shifts) {
          if (new Date(shift.start_date).getDate() == day.getDate()) {
            return shift.start_date
          }
        }
        return 0
      },
      getOperatorShiftEndAtDay(shifts, day) {
        for (const shift of shifts) {
          if (new Date(shift.start_date).getDate() == day.getDate()) {
            return shift.end_date
          }
        }
        return 0
      },
      getOperatorShiftDelaysAtDay(shifts, day) {
        for (const shift of shifts) {
          if (new Date(shift.start_date).getDate() == day.getDate()) {
            return shift.delays
          }
        }
        return 0
      },
      getOperatorShiftTimeAtDay(shifts, day) {
        const shift = shifts.find((shift) => new Date(shift.start_date).getDate() === day.getDate())
        if (!shift) return 0
        if (new Date(shift.end_date).getTime() < 0) return 0

        const timeout = this.getOperatorShiftTimeoutsTimeAtDay(shifts, day)
        return ((new Date(shift.end_date) - new Date(shift.start_date)) / 1000) - timeout
      },
      getOperatorShiftTimeoutsTimeAtDay(shifts, day) {
        if (shifts.length === 0) return 0 
        const shift = shifts.find((shift) => new Date(shift.start_date).getDate() === day.getDate())
        if (!shift || shift.timeouts.length === 0) return 0
        if (new Date(shift.end_date).getTime() < 0) return 0
        
        return shift.timeouts.reduce((a, v) => a + (new Date(v.end_date) - new Date(v.start_date)), 0) / 1000
      },
      getOperatorShiftTimeoutsCountAtDay(shifts, day) {
        if (shifts.length === 0) return 0 
        const shift = shifts.find((shift) => new Date(shift.start_date).getDate() === day.getDate())
        if (!shift) return 0

        return shift.timeouts.length
      },
      getMiddleOperatorShiftsStart(shifts) {
        if (shifts.length === 0) return 0

        const minutes = Math.round(shifts.reduce((a, v) => {
          return a + new Date(v.start_date).getHours() * 60 + new Date(v.start_date).getMinutes()
        }, 0) / shifts.length)
        
        return new Date(new Date().getFullYear(), new Date().getMonth(), 1, Math.floor(minutes / 60), minutes % 60, 0, 0)
      },
      getMiddleOperatorShiftsEnd(shifts) {
        if (shifts.length === 0) return 0

        const minutes = Math.round(shifts.reduce((a, v) => {
          return a + new Date(v.end_date).getHours() * 60 + new Date(v.end_date).getMinutes()
        }, 0) / shifts.length)
        
        return new Date(new Date().getFullYear(), new Date().getMonth(), 1, Math.floor(minutes / 60), minutes % 60, 0, 0)
      },
      getMiddleOperatorShiftsTimeoutsTime(shifts) {
        if (shifts.length === 0) return 0

        return shifts.reduce(
          (shiftsAcc, shift) => {
            if (new Date(shift.end_date).getTime() < 0) return shiftsAcc
            return shiftsAcc + shift.timeouts.reduce((timeoutsAcc, timeout) => {
              return timeoutsAcc + ((new Date(timeout.end_date) - new Date(timeout.start_date)) / 1000)
            }, 0) / shift.timeouts.length
          }, 0
        ) / shifts.length
      },
      getMiddleOperatorShiftsTime(shifts) {
        if (shifts.length === 0) return 0

        return shifts.reduce(
          (shiftsAcc, shift) => {
            if (new Date(shift.end_date).getTime() < 0) return shiftsAcc

            const timeouts = shift.timeouts.reduce(
              (timeoutsAcc, timeout) => {
                return timeoutsAcc + ((new Date(timeout.end_date) - new Date(timeout.start_date)) / 1000)
              }, 0
            )

            return shiftsAcc + ((new Date(shift.end_date) - new Date(shift.start_date)) / 1000) - timeouts
          }, 0
        ) / shifts.length
      },
    },
  }
</script>


<style scoped>
  .table__wrapper {
    display: block;
    overflow-x: auto;
    max-width: 100%;
  }

  .table__table {
    display: block;
    max-width: 0;
    white-space: nowrap;
    padding-bottom: 20px;
  }

  .table__th,
  .table__td {
    text-align: center;
    padding: 4px 8px;
    vertical-align: top;
    border-top: 2px solid var(--lt__main);
    border-left: 2px solid var(--lt__main); 
  }

  .table__th:last-child,
  .table__td:last-child {
    border-right: 2px solid var(--lt__main); 
  }

  .table__tr:last-child > * {
    border-bottom: 2px solid var(--lt__main); 
  } 

  .table__login {
    text-align: left;
  }

  .table__hours {
    text-align: center;
  }


  /* DARK THEME */

  .dark__theme .table__th,
  .dark__theme .table__td {
    border-color: var(--dt__header);
  }
</style>