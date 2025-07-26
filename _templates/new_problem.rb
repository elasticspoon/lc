problem_name = ARGV[0]
if problem_name.nil? || problem_name.empty?
  puts 'Please provide a problem name. e.g. mise ruby "two sum"'
  exit 1
end

file_name = "#{problem_name.gsub(/[ .]/, '_').downcase}.rb"

template = <<~TEMPLATE
  require "minitest/autorun"

  # @param {String} s
  # @return {Integer}
  def function(s)
  end

  class Test < Minitest::Test
    def test_one
      # assert_equal(0, function(""))
    end
  end
TEMPLATE

File.write(file_name, template)

puts "Created #{file_name}"
