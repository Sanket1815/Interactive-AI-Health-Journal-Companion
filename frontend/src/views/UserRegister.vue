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
          
          <h1 class="auth-title">Create Account</h1>
          <p class="auth-subtitle">Start your mental wellness journey today</p>
        </div>

        <form @submit.prevent="register" class="auth-form">
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
                placeholder="Create a strong password"
                required
                :disabled="loading"
                minlength="6"
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
            <div class="password-strength">
              <div class="strength-bar">
                <div class="strength-fill" :class="passwordStrengthClass"></div>
              </div>
              <span class="strength-text" :class="passwordStrengthClass">
                {{ passwordStrengthText }}
              </span>
            </div>
          </div>

          <div class="form-group">
            <label class="checkbox-label">
              <input
                v-model="agreeToTerms"
                type="checkbox"
                class="checkbox-input"
                required
              />
              <span class="checkbox-custom"></span>
              <span class="checkbox-text">
                I agree to the <a href="#" class="terms-link">Terms of Service</a> and 
                <a href="#" class="terms-link">Privacy Policy</a>
              </span>
            </label>
          </div>

          <button type="submit" class="btn btn-primary btn-full" :disabled="loading || !agreeToTerms">
            <Loader2 v-if="loading" class="btn-icon animate-spin" />
            <UserPlus v-else class="btn-icon" />
            {{ loading ? 'Creating Account...' : 'Create Account' }}
          </button>
        </form>

        <div class="auth-footer">
          <p class="auth-switch">
            Already have an account?
            <router-link to="/login" class="auth-link">Sign in here</router-link>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Brain, Mail, Lock, Eye, EyeOff, UserPlus, ArrowLeft, Loader2 } from 'lucide-vue-next'
import * as THREE from 'three'
import VANTA from 'vanta'

export default {
  name: "UserRegister",
  components: {
    Brain,
    Mail,
    Lock,
    Eye,
    EyeOff,
    UserPlus,
    ArrowLeft,
    Loader2
  },
  data() {
    return {
      email: "",
      password: "",
      showPassword: false,
      agreeToTerms: false,
      loading: false
    };
  },
  computed: {
    passwordStrength() {
      const password = this.password;
      if (password.length === 0) return 0;
      if (password.length < 6) return 1;
      if (password.length < 8) return 2;
      if (password.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/)) return 4;
      return 3;
    },
    passwordStrengthClass() {
      const strength = this.passwordStrength;
      if (strength === 0) return '';
      if (strength === 1) return 'weak';
      if (strength === 2) return 'fair';
      if (strength === 3) return 'good';
      return 'strong';
    },
    passwordStrengthText() {
      const strength = this.passwordStrength;
      if (strength === 0) return '';
      if (strength === 1) return 'Weak';
      if (strength === 2) return 'Fair';
      if (strength === 3) return 'Good';
      return 'Strong';
    }
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
      this.vantaEffect = VANTA.CLOUDS({
        el: this.$refs.vantaRef,
        THREE: THREE,
        mouseControls: true,
        touchControls: true,
        gyroControls: false,
        minHeight: 200.00,
        minWidth: 200.00,
        backgroundColor: 0xf093fb,
        skyColor: 0xf5576c,
        cloudColor: 0xffffff,
        cloudShadowColor: 0x183550,
        sunColor: 0xff6347,
        sunGlareColor: 0xff6347,
        sunlightColor: 0xff8a65,
        speed: 0.8
      })
    },
    async register() {
      if (this.loading) return;
      
      this.loading = true;
      
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

        if (!response.ok) {
          throw new Error("Registration failed");
        }

        alert("Registration successful! Please sign in to continue.");
        this.$router.push("/login");
      } catch (error) {
        alert("Registration failed. Please try again.");
        console.error("Registration error:", error);
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
  color: #f5576c;
  text-decoration: none;
  font-size: 0.875rem;
  font-weight: 500;
  margin-bottom: 1.5rem;
  transition: color 0.2s ease;
}

.back-link:hover {
  color: #e53e3e;
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
  color: #f5576c;
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
  color: #f5576c;
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
  color: #f5576c;
}

.toggle-icon {
  width: 1.25rem;
  height: 1.25rem;
}

.password-strength {
  margin-top: 0.5rem;
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.strength-bar {
  flex: 1;
  height: 0.25rem;
  background: #e2e8f0;
  border-radius: 0.125rem;
  overflow: hidden;
}

.strength-fill {
  height: 100%;
  border-radius: 0.125rem;
  transition: all 0.3s ease;
  width: 0;
}

.strength-fill.weak {
  width: 25%;
  background: #f56565;
}

.strength-fill.fair {
  width: 50%;
  background: #ed8936;
}

.strength-fill.good {
  width: 75%;
  background: #38b2ac;
}

.strength-fill.strong {
  width: 100%;
  background: #48bb78;
}

.strength-text {
  font-size: 0.75rem;
  font-weight: 500;
}

.strength-text.weak { color: #f56565; }
.strength-text.fair { color: #ed8936; }
.strength-text.good { color: #38b2ac; }
.strength-text.strong { color: #48bb78; }

.checkbox-label {
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
  cursor: pointer;
  font-size: 0.875rem;
  line-height: 1.4;
}

.checkbox-input {
  display: none;
}

.checkbox-custom {
  width: 1.25rem;
  height: 1.25rem;
  border: 2px solid #e2e8f0;
  border-radius: 0.25rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  flex-shrink: 0;
  margin-top: 0.125rem;
}

.checkbox-input:checked + .checkbox-custom {
  background: #f5576c;
  border-color: #f5576c;
}

.checkbox-input:checked + .checkbox-custom::after {
  content: '✓';
  color: white;
  font-size: 0.75rem;
  font-weight: bold;
}

.checkbox-text {
  color: #4a5568;
}

.terms-link {
  color: #f5576c;
  text-decoration: none;
  font-weight: 500;
}

.terms-link:hover {
  text-decoration: underline;
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
  color: #f5576c;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.2s ease;
}

.auth-link:hover {
  color: #e53e3e;
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