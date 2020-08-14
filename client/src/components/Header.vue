<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-navbar-brand href="#">Pizza Delivery</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
          </b-nav-form>

          <b-button 
            v-if="!$store.state.isUserLoggedIn" 
            :to="{
            name: 'login'
            }">
            Login
          </b-button>

          <b-button 
            v-if="!$store.state.isUserLoggedIn"
            :to="{
            name: 'register'
            }">
            Sign up
          </b-button>
          <b-nav-item-dropdown right v-if="$store.state.isUserLoggedIn">
            <!-- Using 'button-content' slot -->
            <template v-slot:button-content>
              <em>User</em>
            </template>
            <b-dropdown-item href="#">Profile</b-dropdown-item>
            <b-dropdown-item href="#" @click="logout">Sign Out</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
export default {
  methods: {
    logout () {
      this.$store.dispatch('setUser', null)
      this.$router.push({
        name: 'pizzas'
      })
    }
  }
}
</script>

<style scoped>
</style>
