<template>
  <div>
    <b-navbar toggleable="lg" class="header">
      <b-navbar-brand 
        :to="{
        name: 'pizzas'
        }">
          Pizza Delivery
      </b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
          <b-navbar-nav v-if="$store.state.isUserLoggedIn">
            <b-nav-item 
            :to="{
            name: 'history'
            }">
              History
            </b-nav-item>
          </b-navbar-nav>
        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <!--<b-nav-form>
            <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
          </b-nav-form>-->
          <b-nav-item
          :to="{
            name: 'order'
          }">
            <span class="fa-layers fa-fw fa-2x">
              <i class="fas fa-shopping-cart "></i>
              <span v-if="$store.state.orders.length > 0" class="fa-layers-counter fa-1x" style="background:Tomato">{{$store.state.orders.length}}</span>
            </span>
          </b-nav-item>
          <b-button 
            v-if="!$store.state.isUserLoggedIn" 
            :to="{
            name: 'login'
            }" 
            variant="light">
            LOGIN
          </b-button>

          <b-button 
            v-if="!$store.state.isUserLoggedIn"
            :to="{
            name: 'register'
            }"
            variant="light">
            SIGN UP
          </b-button>
          <b-nav-item-dropdown right v-if="$store.state.isUserLoggedIn">
            <!-- Using 'button-content' slot -->
            <template v-slot:button-content>
              <em>User</em>
            </template>
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
.header {
  border-bottom: 1px solid #8e8e8e;
}

.icon {
  padding-left: 10px;
}
</style>
