<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

const matches = ref([])
const inputText = ref('')
const regexPattern = ref('')
const generatedRegexPattern = ref('')

function performRegexMatching() {
  if (regexPattern.value && inputText.value) {
    axios
      .post('http://localhost:8080/api/match', {
        regexPattern: regexPattern.value,
        inputText: inputText.value
      })
      .then((response) => {
        if (response.data && response.data.matches) {
          matches.value = response.data.matches
        } else {
          matches.value = []
        }
      })
      .catch((error) => {
        console.error(error)
      })
  } else {
    matches.value = []
  }
}

function generateRegexPattern() {
  generatedRegexPattern.value = escapeRegex(inputText.value)
}

function notify() {
  toast.success('Copied to clipboard!', {
    autoClose: 2000
  })
}

function copyToClipboard() {
  navigator.clipboard.writeText(generatedRegexPattern.value)
  notify()
}

function escapeRegex(text) {
  // escapes special chars to make it suitable as a regex pattern.
  // https://stackoverflow.com/questions/3561493/is-there-a-regexp-escape-function-in-javascript
  return text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
}
</script>

<template>
  <div class="container">
    <h1>Regex Interpreter</h1>

    <div class="input-group">
      <label for="regex-input">Regex Pattern:</label>
      <input type="text" id="regex-input" v-model="regexPattern" @input="performRegexMatching" />
    </div>

    <div class="input-group">
      <label for="text-input">Input Text:</label>
      <textarea id="text-input" v-model="inputText" @input="performRegexMatching"></textarea>
    </div>

    <div class="results">
      <h2>Matching Results:</h2>

      <ul>
        <li v-for="(match, index) in matches" :key="index">
          <span class="match-summary">
            {{ match }}
          </span>
        </li>
      </ul>

      <p v-if="matches.length === 0">No matches found.</p>
    </div>

    <div class="generate-regex">
      <h2>Generate Regex Pattern:</h2>

      <button @click="generateRegexPattern">Generate</button>

      <div v-if="generatedRegexPattern">
        <button @click="copyToClipboard">Copy to Clipboard</button>
        <code>{{ generatedRegexPattern }}</code>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  margin-bottom: 20px;
}

.input-group {
  margin-bottom: 20px;
}

label {
  font-weight: bold;
}

input[type='text'],
textarea {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border: 0;
  border-radius: 4px;
  font-family: inherit;
  background-color: #333;
  color: white;
}

.results {
  margin-top: 20px;
}

.results ul {
  padding: 0;
}

.results ul li {
  margin-bottom: 5px;
}

.results ul .match-summary {
  font-weight: bold;
}

.results p {
  font-style: italic;
  color: #888;
}

.generate-regex {
  margin-top: 20px;
}

.generate-regex h2 {
  margin-bottom: 10px;
}

.generate-regex button {
  padding: 6px 12px;
  margin-bottom: 15px;
  font-size: 14px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.generate-regex button:hover {
  background-color: #0069d9;
}

.generate-regex p {
  margin-top: 10px;
  font-weight: bold;
}

.generate-regex code {
  display: block;
  padding: 10px;
  border: 1px solid gray;
  border-radius: 4px;
  font-family: monospace;
  white-space: pre;
}
</style>
