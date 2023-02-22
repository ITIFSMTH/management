<template>
  <div class="navbar-wrapper"
    :class="{'open': this.$store.getters.isNavbarOpen}">
    <div class="navbar">
      <ul class="navbar-links">
        <li class="navbar-link navbar-theme"
          :class="{'dark__theme-active': this.$store.getters.isThemeDark}"
          @click="switchTheme">
          <span class="navbar-link__icon navbar-theme__icon material-icons">dark_mode</span>
          <span class="navbar-link__text navbar-theme__text">Тема</span>
        </li>

        <!--  ONLY FOR ADMIN -->
        <router-link to="/admin" custom v-slot="{ href, navigate, isActive }"
          v-if="this.$store.getters.isUserAdmin">
          <li class="navbar-link"
            :class="{'active': isActive}">
            <a class="navbar-link__link"
              :href="href"
              @click="navigate">
              <span class="navbar-link__icon material-icons">admin_panel_settings</span>
              <span class="navbar-link__text">Админ</span>
            </a>
          </li>
        </router-link>

        <router-link to="/statistic" custom v-slot="{ href, navigate, isActive }"
          v-if="this.$store.getters.isUserAdmin">
          <li class="navbar-link"
            :class="{'active': isActive}">
            <a class="navbar-link__link"
              :href="href"
              @click="navigate">
              <span class="navbar-link__icon material-icons">monitoring</span>
              <span class="navbar-link__text">Статистика</span>
            </a>
          </li>
        </router-link>

        <!-- ONLY FOR OPERATORS -->
        <router-link to="/operator" custom v-slot="{ href, navigate, isActive }"
          v-if="this.$store.getters.isUserOperator">
          <li class="navbar-link"
            :class="{'active': isActive}">
            <a class="navbar-link__link"
              :href="href"
              @click="navigate">
              <span class="navbar-link__icon material-icons">person</span>
              <span class="navbar-link__text">Оператор</span>
            </a>
          </li>
        </router-link>

        <!-- ONLY FOR JUNIOR OPERATOR -->
        <!-- <router-link to="/poll/rating" custom v-slot="{ href, navigate, isActive }"
          v-if="this.$store.getters.isUserJuniorOperator">
          <li class="navbar-link"
            :class="{'active': isActive}">
            <a class="navbar-link__link"
              :href="href"
              @click="navigate">
              <span class="navbar-link__icon material-icons">how_to_vote</span>
              <span class="navbar-link__text">Голосование</span>
            </a>
          </li>
        </router-link> -->

        <!-- ONLY FOR SENIOR OPERATOR -->
        <!-- <router-link to="/poll/budget" custom v-slot="{ href, navigate, isActive }"
          v-if="this.$store.getters.isUserSeniorOperator">
          <li class="navbar-link"
            :class="{'active': isActive}">
            <a class="navbar-link__link"
              :href="href"
              @click="navigate">
              <span class="navbar-link__icon material-icons">how_to_vote</span>
              <span class="navbar-link__text">Голосование</span>
            </a>
          </li>
        </router-link> -->

        <li class="navbar-link navbar-show"
          @click="switchNavbar">
          <span class="navbar-link__icon navbar-show__icon material-icons">arrow_forward_ios</span>
          <span class="navbar-link__text navbar-show__text">Скрыть</span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'navbar-component',
    methods: {
      switchTheme() {
        this.$store.dispatch('editTheme')
      },
      switchNavbar() {
        this.$store.commit('switchNavbar')
      }
    }
  }
</script>

<style scoped>
  .navbar-wrapper {
    min-width: 74px;
  }

  .navbar {
    position: fixed;
    top: 0;
    display: inline-block;
    padding: 20px 10px;
    height: 100vh;
    box-sizing: border-box;
    z-index: 1000;
    background-color: var(--white);
  }

  .navbar-links {
    list-style-type: none;
  }

  .navbar-link {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
    cursor: pointer;
  }

  .navbar-link__link {
    display: flex;
    align-items: center;
    justify-content: left;
    background-color: var(--lt__primary);
    width: 100%;
    border-radius: 5px;
  }

  .navbar-link__text {
    display: none;
    padding: 0 10px;
    font-weight: 500;
    font-size: 20px;
    color: var(--grey);
  }

  .navbar-link__icon {
    font-size: 28px;
    padding: 13px;
    border-radius: 5px;
    cursor: pointer;
    color: var(--grey);
  }

  .navbar-link.active .navbar-link__icon,
  .navbar-link.active .navbar-link__text {
    color: var(--lt__main);
  }

  .navbar-theme__icon,
  .navbar-theme__text  {
    color: var(--lt__main-darker_1)
  }

  .dark__theme-active .navbar-theme__icon,
  .dark__theme-active .navbar-theme__text {
    color: var(--dt__yellow);
  }
  
  .navbar-show {
    position: absolute;
    bottom: 0;
  }

  .navbar-show__icon {
    background-color: transparent !important;
    transition: .2s transform;
    color: var(--lt__main);
  }

  .navbar-show__text {
    font-weight: 600;
    color: var(--lt__main);
  }


  .navbar-wrapper.open {
    min-width: 210px;
  }

  .navbar-wrapper.open .navbar-link__text {
    display: inline-block;
  }

  .navbar-wrapper.open .navbar-show__icon {
    transform: rotateZ(180deg);
  }


  /* MEDIA */


  @media screen and (max-width: 800px) {
    .navbar-wrapper {
      min-width: auto;
    }

    .navbar {
      padding: 10px;
      height: auto;
      bottom: 0;
      left: 0;
      right: 0;
      top: auto;
      background-color: var(--lt__primary) !important;
      box-shadow: 0px -5px 15px 2px rgba(220, 231, 241, 1);
    }

    .navbar-links {
      display: flex;
    }

    .navbar-link {
      margin-bottom: 0;
      margin-right: 10px;
    }

    .navbar-show {
      display: none;
    }

    /* DARK THEME */

    .dark__theme .navbar {
      background-color: var(--dt__primary) !important;
      box-shadow: 0px -5px 10px 2px rgba(14, 20, 26, 1);
    }
  }



  /* DARK THEME */


  .dark__theme .navbar {
    background-color: var(--dt__primary-lighter);
  }

  .dark__theme .navbar-link__link {
    background-color: var(--dt__primary);
  }


</style>