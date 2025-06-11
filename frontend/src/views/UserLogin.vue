<template>
  <div class="auth-container">
    <div class="auth-background" ref="vantaRef"></div>
    
    <div class="auth-content">
      <div class="auth-card">
        <div class="auth-header">
          <router-link to="/" class="back-link">
            <ArrowLeft class="back-icon" />
            Back to Home
          </router-link>
          
          <div class="logo">
            <Brain class="logo-icon" />
            <span class="logo-text">MindJournal</span>
          </div>
          
          <h1 class="auth-title">Welcome Back</h1>
          <p class="auth-subtitle">Sign in to continue your wellness journey</p>
        </div>

        <form @submit.prevent="login" class="auth-form">
          <div class="form-group">
            <label for="email" class="form-label">
              <Mail class="label-icon" />
              Email Address
            </label>
            <input
              v-model="email"
              type="email"
              id="email"
              class="form-control"
              placeholder="Enter your email"
              required
              :disabled="loading"
            />
          </div>

          <div class="form-group">
            <label for="password" class="form-label">
              <Lock class="label-icon" />
              Password
            </label>
            <div class="password-input">
              <input
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                id="password"
                class="form-control"
                placeholder="Enter your password"
                required
                :disabled="loading"
              />
              <button
                type="button"
                class="password-toggle"
                @click="showPassword = !showPassword"
              >
                <Eye v-if="!showPassword" class="toggle-icon" />
                <EyeOff v-else class="toggle-icon" />
              </button>
            </div>
          </div>

          <button type="submit" class="btn btn-primary btn-full" :disabled="loading">
            <Loader2 v-if="loading" class="btn-icon animate-spin" />
            <LogIn v-else class="btn-icon" />
            {{ loading ? 'Signing In...' : 'Sign In' }}
          </button>
        </form>

        <div class="auth-footer">
          <p class="auth-switch">
            Don't have an account?
            <router-link to="/register" class="auth-link">Sign up here</router-link>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Brain, Mail, Lock, Eye, EyeOff, LogIn, ArrowLeft, Loader2 } from 'lucide-vue-next'
import * as THREE from 'three'
import VANTA from 'vanta'

export default {
  name: "UserLogin",
  components: {
    Brain,
    Mail,
    Lock,
    Eye,
    EyeOff,
    LogIn,
    ArrowLeft,
    Loader2
  },
  data() {
    return {
      email: "",
      password: "",
      showPassword: false,
      loading: false
    };
  },
  mounted() {
    this.initVanta()
  },
  beforeUnmount() {
    if (this.vantaEffect) {
      this.vantaEffect.destroy()
    }
  },
  methods: {
    initVanta() {
      this.vantaEffect = VANTA.WAVES({
        el: this.$refs.vantaRef,
        THREE: THREE,
        mouseControls: true,
        touchControls: true,
        gyroControls: false,
        minHeight: 200.00,
        minWidth: 200.00,
        scale: 1.00,
        scaleMobile: 1.00,
        color: 0x667eea,
        shininess: 30.00,
        waveHeight: 15.00,
        waveSpeed: 0.75,
        zoom: 0.65
      })
    },
    async login() {
      if (this.loading) return;
      
      this.loading = true;
      
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

        if (!response.ok) {
          throw new Error("Invalid credentials");
        }

        const data = await response.json();
        localStorage.setItem("jwt", data.token);
        
        // Success feedback
        this.$router.push("/journal");
      } catch (error) {
        alert("Login failed. Please check your credentials and try again.");
        console.error("Login error:", error);
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
.auth-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  padding: 2rem;
}

.auth-background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}

.auth-content {
  width: 100%;
  max-width: 400px;
  z-index: 1;
}

.auth-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 1.5rem;
  padding: 2.5rem;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.auth-header {
  text-align: center;
  margin-bottom: 2rem;
}

.back-link {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  color: #667eea;
  text-decoration: none;
  font-size: 0.875rem;
  font-weight: 500;
  margin-bottom: 1.5rem;
  transition: color 0.2s ease;
}

.back-link:hover {
  color: #5a67d8;
}

.back-icon {
  width: 1rem;
  height: 1rem;
}

.logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  margin-bottom: 1.5rem;
}

.logo-icon {
  width: 2rem;
  height: 2rem;
  color: #667eea;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: 700;
  color: #2d3748;
}

.auth-title {
  font-size: 2rem;
  font-weight: 700;
  color: #2d3748;
  margin-bottom: 0.5rem;
}

.auth-subtitle {
  color: #4a5568;
  font-size: 0.875rem;
}

.auth-form {
  margin-bottom: 1.5rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: #2d3748;
  margin-bottom: 0.5rem;
}

.label-icon {
  width: 1rem;
  height: 1rem;
  color: #667eea;
}

.password-input {
  position: relative;
}

.password-toggle {
  position: absolute;
  right: 0.75rem;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  cursor: pointer;
  color: #4a5568;
  padding: 0.25rem;
  border-radius: 0.25rem;
  transition: color 0.2s ease;
}

.password-toggle:hover {
  color: #667eea;
}

.toggle-icon {
  width: 1.25rem;
  height: 1.25rem;
}

.btn-full {
  width: 100%;
  justify-content: center;
}

.btn-icon {
  width: 1.25rem;
  height: 1.25rem;
}

.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.auth-footer {
  text-align: center;
}

.auth-switch {
  color: #4a5568;
  font-size: 0.875rem;
}

.auth-link {
  color: #667eea;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.2s ease;
}

.auth-link:hover {
  color: #5a67d8;
  text-decoration: underline;
}

@media (max-width: 480px) {
  .auth-container {
    padding: 1rem;
  }
  
  .auth-card {
    padding: 2rem;
  }
  
  .auth-title {
    font-size: 1.75rem;
  }
}
</style>