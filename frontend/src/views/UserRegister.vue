<template>
  <div class="auth-wrapper">
    <div class="auth-inner">
      <div class="card">
        <div class="card-body">
          <h4 class="card-title mb-1">Create Your Account 🚀</h4>
          <p class="card-text mb-2">Sign up to get started</p>
          <form @submit.prevent="register" class="auth-register-form">
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
              Sign up
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
    async register() {
      try {
        const response = await fetch("http://localhost:8080/register", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            email: this.email,
            password: this.password,
          }),
        });

        if (!response.ok) throw new Error("Registration failed");

        alert("Registration successful");
        this.$router.push("/login");
      } catch (error) {
        alert("Registration failed");
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
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
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
  border-color: #f5576c;
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
  color: #f5576c;
}

.custom-btn {
  background-color: #f5576c;
  border: none;
  padding: 0.75rem;
  font-size: 1rem;
  border-radius: 5px;
  transition: background-color 0.3s ease;
}

.custom-btn:hover {
  background-color: #f093fb;
}
</style>
