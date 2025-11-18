<script>
  import { Shield } from 'lucide-svelte';
</script>

<footer class="py-10 px-6 border-t border-gray-200">
  <div class="max-w-6xl mx-auto flex flex-col md:flex-row items-center justify-between gap-6">

    <div class="flex items-center gap-2.5">
      <Shield class="w-5 h-5 text-blue-600" />
      <span class="font-semibold text-gray-900">AegisBox</span>
    </div>

    <div class="flex items-center gap-8 text-sm text-gray-600">
      <a href="#" class="hover:text-gray-900 transition-colors">Documentation</a>
      <a href="#" class="hover:text-gray-900 transition-colors">Privacy</a>
      <a href="#" class="hover:text-gray-900 transition-colors">Terms</a>
    </div>

    <div class="text-sm text-gray-500">
      © 2025 AegisBox. Built with 🛡️
    </div>

  </div>
</footer>
