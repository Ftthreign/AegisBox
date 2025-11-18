<script lang="ts">
  import Navbar from './components/Navbar.svelte';
  import Hero from './components/Hero.svelte';
  import Features from './components/Features.svelte';
  import Security from './components/Security.svelte';
  import CTA from './components/CTA.svelte';
  import Footer from './components/Footer.svelte';
</script>

<svelte:head>
  <title>AegisBox - Your Digital Security Vault</title>
  <meta name="description" content="Password manager lokal dengan enkripsi tingkat militer. Zero-knowledge, local-first, dan open source." />
</svelte:head>

<div class="min-h-screen bg-white text-gray-900">
  <Navbar />
  <Hero />
  <Features />
  <Security />
  <CTA />
  <Footer />
</div>