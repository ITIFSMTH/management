<template>
  <div class="app">
    <component :is="layout">
      <router-view />
    </component>

    <div class="modal__background" 
      :class="{'modal__background-active': this.$store.getters.isModalActive}"></div>
  </div>
</template>

<script>
  import LoginLayout from '@/layouts/LoginLayout';
  import PanelLayout from '@/layouts/PanelLayout'

  export default {
    name: 'App',
    components: {
      LoginLayout,
      PanelLayout
    },
    computed: {
      layout() {
        return this.$route.meta.layout || 'login-layout'
      }
    },
    mounted() {
      this.$store.commit('updateUser')
      
      if (this.$store.getters.isUserAuthenticated) {
        this.$store.dispatch('getTheme')
      }
    }
  }
</script>

<style>
  :root {
    --white: #fff;
    --grey: rgb(127, 127, 127);


    --lt__green: #4fbe87;
    --lt__red: #f3616d;
    --lt__yellow: #cca825;
    --lt__orange: #fd7e14;
    --lt__grey: #e9ecef;
    --lt__light-blue: rgb(13, 202, 240);

    --lt__primary: #dce7f1;

    --lt__main-lighter_1: #8f6ccf;
    --lt__main-lighter: #815aca;
    --lt__main: #6f42c1;
    --lt__main-darker: #5a2eac;
    --lt__main-darker_1: #40217c;

    --lt__dark-text: #25396f;
    --lt__light-text: #aaaeb8;
    

    --dt__green: #21b46a;
    --dt__red: #ec3746;
    --dt__yellow: #e4b60f;

    --dt__primary: #141B24;
    --dt__primary-lighter: #1E2A3D;
    --dt__primary-darker: #0e141a;

    --dt__main-lighter: #4b2694;
    --dt__main: #452485;
    --dt__main-darker: #3c1f77;

    --dt__header: #8e65df;
    --dt__text-lighter: #c3adf4;
    --dt__text: #815aca;
    --dt__text-darker: #5f34af;
    --dt__skyblue: #8090b4;
  }

  * {
    margin: 0;
    padding: 0;
    font-family: 'Montserrat', sans-serif;
  }

  a {
    color: var(--lt__main);
    text-decoration: none;
  }

  body {
    background-color: var(--lt__primary);
  }

  .container {
    max-width: 1400px;
    padding: 0 40px;
    width: 100%;
    margin: 0 auto;
    box-sizing: border-box;
  }

  .input-box {
    width: 100%;
    padding-bottom: 15px;
  }

  .input-box__input {
    width: 100%;
  }

  .input-box__label {
    font-size: 15px;
    padding-left: 5px;
    font-weight: 600;
    color: var(--lt__main-darker_1);
  }

  .input-box__input {
    font-size: 16px;
    font-weight: 500;
    margin-top: 5px;
    padding: 8px 10px;
    width: 100%;
    border: 1px solid var(--lt__primary);
    background-color: var(--lt__primary);
    box-sizing: border-box;
    border-radius: 5px;
    transition: .3s box-shadow;
  }

  .input-box__input:focus {
    outline: none;
    box-shadow: 0px 0px 5px 2px rgba(129,90,202,0.2);
  }

  .button {
    display: flex;
    font-size: 18px;
    font-weight: 500;
    padding: 10px 15px;
    color: var(--white);
    width: 100%;
    height: 38px;
    box-sizing: border-box;
    background-color: var(--lt__main-darker);
    border: none;
    cursor: pointer;
    border-radius: 5px;
    transition: .2s;
    justify-content: center;
    align-items: center;
  }

  .button:hover {
    opacity: 0.88;
  }

  .button .button-icon {
    font-size: 22px;
    padding-right: 5px;
  }

  .button .button-icon-right {
    font-size: 22px;
    padding-left: 5px;
  }

  .button .button-icon-no-gap {
    font-size: 22px;
  }

  .button-green {
    background-color: var(--lt__green);
  }

  .button-red {
    background-color: var(--lt__red);
  }

  .page {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
  }

  .page-block {
    width: 100%;
    padding: 10px;
    background-color: white;
    border-radius: 5px;
    box-shadow: 0px 5px 15px 2px rgba(34, 60, 80, 0.19);
  }

  .page-block__header {
    font-size: 20px;
    color: var(--lt__main);
    padding-left: 3px;
    padding-bottom: 15px;
  }

  .page-block__subheader {
    font-size: 18px;
    color: var(--lt__main);
    padding-left: 3px;
    padding-bottom: 5px
  }

  .table__table {
    width: 100%;
  }

  .table__header {
    font-weight: 600;
    font-size: 18px;
    text-align: left;
    color: var(--dt__header);
    padding: 10px 0;
  }

  .table__td {
    padding-top: 5px;
  }

  .table__td {
    text-align: center;
    font-weight: 500;
    vertical-align: top;
  }

  .table__th {
    color: var(--lt__main-darker_1);
  }

  .table__th:first-child,
  .table__td:first-child {text-align: left;}
  .table__th:last-child,
  .table__td:last-child {text-align: right;}

  .table__button {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 25px;
    height: 25px;
    border: none;
    border-radius: 5px;
    margin-left: 5px;
    cursor: pointer;
    transition: .1s;
  }

  .table__button:hover {
    opacity: 0.85;
  }

  .table__button-icon {
    color: var(--white);
    font-size: 18px;
  }

  .table__button-delete {
    background-color: var(--lt__red);
  }

  .table__button-add {
    background-color: var(--lt__green);
  }

  .Vue-Toastification__toast-body {
    font-weight: 600;
  }

  .modal__background.modal__background-active {
    position: fixed;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, .8);
    z-index: 1001;
  }


  .modal-open {
    display: block !important;
  }

  .modal {
    display: none;
    padding: 20px;
    width: calc(100% - 80px);
    max-width: 600px;
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    background-color: var(--white);
    border-radius: 5px;
    z-index: 1002;
    text-align: left;
  }

  .modal__header {
    padding-bottom: 20px;
    font-size: 24px;
    color: var(--lt__main);
  }

  .modal__description {
    font-weight: 500;
    padding-bottom: 20px;
    color: var(--lt__dark-text);
  }

  .modal-footer {
    display: flex;
    justify-content: right;
    padding-top: 20px;
  }

  .modal__button {
    padding: 8px 15px;
    font-size: 18px;
    font-weight: 600;
    color: var(--white);
    background-color: var(--lt__main-darker);
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: .1s;
    margin-left: 10px;
    min-width: 100px
  }

  .modal__button-agree {
    background-color: var(--lt__green);
  }

  .modal__button-cancel {
    background-color: var(--lt__red);
  }

  .modal__button:hover {
    opacity: .88;
  }


  /* MEDIA */


  @media screen and (max-width: 840px) {
    .container {
      padding: 0 10px;
    }
  }


  /* DARK THEME */


  .dark__theme .modal__background.modal__background-active {
    background-color: rgba(0, 0, 0, .7);
  }

  .dark__theme {
    background-color: var(--dt__primary);
  }

  .dark__theme .page-block {
    background-color: var(--dt__primary-lighter);
    box-shadow: 0px 5px 15px 2px rgba(14, 20, 26, 1);
  }

  .dark__theme .input-box__label {
    color: var(--dt__header)
  }

  .dark__theme .input-box__input {
    background-color: var(--dt__primary);
    border-color: var(--dt__primary);
    color: var(--white);
  }

  .dark__theme .button-green {
    background-color: var(--dt__green);
  }

  .dark__theme .button-red {
    background-color: var(--dt__red);
  }

  .dark__theme .page-block__header {
    color: var(--dt__header);
  }

  .dark__theme .page-block__subheader {
    color: var(--dt__header);
  }

  .dark__theme .table__th {
    color: var(--dt__header);
  }

  .dark__theme .table__td {
    color: var(--dt__text-lighter);
  }

  .dark__theme .table__button-delete {
    background-color: var(--dt__red);
  }

  .dark__theme .table__button-add {
    background-color: var(--dt__green);
  }

  .dark__theme .modal {
    background-color: var(--dt__primary-lighter);
  }

  .dark__theme .modal__header {
    color: var(--dt__header);
  }

  .dark__theme .modal__description {
    color: var(--dt__text-lighter);
  }

  .dark__theme .modal__button-agree {
    background-color: var(--dt__green);
  }

  .dark__theme .modal__button-cancel {
    background-color: var(--dt__red);
  }
</style>