module Kernel
  def method_missing(cmd,*args)
    args.empty? ? cmd.to_s : "#{cmd} #{args.join(" ")}"
  end
end

puts this is terrible code

x = [1,2,3]
eval(x.join('+'))
