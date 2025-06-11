<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
  line-height: 1.6;
  color: #1a202c;
  background-color: #f7fafc;
}

#app {
  min-height: 100vh;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

/* Global utility classes */
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.75rem 1.5rem;
  font-size: 0.875rem;
  font-weight: 500;
  line-height: 1.25rem;
  border-radius: 0.5rem;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease-in-out;
  text-decoration: none;
  gap: 0.5rem;
}

.btn:focus {
  outline: 2px solid transparent;
  outline-offset: 2px;
  box-shadow: 0 0 0 3px rgba(66, 153, 225, 0.5);
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover {
  transform: translateY(-1px);
  box-shadow: 0 10px 25px rgba(102, 126, 234, 0.3);
}

.btn-secondary {
  background: white;
  color: #4a5568;
  border: 1px solid #e2e8f0;
}

.btn-secondary:hover {
  background: #f7fafc;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.form-control {
  width: 100%;
  padding: 0.75rem 1rem;
  font-size: 0.875rem;
  line-height: 1.25rem;
  color: #2d3748;
  background-color: white;
  border: 1px solid #e2e8f0;
  border-radius: 0.5rem;
  transition: all 0.2s ease-in-out;
}

.form-control:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.card {
  background: white;
  border-radius: 1rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  overflow: hidden;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>