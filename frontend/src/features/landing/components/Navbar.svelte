<script lang="ts">
  import { Shield, Github } from 'lucide-svelte';
  import { onMount } from 'svelte';

  let scrollY = 0;

  onMount(() => {
    const handleScroll = () => (scrollY = window.scrollY);
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  });
</script>

<nav
  class="
    fixed top-0 w-full z-50 transition-all duration-300
    bg-white/40 backdrop-blur-xl
    border-b border-white/30
    {scrollY > 50 ? 'shadow-[0_4px_20px_rgba(0,0,0,0.08)] bg-white/60' : ''}
  "
>
  <div class="absolute inset-0 pointer-events-none">
    <div class="absolute top-0 left-0 w-full h-px bg-white/70"></div>
    <div class="absolute top-0 left-0 w-full h-6 bg-linear-to-b from-white/50 to-transparent"></div>
  </div>

  <div class="relative max-w-7xl mx-auto px-6 py-4 flex items-center justify-between">
    <div class="flex items-center gap-2.5">
      <div class="relative">
        <Shield class="w-6 h-6 text-blue-600" strokeWidth="2.5" />
        <div class="absolute inset-0 bg-blue-500/20 blur-md rounded-full"></div>
      </div>
      <span class="text-lg font-semibold text-gray-900 select-none">AegisBox</span>
    </div>
    <div class="hidden md:flex items-center gap-8 text-sm font-medium">
      <a href="#features" class="text-gray-600 hover:text-gray-900 transition-colors">Features</a>
      <a href="#security" class="text-gray-600 hover:text-gray-900 transition-colors">Security</a>
      <a href="/docs" class="text-gray-600 hover:text-gray-900 transition-colors">Docs</a>

      <a
        href="https://github.com"
        target="_blank"
        rel="noopener noreferrer"
        class="text-gray-600 hover:text-gray-900 transition-colors flex items-center"
      >
        <Github class="w-5 h-5" />
      </a>
    </div>

  </div>
</nav>
