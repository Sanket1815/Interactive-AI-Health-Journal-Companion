<template>
  <div class="auth-wrapper">
    <div class="auth-inner">
      <div class="card">
        <div class="card-body">
          <h4 class="card-title mb-1">Welcome Back! 👋</h4>
          <p class="card-text mb-2">Please sign-in to continue</p>
          <form @submit.prevent="login" class="auth-login-form">
            <!-- Email Field -->
            <div class="mb-1 input-group">
              <input
                v-model="email"
                type="email"
                class="form-control custom-input"
                id="email"
                placeholder="Enter your email"
                required
              />
              <span class="input-icon">@</span>
            </div>

            <!-- Password Field -->
            <div class="mb-1 input-group">
              <input
                v-model="password"
                type="password"
                class="form-control custom-input"
                id="password"
                placeholder="Enter your password"
                required
              />
              <span class="input-icon">&#x1f512;</span>
            </div>

            <!-- Submit Button -->
            <button type="submit" class="btn btn-primary w-100 custom-btn">
              Sign in
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    async login() {
      try {
        const response = await fetch("http://localhost:8080/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            email: this.email,
            password: this.password,
          }),
        });

        if (!response.ok) throw new Error("Login failed");

        const data = await response.json();
        localStorage.setItem("jwt", data.token);
        this.$router.push("/journal");
      } catch (error) {
        alert("Login failed");
      }
    },
  },
};
</script>

<style scoped>
.auth-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(135deg, #6dd5fa 0%, #2980b9 100%);
}

.auth-inner {
  max-width: 400px;
  width: 100%;
}

.card {
  border-radius: 10px;
  overflow: hidden;
}

.card-body {
  padding: 2rem;
  background-color: #ffffff;
  box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.1);
}

.custom-input {
  padding: 0.75rem;
  border-radius: 5px;
  border: 1px solid #ccc;
  width: calc(100% - 30px);
  transition: border-color 0.3s ease;
}

.custom-input:focus {
  border-color: #2980b9;
  outline: none;
}

.input-group {
  position: relative;
  margin-bottom: 1rem;
}

.input-icon {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 1.2rem;
  color: #2980b9;
}

.custom-btn {
  background-color: #2980b9;
  border: none;
  padding: 0.75rem;
  font-size: 1rem;
  border-radius: 5px;
  transition: background-color 0.3s ease;
}

.custom-btn:hover {
  background-color: #3498db;
}
</style>
